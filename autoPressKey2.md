```go
package main

import "C"
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
	Longtime      int
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
	hasPause := true
	inLongPress := false
	longPressDone := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	robotgo.EventHook(hook.MouseDown, []string{"mleft"}, func(e hook.Event) {
		if !hasPause && e.Button == 1 {
			cancel()
			fmt.Println("auto press paused!")
		}
	})
	longPress := func (c Config) {
		for {
			select {
			case <-longPressDone:
				robotgo.KeyToggle(c.Key, "up")
				return
			default:
				robotgo.KeyToggle(c.Key, "down")
				robotgo.MilliSleep(10)
			}
		}
	}
	startup := func(in chan string, c Config) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("ignore:", err)
			}
		}()
		for {
			select {
			case <- ctx.Done():
				hasPause = true
				_, ok := <-in
				if ok {
					close(in)
				}
				return
			default:
				robotgo.MilliSleep(10 * c.FirstMoment)
				if c.Longtime > 9 {
					inLongPress = true
					go longPress(c)
					robotgo.MilliSleep(10 * c.Longtime)
					longPressDone <- true
					inLongPress = false
				} else {
					if !inLongPress {
						in <- c.Key
					}
				}
				robotgo.MilliSleep(10 * (c.Period - c.FirstMoment - c.Longtime))
			}
		}
	}
	robotgo.EventHook(hook.KeyDown, []string{"a"}, func(e hook.Event) {
		if hasPause {
			hasPause = false
			ctx, cancel = context.WithCancel(context.Background())
			ch := make(chan string)
			fmt.Println("loading config!")
			conf := [5]Config{}
			LoadConf(&conf)
			fmt.Println("auto press startup!")
			for _, c := range conf {
				go startup(ch,c)
			}
			go func(in chan string) {
				for key := range in{
					robotgo.KeyTap(key)
				}
			}(ch)
		}
	})
}


/**
 * 构建命令
 * go build -o autoKey.exe .\autoPressKey.go
 */
```

### conf.json
```json
[
  { "key": "q", "firstMoment": 3, "period": 270, "longtime": 0 },
  { "key": "w", "firstMoment": 9, "period": 29 , "longtime": 0 },
  { "key": "e", "firstMoment": 7, "period": 37 , "longtime": 0 },
  { "key": "r", "firstMoment": 5, "period": 190, "longtime": 0 },
  { "key": "g", "firstMoment": 1, "period": 400, "longtime": 0 }
]
```
