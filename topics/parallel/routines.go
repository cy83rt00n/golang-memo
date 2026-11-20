package parallel

import (
	"fmt"
	"time"
)

// secondsBacktick выполняет обратный отсчёт от counterStart до 1
// с задержкой delay между шагами и печатает каждый шаг с меткой "Timer #1".
//
// delay передаётся как готовый time.Duration (например, 4*time.Second),
// поэтому функция не интерпретирует число и не домножает его на единицы
// времени — она просто спит ровно столько, сколько ей сказали.
func secondsBacktick(delay time.Duration, counterStart int) {
	for d := counterStart; d > 0; d-- {
		fmt.Printf("Timer #1 %v\n", d)
		time.Sleep(delay)
	}
}

// millisecondsBacktick выполняет обратный отсчёт от counterStart до 1
// с задержкой delay между шагами и печатает каждый шаг с меткой "Timer #2".
//
// Как и secondsBacktick, принимает готовый time.Duration
// (например, 1000*time.Millisecond) и не домножает его на единицы времени.
func millisecondsBacktick(delay time.Duration, counterStart int) {
	for d := counterStart; d > 0; d-- {
		fmt.Printf("Timer #2 %v\n", d)
		time.Sleep(delay)
	}
}
