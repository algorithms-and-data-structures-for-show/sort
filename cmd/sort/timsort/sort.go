package timsort

import (
	"github.com/algorithms-and-data-structures-for-show/sort/cmd/sort/insertion"
	"github.com/algorithms-and-data-structures-for-show/sort/cmd/sort/merge"
)

func Sort(lst []int) []int {
	var runs [][]int
	length := len(lst)
	newRun := []int{lst[0]}

	lengthN := 0
	for i := 1; true; i++ {
		if i == length-1 {
			newRun = append(newRun, lst[i])
			insertion.Sort(newRun)
			runs = append(runs, newRun)
			break
		}
		if lst[i] < lst[i-1] && lengthN > 200 {
			lengthN = 0
			insertion.Sort(newRun)
			runs = append(runs, newRun)
			newRun = []int{lst[i]}
		} else {
			newRun = append(newRun, lst[i])
		}
		lengthN++
	}

	var sortedArray []int
	for _, run := range runs {
		sortedArray = merge.Merge(sortedArray, run)
	}

	return sortedArray
}
