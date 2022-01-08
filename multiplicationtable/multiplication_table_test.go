package multiplicationtable

import (
	"reflect"
	"testing"
)

func TestMultiplicationTable(t *testing.T) {
	got := MultiplicationTable(3)
	want := [][]int{[]int{1,2,3}, []int{2,4,6}, []int{3,6,9}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}