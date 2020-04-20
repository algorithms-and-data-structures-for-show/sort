package merge

func Merge(left []int, right []int) []int {
	lst := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			lst = append(lst, left[0])
			left = left[1:]
		} else {
			lst = append(lst, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		lst = append(lst, left...)
	}
	if len(right) > 0 {
		lst = append(lst, right...)
	}
	return lst
}

func Sort(lst []int) []int {
	length := len(lst)
	if length >= 2 {
		mid := length / 2
		lst = Merge(Sort(lst[:mid]), Sort(lst[mid:]))
	}
	return lst
}
