package maps

import (
	"fmt"
	"time"
)

type message struct {
	code      uint
	text      string
	isSuccess bool
	isError   bool
}

var logMessages map[int]string = map[int]string{
	200: "Operation successful",
	404: "Undefined operation",
	503: "Critical error",
	100: "Debug message",
}

var logLevels map[string]message = map[string]message{
	"info":    {code: 200, text: "Successful operation", isSuccess: true, isError: false},
	"debug":   {code: 201, text: "Debug message", isSuccess: false, isError: false},
	"warning": {code: 202, text: "Warning message", isSuccess: false, isError: false},
	"error":   {code: 203, text: "Error message", isSuccess: false, isError: true},
}

func Test() {
	fmt.Println("-== Maps ==-")
	now := time.Now()
	for level, message := range logLevels {
		fmt.Printf("[%v][%v] (%v) %v\n", now.Format(time.RFC3339), level, message.code, message.text)
	}
}
