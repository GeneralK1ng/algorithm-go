package ConstructLongestStr_2745

func longestString(x, y, z int) int {
	// dp[a][b][c][0] —— 'a' "AA", 'b' "BB", 'c' "AB"
	// dp[a][b][c][1] —— 'a' "AA", 'b' "BB", 'c' "AB"
	// dp[a][b][c][2] —— 'a' "AA", 'b' "BB", 'c' "AB"
	dp := make([][][][3]int, x+1)
	for i := range dp {
		dp[i] = make([][][3]int, y+1)
		for j := range dp[i] {
			dp[i][j] = make([][3]int, z+1)
		}
	}

	for aa := 0; aa <= x; aa++ {
		for bb := 0; bb <= y; bb++ {
			for ab := 0; ab <= z; ab++ {
				if aa > 0 {
					dp[aa][bb][ab][0] = 2 + max(dp[aa-1][bb][ab][1], dp[aa-1][bb][ab][2])
				}
				if bb > 0 {
					dp[aa][bb][ab][1] = 2 + dp[aa][bb-1][ab][0]
				}
				if ab > 0 {
					dp[aa][bb][ab][2] = 2 + max(dp[aa][bb][ab-1][1], dp[aa][bb][ab-1][2])
				}
			}
		}
	}

	maxLenEndingWithAA := dp[x][y][z][0]
	maxLenEndingWithBB := dp[x][y][z][1]
	maxLenEndingWithAB := dp[x][y][z][2]

	return max(maxLenEndingWithAA, maxLenEndingWithBB, maxLenEndingWithAB)
}
