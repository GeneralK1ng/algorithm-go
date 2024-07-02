package MinNumFrog_1419

func minNumberOfFrogs(croakOfFrogs string) int {
	counter := [5]int{}
	frogs := 0

	for _, c := range croakOfFrogs {
		switch c {
		case 'c':
			if counter[4] > 0 {
				counter[4]--
			} else {
				frogs++
			}
			counter[0]++
		case 'r':
			if counter[0] == 0 {
				return -1
			}
			counter[0]--
			counter[1]++
		case 'o':
			if counter[1] == 0 {
				return -1
			}
			counter[1]--
			counter[2]++
		case 'a':
			if counter[2] == 0 {
				return -1
			}
			counter[2]--
			counter[3]++
		case 'k':
			if counter[3] == 0 {
				return -1
			}
			counter[3]--
			counter[4]++
		default:
			return -1
		}
	}

	for i := 0; i < 4; i++ {
		if counter[i] != 0 {
			return -1
		}
	}

	return frogs
}
