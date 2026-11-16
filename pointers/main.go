package main

import (
	"fmt"

	"pointers/pkg/util"
)

func main() {
	var messages []string = []string{
		"Success message",
		"Error message",
		"Debug message",
		"Warning message",
	}

	var numbers []uint8 = []uint8{1, 2, 3, 4, 5}
	var pMsg *[]string = &messages

	fmt.Println(numbers)
	fmt.Println(messages)

	println("Изменение по значению")
	util.Changer(messages)
	fmt.Println(messages)

	// Для версии Golang 1.26.8: передача происходит по ссылке
	// [Success message Error message Debug message Warning message]
	// Изменение по значению
	// [Success message Error message Notice message Warning message]

	println("Изменение по ссылке")
	util.PChanger(pMsg)
	fmt.Println(messages)
}
