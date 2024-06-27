package LexicographicallySmallestSubstring_2734

func smallestString(s string) string {
	arr := []byte(s)
	n := len(arr)
	replace := false
	for i := 0; i < n; i++ {
		if arr[i] == 'a' && replace {
			break
		}
		if arr[i] == 'a' {
			continue
		}
		arr[i] -= byte(1)
		replace = true
	}
	if !replace {
		arr[n-1] = 'z'
	}

	return string(arr)
}
