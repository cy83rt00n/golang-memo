package aliases

import (
	"fmt"
	"time"
)

type Message string
type MessageType string
type logger func(msg Message, msgType MessageType)
type LogWithCode func(msg Message, msgType MessageType, code int)

func printLog() logger {
	var counter int = 1
	return func(msg Message, msgType MessageType) {
		var now time.Time = time.Now()
		fmt.Printf("#%v [%v][%v] %v\n", counter, now.Format(time.RFC3339), msgType, msg)
		counter++
	}
}

func Logger() LogWithCode {
	var counter int = 1
	return func(msg Message, msgType MessageType, code int) {
		var now time.Time = time.Now()
		fmt.Printf("#%v [%v][%v] (%v) %v\n", counter, now.Format(time.RFC3339), msgType, code, msg)
		counter++
	}
}

func Test() {
	var log logger = printLog()
	fmt.Println("-== Aliases ==-")
	log("First message", "debug")
	time.Sleep(1 * time.Second)
	log("Second message", "error")
	time.Sleep(2 * time.Second)
	log("Third message", "warning")
}
