// Package aliases демонстрирует объявление именованных и функциональных типов.
package aliases

import (
	"fmt"
	"time"
)

// Message — именованный строковый тип для текста сообщения.
type Message string

// MessageType — именованный строковый тип для уровня/типа сообщения.
type MessageType string

type logger func(msg Message, msgType MessageType)

// LogWithCode — функциональный тип логгера, принимающего сообщение, его тип и числовой код.
type LogWithCode func(msg Message, msgType MessageType, code int)

func printLog() logger {
	var counter int = 1
	return func(msg Message, msgType MessageType) {
		var now time.Time = time.Now()
		fmt.Printf("#%v [%v][%v] %v\n", counter, now.Format(time.RFC3339), msgType, msg)
		counter++
	}
}

// Logger возвращает логгер LogWithCode с внутренним счётчиком сообщений.
func Logger() LogWithCode {
	var counter int = 1
	return func(msg Message, msgType MessageType, code int) {
		var now time.Time = time.Now()
		fmt.Printf("#%v [%v][%v] (%v) %v\n", counter, now.Format(time.RFC3339), msgType, code, msg)
		counter++
	}
}

// Test — точка входа для демонстрационных запусков пакета aliases.
func Test() {
	var log logger = printLog()
	fmt.Println("-== Aliases ==-")
	log("First message", "debug")
	time.Sleep(1 * time.Second)
	log("Second message", "error")
	time.Sleep(2 * time.Second)
	log("Third message", "warning")
}
