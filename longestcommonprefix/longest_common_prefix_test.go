package longestcommonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	got := LongestCommonPrefix([]string{"flower", "flow", "flight"})
	want := "fl"

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
