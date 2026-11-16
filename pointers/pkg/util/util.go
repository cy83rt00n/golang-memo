package util

import "fmt"

/*
Проверяем автодокументацию

@param pVictim []string - указатель на массив с текстами сообщений
*/
func PChanger(pVictim *[]string) {
	(*pVictim)[2] = "Info message"
	fmt.Printf("Victim size is %d\n", len(*pVictim))
}

func Changer(victim []string) {
	victim[2] = "Notice message"
	fmt.Printf("Victim size is %d\n", len(victim))
}
