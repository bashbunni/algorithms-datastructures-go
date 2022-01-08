package main

import (
	"algorithms/datastructures"
	mergetwosortedlists "algorithms/merge-two-sorted-lists"
	"algorithms/multiplicationtable"
	"algorithms/validparentheses"
	"fmt"
)

func main() {
	datastructures.BuildLinkedList()
	fmt.Println(multiplicationtable.MultiplicationTable(3))
	fmt.Println(validparentheses.IsValidParentheses("{[]}"))
	fmt.Println(mergetwosortedlists.MergeTwoLists(&mergetwosortedlists.ListNode{1, &mergetwosortedlists.ListNode{2, &mergetwosortedlists.ListNode{4, nil}}},
		&mergetwosortedlists.ListNode{1, &mergetwosortedlists.ListNode{3, &mergetwosortedlists.ListNode{4, nil}}}))
}
