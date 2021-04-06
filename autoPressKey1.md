```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"math"
	"os"
	"time"
)
type Point struct {
	X float64
	Y float64
}
func(s Point) measure(t Point) float64 {
	square := math.Pow(s.X - t.X, 2) + math.Pow(s.Y - t.Y, 2)
	return math.Sqrt(square)
}
type Config struct {
	Key           string
	FirstMoment   int
	Period        int
	HoldTime      int
}

var (
	filePath string
)
func main() {
	filePath = "conf.json"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}
	robotgo.EventHook(hook.KeyDown, []string{"ctrl", "shift", "space"}, func(e hook.Event) {
		fmt.Println("listening program over!")
		robotgo.EventEnd()
	})
	initAuto()
	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}
func LoadConf(conf *[5]Config)  {
	file, _ := os.Open(filePath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(conf)
	if err != nil {
		fmt.Println("config load error:", err)
	}
}

func initAuto() {
	w,h := robotgo.GetScreenSize()
	playerPos := Point{float64(w/2), float64(h/2)}
	fmt.Printf("player default position:%v \n", playerPos)
	isPausing := true
	isHolding := false
	nearTime  := time.Now()
	holdCtx, holdCancel := context.WithCancel(context.TODO())
	ctx, cancel := context.WithCancel(context.Background())
	holdPress := func (c Config) {
		for {
			select {
			case <-holdCtx.Done():
				robotgo.KeyToggle(c.Key, "up")
				isHolding = false
				return
			default:
				robotgo.KeyToggle(c.Key, "down")
				robotgo.MilliSleep(10)
			}
		}
	}
	startup := func(c Config) {
		for {
			select {
			case <- ctx.Done():
				return
			default:
				robotgo.MilliSleep(10 * c.FirstMoment)
				if c.HoldTime > 9 {
					holdCtx, holdCancel = context.WithCancel(context.TODO())
					isHolding = true
					go holdPress(c)
					robotgo.MilliSleep(10 * c.HoldTime)
					holdCancel()
				} else {
					if !isHolding {
						robotgo.KeyTap(c.Key)
					}
				}
				robotgo.MilliSleep(10 * (c.Period - c.FirstMoment - c.HoldTime))
			}
		}
	}
	fight := func() {
		if isPausing {
			isPausing = false
			ctx, cancel = context.WithCancel(context.Background())
			conf := [5]Config{}
			LoadConf(&conf)
			fmt.Println("auto fight start!")
			for _, c := range conf {
				go startup(c)
			}
		}
	}
	pause := func() {
		cancel()
		holdCancel()
		isPausing = true
		fmt.Println("auto fight paused!")
	}
	robotgo.EventHook(hook.KeyDown, []string{"a"}, func(e hook.Event) {
		fight()
	})
	robotgo.EventHook(hook.MouseDown, []string{}, func(e hook.Event) {
		// 双击左键 暂停战斗
		if !isPausing && e.Button == 1 && e.Clicks == 2 {
			pause()
		}
	})
	robotgo.EventHook(hook.MouseDown, []string{}, func(e hook.Event) {
		// 双击右键 重新定位角色位置
		if e.Button == 2 && e.Clicks == 2 {
			//fmt.Printf("%v \n", e)
			playerPos = Point{ float64(e.X), float64(e.Y)}
			fmt.Printf("player new position:%v \n", playerPos)
		}
	})
	// 辅助走位
	robotgo.EventHook(hook.KeyHold, []string{"d"}, func(e hook.Event) {
		if !isPausing {
			pause()
		}
		curr := time.Now()
		if curr.Sub(nearTime) > 150*time.Millisecond {
			robotgo.MouseClick("left")
			nearTime = curr
		}
	})
}

/**
 * 构建命令
 * go build -o autoKey.exe .\autoKey.go
 */
```
配置文件conf.json
```json
[
  { "key": "1", "firstMoment": 7, "period": 230, "holdTime": 0 },
  { "key": "2", "firstMoment": 9, "period": 25 , "holdTime": 0 },
  { "key": "3", "firstMoment": 5, "period": 410 , "holdTime": 0 },
  { "key": "4", "firstMoment": 3, "period": 590, "holdTime": 0 },
  { "key": "5", "firstMoment": 1, "period": 310, "holdTime": 0 }
]
```
