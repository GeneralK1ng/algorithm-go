package AvaliableCapturesRook_999

func numRookCaptures(board [][]byte) int {
	var rookRow, rookCol int

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				rookRow, rookCol = i, j
				break
			}
		}
	}

	captures := 0

	for i := rookRow - 1; i >= 0; i-- {
		if board[i][rookCol] == 'B' {
			break
		}
		if board[i][rookCol] == 'p' {
			captures++
			break
		}
	}

	for i := rookRow + 1; i < 8; i++ {
		if board[i][rookCol] == 'B' {
			break
		}
		if board[i][rookCol] == 'p' {
			captures++
			break
		}
	}

	for j := rookCol - 1; j >= 0; j-- {
		if board[rookRow][j] == 'B' {
			break
		}
		if board[rookRow][j] == 'p' {
			captures++
			break
		}
	}

	for j := rookCol + 1; j < 8; j++ {
		if board[rookRow][j] == 'B' {
			break
		}
		if board[rookRow][j] == 'p' {
			captures++
			break
		}
	}

	return captures
}
