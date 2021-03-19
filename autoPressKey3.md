```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"os"
)

type Config struct {
	Key           string
	FirstMoment   int
	Period        int
	HoldTime      int
}

func main() {
	robotgo.EventHook(hook.KeyDown, []string{"ctrl", "shift", "space"}, func(e hook.Event) {
		fmt.Println("listening program over!")
		robotgo.EventEnd()
	})
	initAutoPress()
	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}
func LoadConf(conf *[5]Config)  {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(conf)
	if err != nil {
		fmt.Println("config load error:", err)
	}
}

func initAutoPress() {
	isPausing := true
	isHolding := false
	holdCtx, holdCancel := context.WithCancel(context.TODO())
	ctx, cancel := context.WithCancel(context.Background())
	robotgo.EventHook(hook.KeyHold, []string{"d"}, func(e hook.Event) {
		fmt.Printf("%v \n", e)
		robotgo.KeyTap("p")
		//robotgo.MouseToggle("down")
	})
	robotgo.EventHook(hook.KeyUp, []string{"d"}, func(e hook.Event) {
		fmt.Printf("keyup:%v \n", e)
		// robotgo.MouseToggle("up")
	})
	robotgo.EventHook(hook.MouseDown, []string{"mleft"}, func(e hook.Event) {
		if !isPausing && e.Button == 1 {
			cancel()
			holdCancel()
			isPausing = true
			fmt.Println("auto press paused!")
		}
	})
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
	robotgo.EventHook(hook.KeyDown, []string{"a"}, func(e hook.Event) {
		if isPausing {
			isPausing = false
			ctx, cancel = context.WithCancel(context.Background())
			fmt.Println("loading config!")
			conf := [5]Config{}
			LoadConf(&conf)
			fmt.Println("auto press startup!")
			for _, c := range conf {
				go startup(c)
			}
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
