package longestcommonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	t.Run("fl" , func(t *testing.T) {
		got := LongestCommonPrefix([]string{"flower", "flow", "flight"})
		want := "fl"

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})

	t.Run("leet" , func(t *testing.T) {
		got := LongestCommonPrefix([]string{"leet", "leets", "leetcode"})
		want := "leet"

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})

	t.Run("not found" , func(t *testing.T) {
		got := LongestCommonPrefix([]string{"leet", "raccoon", "skater"})
		want := ""

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})

}
