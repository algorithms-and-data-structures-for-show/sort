package bubble

func sort(sliceToSort []int64) []int64 {
	for i := 0; i < len(sliceToSort); i++ {
		for j := i + 1; j < len(sliceToSort); j++ {
			if sliceToSort[i] > sliceToSort[j] {
				sliceToSort[i], sliceToSort[j] = sliceToSort[j], sliceToSort[i]
			}
		}
	}
	return sliceToSort
}
