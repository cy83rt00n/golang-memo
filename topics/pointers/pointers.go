package pointers

import (
	"fmt"
)

func PChanger(pVictim *[]string) {
	(*pVictim)[2] = "Info message"
	fmt.Printf("Victim size is %d\n", len(*pVictim))
}

func Changer(victim []string) {
	victim[2] = "Notice message"
	fmt.Printf("Victim size is %d\n", len(victim))
}

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
	println(pMsg)

	println("Изменение по значению")
	Changer(messages)
	fmt.Println(messages)

	println("Изменение по ссылке")
	PChanger(pMsg)
	fmt.Println(messages)
}
