package validparentheses

import "testing"

func TestIsValidParentheses(t *testing.T) {
	type test struct {
		input string
		want bool
	}
	tests := []test{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
	}
	for _, tcase := range tests {
		got := IsValidParentheses(tcase.input)
		if got != tcase.want {
			t.Errorf("got %v, want %v", got, tcase.want)
		}
	}
}
