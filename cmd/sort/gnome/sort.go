package gnome

func sort(sliceToSort []int64) []int64 {
	i := 1
	j := 2
	for i < len(sliceToSort) {
		if sliceToSort[i-1] < sliceToSort[i] {
			i = j
			j++
		} else {
			sliceToSort[i-1], sliceToSort[i] = sliceToSort[i], sliceToSort[i-1]
			i--
			if i == 0 {
				i = j
				j++
			}
		}
	}
	return sliceToSort
}
