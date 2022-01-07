package main

import (
	"algorithms/datastructures"
	"algorithms/multiplicationtable"
	"fmt"
)

func main() {
	datastructures.BuildLinkedList()
	fmt.Println(multiplicationtable.MultiplicationTable(3))
}
