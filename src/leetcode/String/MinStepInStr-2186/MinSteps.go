package MinStepInStr_2186

func minSteps(s string, t string) int {
	byteCountSlice := make([]int, 26)
	for i := 0; i < len(s); i++ {
		var idx = int(s[i] - 'a')
		byteCountSlice[idx]--
	}
	for j := 0; j < len(t); j++ {
		var idx = int(t[j] - 'a')
		byteCountSlice[idx]++
	}
	var count = 0
	for i := 0; i < len(byteCountSlice); i++ {
		if byteCountSlice[i] < 0 {
			count = count - byteCountSlice[i]
		} else {
			count = count + byteCountSlice[i]
		}
	}
	return count
}
