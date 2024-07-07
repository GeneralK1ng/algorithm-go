package _733_MinNumOfPeopleTeach

import (
	"sort"
)

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	m := len(languages)
	// 用于记录有冲突的好友们各语言的计数
	count := make([]int, n+1)
	// 标记哪些用户有冲突
	flag := make([]bool, m)
	// 将languages转化为二维布尔数组方便查询各用户语言掌握情况
	l := make([][]bool, m)
	for i := range l {
		l[i] = make([]bool, n+1)
		for _, j := range languages[i] {
			l[i][j] = true
		}
	}
	// 比对各个好友冲突情况
	for _, f := range friendships {
		a, b := f[0]-1, f[1]-1
		conflict := true
		for i := 1; i <= n; i++ {
			if l[a][i] && l[b][i] {
				conflict = false
				break
			}
		}
		if conflict {
			flag[a], flag[b] = true, true
		}
	}
	// 遍历并计数各语言掌握数量, c用于计数一共有多少冲突用户
	c := 0
	for i, f := range flag {
		if !f {
			continue
		}
		c++
		for j := 1; j <= n; j++ {
			if l[i][j] {
				count[j]++
			}
		}
	}
	sort.Ints(count)
	// 找到冲突用户中掌握数量最多的语言，教会没有掌握这门语言的其他冲突用户即可
	return c - count[n]
}
