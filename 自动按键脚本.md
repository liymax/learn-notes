```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"os"
	"sync"
	"time"
)

type AutoState struct {
	mu sync.RWMutex
	autoActive bool
}
func setAutoActive(state *AutoState, flag bool) {
	state.mu.Lock()
	state.autoActive = flag
	state.mu.Unlock()
}
type config struct {
	Key string
	FirstMoment   int
	Period    int
}

func main() {
	robotgo.EventHook(hook.KeyDown, []string{"ctrl", "shift", "space"}, func(e hook.Event) {
		fmt.Println("listening program over!")
		robotgo.EventEnd()
	})
	autoPressByChan()
	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}
func loadConf(conf *[5]config)  {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(conf)
	if err != nil {
		fmt.Println("config load error:", err)
	}
}
func autoPressByChan() {
	shouldClose := false
	hasPause := true
	robotgo.EventHook(hook.MouseDown, []string{"mleft"}, func(e hook.Event) {
		//fmt.Printf("%v \n", e)
		if !shouldClose && e.Button == 1 {
			shouldClose = true
			fmt.Println("auto press paused!")
		}
	})
	startup := func(in chan string, key string, firstMoment int,period int) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("ignore:", err)
			}
		}()
		for {
			if shouldClose{
				hasPause = true
				_, ok := <-in
				if ok {
					close(in)
				}
				return
			} else {
				time.Sleep(time.Duration(firstMoment * 1e7))
				in <- key
				time.Sleep(time.Duration((period - firstMoment) * 1e7))
			}
		}
	}
	robotgo.EventHook(hook.KeyDown, []string{"a"}, func(e hook.Event) {
		if hasPause {
			shouldClose = false
			hasPause = false
			ch := make(chan string)
			fmt.Println("loading config!")
			conf := [5]config{}
			loadConf(&conf)
			fmt.Println("auto press startup!")
			for _, v := range conf {
				go startup(ch,v.Key, v.FirstMoment, v.Period)
			}
			go func(in chan string) {
				for key := range in{
					robotgo.KeyTap(key)
				}
			}(ch)
		}
	})
}

/*func autoPress() {
	autoState := AutoState{}
	//seq := makeSeq(0)
	robotgo.EventHook(hook.MouseDown, []string{"mleft"}, func(e hook.Event) {
		if autoState.autoActive {
			setAutoActive(&autoState, false)
			fmt.Println("auto pressing paused!")
		}
	})
	robotgo.EventHook(hook.KeyDown, []string{"a"}, func(e hook.Event) {
		if autoState.autoActive == false {
			setAutoActive(&autoState, true)
			go func() {
				for autoState.autoActive {
					time.Sleep(7 *1e7)
					robotgo.KeyTap("q")
					time.Sleep(24 *1e7)
				}
			}()
			go func() {
				for autoState.autoActive {
					time.Sleep(9 *1e7)
					robotgo.KeyTap("w")
					time.Sleep(28 *1e7)
				}
			}()
			go func() {
				for autoState.autoActive {
					time.Sleep(1 *1e7)
					robotgo.KeyTap("e")
					time.Sleep(399 *1e7)
				}
			}()
			go func() {
				for autoState.autoActive {
					time.Sleep(3 *1e7)
					robotgo.KeyTap("r")
					time.Sleep(247 *1e7)
				}
			}()
			go func() {
				for autoState.autoActive {
					time.Sleep(5 *1e7)
					robotgo.KeyTap("g")
					time.Sleep(195 *1e7)
				}
			}()
		}
	})
}*/

func makeSeq(init int) func() int{
	index := init
	return func() int {
		index++
		return index
	}
}

/**
 * 构建命令
 * go build -o autoKey.exe .\autoKeyPress.go
 */
```
配置文件conf.json
```json
[
  { "key": "q", "firstMoment": 1, "period": 500 },
  { "key": "w", "firstMoment": 9, "period": 320 },
  { "key": "e", "firstMoment": 7, "period": 29  },
  { "key": "r", "firstMoment": 3, "period": 300 },
  { "key": "g", "firstMoment": 5, "period": 200 }
]
```
