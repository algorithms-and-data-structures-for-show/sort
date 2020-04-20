package insertion

func Sort(sliceToSort []int) {
	for i := 1; i < len(sliceToSort); i++ {
		stone := sliceToSort[i]
		j := i - 1
		for j >= 0 && sliceToSort[j] > stone {
			sliceToSort[j+1] = sliceToSort[j]
			j--
		}
		sliceToSort[j+1] = stone
	}
}

func sortBadVariant(sliceToSort []int) []int {
	for i := 1; i < len(sliceToSort); i++ {
		for j := i; j >= 1 && sliceToSort[j-1] > sliceToSort[j]; j-- {
			sliceToSort[j], sliceToSort[j-1] = sliceToSort[j-1], sliceToSort[j]
		}
	}
	return sliceToSort
}
