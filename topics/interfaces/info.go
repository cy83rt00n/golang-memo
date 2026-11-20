package interfaces

import (
	"fmt"
	"time"
)

type info struct {
	code int
	text string
}

func (msg *info) setCode(code int) {
	msg.code = code
}

func (msg *info) setMessage(text string) {
	msg.text = text
}

func (msg info) print() {
	now := time.Now()
	fmt.Printf("[%v][info] (%v) %v\n", now.Format(time.RFC3339), msg.code, msg.text)
}
