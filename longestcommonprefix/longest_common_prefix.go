package longestcommonprefix

import (
	"fmt"
	"strings"
)

func LongestCommonPrefix(strs []string) string {
	// get shortest word -> the prefix will not be longer than the shortest word
	// can do vertical scanning -> it will short circuit once we've hit the max len of a word
	// short circuit: stop if len(strs) == 0
	min := getSmallestWord(strs)
	for _, str := range strs {
		fmt.Println(strings.Index(str, min))
		for strings.Index(str, min) != 0 && len(min) != 0 {
			min = min[:len(min)-1]
		}
		if len(min) == 0 {
			return ""
		}
	}
	return min
}

// time: O(N) where n == len(strs)
// space: O(1) ?
func getSmallestWord(strs []string) string {
	var min string
	for _, str := range(strs) {
		if len(str) < len(min) || min == "" {
			min = str
		}
	}
	return min
}
