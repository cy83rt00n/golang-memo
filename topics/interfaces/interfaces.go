// Package interfaces демонстрирует определение и реализацию интерфейсов,
// а также проверку типов через type switch.
package interfaces

import (
	"fmt"
)

// Logable — интерфейс сущности, умеющей печатать себя и менять свой код и текст.
type Logable interface {
	setCode(int)
	setMessage(string)
	print()
}

type undefined struct{}

func log(msg Logable) {
	msg.print()
}

func checkType(variable interface{}) {
	switch variable.(type) {
	case debug:
		fmt.Printf("Variable of type 'debug': %v\n", variable)
	case info:
		fmt.Printf("Variable of type 'info': %v\n", variable)
	case Logable:
		fmt.Printf("Variable of type 'Logable' interface: %v\n", variable)
	default:
		fmt.Println("Variable of type 'undefined'")
	}
}

// Test — точка входа для демонстрационных запусков пакета interfaces.
func Test() {
	dbg := debug{}
	inf := info{}
	var edbg Logable = &debug{}

	pDbg := &dbg
	pInf := &inf

	fmt.Println("-== Interfaces ==-")
	dbg.setCode(100)
	dbg.setMessage("Debug message")
	inf.setCode(299)
	inf.setMessage("Info message")
	edbg.setCode(900)
	edbg.setMessage("Extended debug message")
	log(pDbg)
	log(pInf)
	log(edbg)
	checkType(dbg)
	checkType(edbg)
	checkType(undefined{})

}
