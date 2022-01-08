package mergetwosortedlists

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	type test struct {
		list1 *ListNode
		list2 *ListNode
		want *ListNode
	}

	tests := []test{
		{&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}}},
	}
	for _, tc := range tests {
		got := MergeTwoLists(tc.list1, tc.list2)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("did not pass test")
		}
	}
}
