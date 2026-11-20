package methods

import (
	"fmt"
	"time"
)

type message struct {
	code    int
	message string
	level   string
}

type messageType struct {
	info    string
	debug   string
	warning string
	error   string
}

var levels messageType = messageType{
	info:    "success",
	debug:   "debug",
	warning: "warning",
	error:   "error",
}

var messages []message = []message{
	{code: 200, message: "Operation finished successfully", level: levels.info},
	{code: 100, message: "Log messages test", level: levels.debug},
	{code: 300, message: "Not critical error. Keep an eye on it", level: levels.warning},
	{code: 500, message: "Critical error", level: levels.error},
}

func (m message) print() {
	now := time.Now()
	fmt.Printf("[%v][%v] (%v) %v\n", now.Format(time.RFC3339), m.level, m.code, m.message)
}

func (m *message) setLevel(lv string, code int) {
	m.level = lv
	if code >= 100 {
		m.code = code
	}
}

func Test() {
	fmt.Println("-== Methods ==-")
	for _, m := range messages {
		m.print()
	}

	pm := &messages[0]
	pm.setLevel(levels.debug, 0)

	fmt.Println("-==Set level check==-")
	for _, m := range messages {
		m.print()
	}
}
