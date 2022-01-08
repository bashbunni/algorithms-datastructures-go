package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type test struct{
		input []int
		target int
		want []int
	}
	tests := []test{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3,2,4}, 6, []int{1, 2}},
		{[]int{3,3}, 6, []int{0, 1}},
	}
	for _, tcase := range tests {
		got := twoSum(tcase.input, tcase.target)
		if !reflect.DeepEqual(got, tcase.want) {
			t.Errorf("got %q, want %q", got, tcase.want)
		}
	}
}
