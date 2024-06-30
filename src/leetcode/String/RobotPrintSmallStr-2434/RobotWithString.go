package RobotPrintSmallStr_2434

import (
	"container/list"
	"strings"
)

func robotWithString(s string) string {
	n := len(s)
	cnt := make([]int, 26)
	for i := 0; i < n; i++ {
		cnt[s[i]-'a']++
	}

	var sb strings.Builder
	stack := list.New()
	minCh := 0 // 剩余最小字母（单调递增）

	for i := 0; i < n; i++ {
		ch := s[i]
		stack.PushBack(ch)
		cnt[ch-'a']--
		for minCh < 25 && cnt[minCh] == 0 {
			minCh++
		}
		for stack.Len() > 0 && int(stack.Back().Value.(byte)-'a') <= minCh {
			sb.WriteByte(stack.Remove(stack.Back()).(byte))
		}
	}

	return sb.String()
}
