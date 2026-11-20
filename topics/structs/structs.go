// Package structs демонстрирует объявление структур и их использование
// совместно с функциональными типами из пакета aliases.
package structs

import (
	"fmt"
	"lessons/topics/aliases"
	"time"
)

type msg struct {
	isError   bool
	isSuccess bool
	code      int
	message   aliases.Message
}
type msgPool []msg

var messages msgPool = msgPool{
	{isError: false, isSuccess: true, code: 101, message: "Completed"},
	{isError: true, isSuccess: false, code: 201, message: "Method not allowed"},
	{isError: false, isSuccess: false, code: 301, message: "Debug info"},
}

// MsgSlice — пример экспортируемого слайса строк, используемый в пакете slices.
var MsgSlice []string = []string{
	"Message #1",
	"Message #2",
	"Message #3",
	"Message #4",
	"Message #5",
	"Message #6",
	"Message #7",
	"Message #8",
}

// Test — точка входа для демонстрационных запусков пакета structs.
func Test() {
	var logger aliases.LogWithCode = aliases.Logger()

	fmt.Println("-== Structs ==-")
	for _, m := range messages {
		var mtype aliases.MessageType
		switch {
		case m.isError:
			mtype = "error"
		case m.isSuccess:
			mtype = "success"
		default:
			mtype = "debug"
		}
		logger(m.message, mtype, m.code)
	}
	time.Sleep(2 * time.Second)
}
