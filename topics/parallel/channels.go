package parallel

import (
	"fmt"
	"sync"
	"time"
)

// testPasses задаёт количество итераций, которое выполняет каждая job.
var testPasses = 10

// logger читает сообщения из канала ioBuffer и выводит их в стандартный
// поток вывода с временной меткой в формате RFC3339.
//
// logger читает канал через range, поэтому завершается автоматически,
// когда канал закрыт и все буферизованные сообщения прочитаны.
// По завершении вызывает wg.Done().
func logger(wg *sync.WaitGroup, ioBuffer chan string) {
	defer wg.Done()
	for msg := range ioBuffer {
		fmt.Printf("[%v][info] %v\n", time.Now().Format(time.RFC3339), msg)
	}
}

// job имитирует полезную работу: выполняет testPasses итераций,
// на каждой итерации отправляет в ioBuffer сообщение о прогрессе
// и спит delay секунд. По завершении отправляет финальное
// сообщение "Job #<id> done" и вызывает wg.Done().
//
// Параметр delay задаётся в секундах.
func job(id int, delay time.Duration, wg *sync.WaitGroup, ioBuffer chan string) {
	defer wg.Done()
	for pass := testPasses; pass > 0; pass-- {
		ioBuffer <- fmt.Sprintf("Job #%v in progress. Passes left %v", id, pass)
		time.Sleep(delay)
	}
	ioBuffer <- fmt.Sprintf("Job #%v done", id)
}

// TestChannels запускает три job с разными задержками и один logger,
// который читает их сообщения из общего канала.
//
// Синхронизация устроена так:
//   - jobsWg отслеживает завершение всех отправителей (job);
//   - после jobsWg.Wait() канал закрывается — это безопасно, так как
//     все отправители уже закончили;
//   - loggerWg.Wait() дожидается, пока logger дочитает канал до конца
//     и завершится по close.
//
// Канал создаётся локально, поэтому TestChannels можно вызывать
// повторно без риска отправки в уже закрытый канал.
func TestChannels() {
	var loggerWg, jobsWg sync.WaitGroup
	var ioBuffer chan string = make(chan string)

	loggerWg.Add(1)
	go logger(&loggerWg, ioBuffer)

	jobsWg.Add(3)
	go job(123, 2*time.Second, &jobsWg, ioBuffer)
	go job(124, 4*time.Second, &jobsWg, ioBuffer)
	go job(125, 1*time.Second, &jobsWg, ioBuffer)

	jobsWg.Wait()
	close(ioBuffer)
	loggerWg.Wait()
}
