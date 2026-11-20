package main

import (
	"lessons/topics/aliases"
	"lessons/topics/basics"
	"lessons/topics/interfaces"
	"lessons/topics/maps"
	"lessons/topics/methods"
	"lessons/topics/parallel"
	"lessons/topics/pointers"
	"lessons/topics/slices"
	"lessons/topics/structs"
)

func main() {
	basics.Test()
	aliases.Test()
	structs.Test()
	methods.Test()
	slices.Test()
	maps.Test()
	interfaces.Test()
	parallel.Test()
	pointers.Test()
}
