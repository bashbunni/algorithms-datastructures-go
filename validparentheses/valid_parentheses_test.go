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
		{"{[]}", true},
		{"{", false},
		{"{{", false},
		{"))", false},
	}
	for _, tcase := range tests {
		got := IsValidParentheses(tcase.input)
		if got != tcase.want {
			t.Errorf("got %v, want %v with %s", got, tcase.want, tcase.input)
		}
	}
}
