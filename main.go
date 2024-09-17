package main

import (
	"fmt"
	types "golearning/types"
)

func main() {
	list := types.NewLinkedList([]string{"first", "second", "third", "fourth", "fifth"})
	list.PrintValues()
	fmt.Println("_______________")
	if list.SearchAndReplace("third", "3rd") {
		list.PrintValues()
	}
	fmt.Println("_______________")
	if list.Delete("fourth") {
		list.PrintValues()
	}
	fmt.Println("_______________")
	list.Add("sixth")
	list.Add("seventh")
	list.PrintValues()
}
