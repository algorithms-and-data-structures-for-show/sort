package cocktail

func sort(sliceToSort []int64) []int64 {
	left := 0
	right := len(sliceToSort) - 1
	for left <= right {
		for i := left; i < right; i++ {
			j := i + 1
			if sliceToSort[i] > sliceToSort[j] {
				sliceToSort[i], sliceToSort[j] = sliceToSort[j], sliceToSort[i]
			}
		}
		right--
		for i := right; i > left; i-- {
			j := i - 1
			if sliceToSort[i] < sliceToSort[j] {
				sliceToSort[i], sliceToSort[j] = sliceToSort[j], sliceToSort[i]
			}
		}
		left++
	}

	return sliceToSort
}
