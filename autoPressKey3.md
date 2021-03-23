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
	"strconv"
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
	nearDistance float64
)
func main() {
	filePath = "conf.json"
	nearDistance = 200
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}
	if len(os.Args) > 2 {
		n, err :=  strconv.ParseFloat(os.Args[2],64)
		if err == nil {
			nearDistance = n
		}
	}
	robotgo.EventHook(hook.KeyDown, []string{"ctrl", "shift", "space"}, func(e hook.Event) {
		fmt.Println("listening program over!")
		robotgo.EventEnd()
	})
	initAutoPress()
	helpHoldLeft()
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
func helpHoldLeft() {
	robotgo.SetMouseDelay(20)
	robotgo.EventHook(hook.KeyHold, []string{"d"}, func(e hook.Event) {
		robotgo.MouseClick("left")
	})
}
func initAutoPress() {
	w,h := robotgo.GetScreenSize()
	playerPos := Point{float64(w/2), float64(h/2)}
	fmt.Printf("player default position:%v \n", playerPos)
	isPausing := true
	isHolding := false
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
	nearby := func() bool{
		x, y := robotgo.GetMousePos()
		mousePos := Point{ float64(x), float64(y)}
		return playerPos.measure(mousePos) < nearDistance
	}
	robotgo.EventHook(hook.KeyDown, []string{"a"}, func(e hook.Event) {
		if nearby() {
			//在怪物附近直接开打
			fight()
		} else {
			//前往怪物附近然后开打
			for {
				if isPausing {
					robotgo.MouseClick("left")
				}
				robotgo.MilliSleep(10 * 60)
				if nearby() {
					fight()
					robotgo.MilliSleep(10 * 100)
					break
				}
			}
		}
	})
	robotgo.EventHook(hook.MouseDown, []string{}, func(e hook.Event) {
		// 双击左键 暂停战斗
		if !isPausing && e.Button == 1 && e.Clicks == 2{
			pause()
		}
	})
	robotgo.EventHook(hook.MouseDown, []string{}, func(e hook.Event) {
		//鼠标右键双击定位角色位置
		if e.Button == 2 && e.Clicks == 2 {
			//fmt.Printf("%v \n", e)
			playerPos = Point{ float64(e.X), float64(e.Y)}
			fmt.Printf("player new position:%v \n", playerPos)
		}
	})
}


/**
 * 构建命令
 * go build -o autoKey.exe .\autoKey.go
 */
```
### conf.json
```json
[
  { "key": "q", "firstMoment": 3, "period": 270, "holdTime": 0 },
  { "key": "w", "firstMoment": 9, "period": 29 , "holdTime": 0 },
  { "key": "e", "firstMoment": 7, "period": 37 , "holdTime": 0 },
  { "key": "r", "firstMoment": 5, "period": 590, "holdTime": 120 },
  { "key": "g", "firstMoment": 1, "period": 400, "holdTime": 0 }
]
```
