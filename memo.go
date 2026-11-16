package main

import "fmt"

// Объявление нового типа для лучшей читаемости
type errorMessages []string
type message struct {
	isError   bool
	isSuccess bool
	code      int
	message   string
}
type msgPool []message

var errors errorMessages = errorMessages{
	"Service unavailable",
	"Invalid request",
}

var messages msgPool = msgPool{
	{isError: false, isSuccess: true, code: 101, message: "Completed"},
	{isError: true, isSuccess: false, code: 201, message: "Method not allowed"},
	{isError: false, isSuccess: false, code: 301, message: "Debug info"},
}

// Возврат значений
func ff() int {
	return 12
}

// Рекурсия
func fd(i, j int) int {
	if i == 2 {
		i = fd(i+1, j)
		fmt.Printf("Recurrent i is %d\n", i)
	}
	i += j
	return i
}

// Переменное число параметров
func fc(numbers ...int) {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("sum = ", sum)
}

// Возвращение именованных параметров
func fb(i, j int) (g, k int) {
	g = i + j
	k = i - j
	return
}

// Функция как параметр
func fp(a, b int, handler func(int, int) int) (response int) {
	response = handler(a, b)
	return
}

// Замыкание. Запоминает свое окружение и возвращает функцию. Wrapper сообщения об ошибках.
func errLogger(message string) func() {
	var n int
	var innerF func() = func() {
		n += 1
		fmt.Printf("Error '%s' #%d\n", message, n)
	}

	return innerF
}

func logger(msg message) func() {
	var counter int
	var trueFn func() = func() {
		counter++
		var prefix string
		if msg.isError {
			prefix = fmt.Sprintf("Error #%d:", counter)
		} else if msg.isSuccess {
			prefix = fmt.Sprintf("Success #%d:", counter)
		} else {
			prefix = fmt.Sprintf("Debug #%d:", counter)
		}
		fmt.Printf("%s %s\n", prefix, msg.message)
	}

	return trueFn
}

func main() {
	var arr [3][3][3]int
	var a int

	arr = [3][3][3]int{
		{
			{1, 3, 5},
			{2, 4, 6},
			{5, 7, 9},
		},
		{
			{1, 3, 5},
			{2, 4, 6},
			{5, 7, 9},
		},
		{
			{1, 3, 5},
			{2, 4, 6},
			{5, 7, 9},
		},
	}
	a = 4

	switch a + ff() {
	case 24:
		fmt.Println("24")
	case 13:
		fmt.Print("13")
	default:
		fmt.Print("default")
	}
	fmt.Println(arr)

	fmt.Printf("Recurrent func returned %d\n", fd(2, 9))

	fc([]int{2, 5, 7}...)

	var fa func(numbers ...int) = fc

	fa(23, 54, 88)

	// Анонимная функция
	fmt.Printf("anonymous result is %d\n", fp(12, 56, func(a, b int) int { return a + b }))

	// Проверка wapper'а сообщений об ошибках
	var serverUnavailable func() = errLogger(errors[0])
	invalidRequest := errLogger(errors[1])
	writeErrorToLog := logger(messages[1])
	writeSuccessToLog := logger(messages[0])
	writeDebugToLog := logger(messages[2])
	serverUnavailable()
	serverUnavailable()
	invalidRequest()
	serverUnavailable()
	invalidRequest()
	serverUnavailable()

	writeDebugToLog()
	writeErrorToLog()
	writeSuccessToLog()
	writeDebugToLog()
	writeDebugToLog()
	writeSuccessToLog()
	writeSuccessToLog()
	writeErrorToLog()
}
