package timsort

import (
	"fmt"
	"github.com/algorithms-and-data-structures-for-show/sort/cmd/sort/utils"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestFirst(t *testing.T) {
	sliceToSort := []int{5, 22, 4, 6, 7, 8, 10, 17, 5}
	sliceSorted := Sort(sliceToSort)

	sclieExpected := []int{4, 5, 5, 6, 7, 8, 10, 17, 22}
	if !reflect.DeepEqual(sliceSorted, sclieExpected) {
		t.Error("Expected [4, 5, 5, 6, 7, 8, 10, 17, 22], got ", sliceSorted)
	}
}

func TestSecond(t *testing.T) {
	sliceToSort := []int{100, 5, 22, 4, 6, 7, 8, 10, 17, 1}
	sliceSorted := Sort(sliceToSort)

	sclieExpected := []int{1, 4, 5, 6, 7, 8, 10, 17, 22, 100}
	if !reflect.DeepEqual(sliceSorted, sclieExpected) {
		t.Error("Expected [1, 4, 5, 6, 7, 8, 10, 17, 22, 100], got ", sliceSorted)
	}
}

func TestThird(t *testing.T) {
	sliceToSort := utils.GenerateRandomSlice(100_000)
	var sliceToSort1 []int
	sliceToSort1 = append(sliceToSort1, sliceToSort...)

	startQuick := time.Now()
	sort.Slice(sliceToSort, func(i, j int) bool { return sliceToSort[i] < sliceToSort[j] })
	fmt.Printf("%s took %v\n", "Quick", time.Since(startQuick))

	startTim := time.Now()
	sliceSorted := Sort(sliceToSort1)
	fmt.Printf("%s took %v\n", "Tim", time.Since(startTim))

	if !reflect.DeepEqual(sliceSorted, sliceToSort) {
		t.Error("sliceSorted not equal  sliceToSort", sliceToSort, sliceSorted)
	}
}
