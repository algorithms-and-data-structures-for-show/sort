package shellsort

func Sort(lst []int) {
	for inc := len(lst) / 2; inc >= 1; inc /= 2 {
		for step := 0; step < inc; step++ {
			insertionSort(lst, step, inc)
		}
	}
}

func insertionSort(lst []int, start int, inc int) {
	for i := start; i < len(lst)-1; i += inc {
		j := min(i+inc, len(lst)-1)
		stone := lst[j]
		j -= inc
		for j >= 0 && lst[j] > stone {
			lst[j+inc] = lst[j]
			j -= inc
		}
		lst[j+inc] = stone
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
