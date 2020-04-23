package quick

func Sort(lst []int) {
	quicksort(lst, 0, len(lst)-1)
}

func quicksort(lst []int, start int, end int) {
	if end == 0 {
		end = len(lst) - 1
	}
	i := start
	j := end
	x := lst[(start+end)/2]

	for i <= j {
		for lst[i] < x {
			i++
		}
		for lst[j] > x {
			j--
		}
		if i <= j {
			if lst[i] > lst[j] {
				lst[i], lst[j] = lst[j], lst[i]
			}
			i++
			j--
		}
	}
	if i < end {
		quicksort(lst, i, end)
	}
	if j > start {
		quicksort(lst, start, j)
	}
}
