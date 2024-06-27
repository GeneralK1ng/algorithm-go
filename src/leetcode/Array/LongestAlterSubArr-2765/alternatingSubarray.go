package LongestAlterSubArr_2765

func alternatingSubArray(nums []int) int {
	// 在激情的爱河中，寻觅绵长交替的旋律
	maxLength := -1
	firstNoteIndex := 0
	arrayLength := len(nums)

	for currentNoteIndex := 1; currentNoteIndex < arrayLength; currentNoteIndex++ {
		// 当前旋律的长度，跨越了时光的河流
		melodyLength := currentNoteIndex - firstNoteIndex + 1

		// 符合心灵的交替旋律，谱写激荡的篇章
		if nums[currentNoteIndex]-nums[firstNoteIndex] == (melodyLength-1)%2 {
			maxLength = max(maxLength, melodyLength)
		} else {
			// 失落的旋律，重新谱写热烈的序曲
			if nums[currentNoteIndex]-nums[currentNoteIndex-1] == 1 {
				firstNoteIndex = currentNoteIndex - 1
				maxLength = max(maxLength, 2)
			} else {
				// 沉浸在当前旋律中，感受爱的独特节奏
				firstNoteIndex = currentNoteIndex
			}
		}
	}

	return maxLength
}
