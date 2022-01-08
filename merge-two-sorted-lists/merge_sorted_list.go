package mergetwosortedlists

type ListNode struct {
	Val int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var newList *ListNode
	// init newList
	if list1 == nil && list2 != nil{
		return list2
	} else if list2 == nil && list1 != nil {
		return list1
	} else if list2 == nil && list1 == nil {
		return nil
	}
	if list1.Val < list2.Val {
		newList = &ListNode{list1.Val, nil}
		list1 = list1.Next
	} else {
		newList = &ListNode{list2.Val, nil}
		list2 = list2.Next
	}
	head := newList
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			// put list1 val before list2 val
			newList.Next = &ListNode{list1.Val, nil}
			newList = newList.Next
			list1 = list1.Next
		} else {
			newList.Next = &ListNode{list2.Val, nil}
			newList = newList.Next
			list2 = list2.Next
		}
	}
	if list1 == nil {
		newList.Next = list2
	} else if list2 == nil {
		newList.Next = list1
	}
	return head
}
