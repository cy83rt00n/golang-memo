// Package pointers демонстрирует передачу данных по значению и по указателю.
package pointers

import (
	"fmt"
)

// PChanger изменяет третий элемент слайса через указатель и печатает его длину.
func PChanger(pVictim *[]string) {
	(*pVictim)[2] = "Info message"
	fmt.Printf("Victim size is %d\n", len(*pVictim))
}

// Changer изменяет третий элемент слайса, переданного по значению, и печатает его длину.
func Changer(victim []string) {
	victim[2] = "Notice message"
	fmt.Printf("Victim size is %d\n", len(victim))
}

// Test — точка входа для демонстрационных запусков пакета pointers.
func Test() {
	messages := []string{
		"Success message",
		"Error message",
		"Debug message",
		"Warning message",
	}

	var numbers []uint8 = []uint8{1, 2, 3, 4, 5}
	var pMsg *[]string = &messages

	fmt.Println("-== Poointers ==-")
	fmt.Println(numbers)
	fmt.Println(messages)
	fmt.Println(pMsg)

	fmt.Println("Изменение по значению")
	Changer(messages)
	fmt.Println(messages)

	fmt.Println("Изменение по ссылке")
	PChanger(pMsg)
	fmt.Println(messages)
}
