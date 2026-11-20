package slices

import (
	"fmt"
	"lessons/topics/structs"
	"reflect"
	"sort"
)

var food []string = []string{"Apple", "Banana", "Chocolate", "Strawberry", "Lemon", "Orange", "Fish"}

func Test() {
	fmt.Println("-==Slices==-")

	fmt.Println("-==Slice from lessons==-")
	fmt.Println(structs.MsgSlice[2:5])

	fmt.Println("-==Copy==-")
	var destSlice = make([]string, 12)
	fmt.Println(destSlice)
	copy(destSlice, food)
	fmt.Println(destSlice)

	fmt.Println("-==Sorting==-")
	fmt.Println(food)
	sort.Strings(food)
	fmt.Println(food)

	fmt.Println("-==Comparison==-")
	reverseSlice := make([]string, 7)
	copy(reverseSlice, food)
	fmt.Println(reverseSlice)
	fmt.Println(reflect.DeepEqual(food, reverseSlice))
	sort.Sort(sort.Reverse(sort.StringSlice(reverseSlice)))
	fmt.Println(reverseSlice)
	fmt.Println(reflect.DeepEqual(food, reverseSlice))
}
