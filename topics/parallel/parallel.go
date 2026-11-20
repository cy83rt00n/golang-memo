// Package parallel демонстрирует конкурентное выполнение задач в Go
// с использованием горутин, каналов и sync.WaitGroup.
package parallel

import (
	"fmt"
	"sync"
	"time"
)

// execute запускает одну из демонстрационных задач по её идентификатору.
//
// taskId == 1 — обратный отсчёт с задержкой 4 секунды (5 шагов);
// taskId == 2 — обратный отсчёт с задержкой 1000 миллисекунд (5 шагов).
//
// По завершении вызывает wg.Done() (через defer, чтобы гарантировать
// вызов даже при панике внутри задачи).
func execute(taskId int, wg *sync.WaitGroup) {
	defer wg.Done()
	switch taskId {
	case 1:
		secondsBacktick(4*time.Second, 5)
	case 2:
		millisecondsBacktick(1000*time.Millisecond, 5)
	}
}

// testExecute параллельно запускает две задачи через execute
// и дожидается их завершения с помощью sync.WaitGroup.
func testExecute() {
	var wg sync.WaitGroup

	wg.Add(2)

	go execute(1, &wg)
	go execute(2, &wg)

	wg.Wait()
}

// Test — точка входа для демонстрационных запусков пакета.
// Сейчас вызывает TestChannels; после её завершения печатает
// итоговое сообщение "All tasks done!".
func Test() {
	fmt.Println("-== Parallel ==-")
	// testExecute()
	TestChannels()
	fmt.Println("All tasks done!")
}
