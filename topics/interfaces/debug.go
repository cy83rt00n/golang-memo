package interfaces

import (
	"fmt"
	"time"
)

type debug struct {
	code int
	text string
}

func (msg *debug) setCode(code int) {
	msg.code = code
}

func (msg *debug) setMessage(text string) {
	msg.text = text
}

func (msg debug) print() {
	now := time.Now()
	fmt.Printf("[%v][debug] (%v) %v\n", now.Format(time.RFC3339), msg.code, msg.text)
}
