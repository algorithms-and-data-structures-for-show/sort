package quick

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
	Sort(sliceToSort)

	sclieExpected := []int{4, 5, 5, 6, 7, 8, 10, 17, 22}
	if !reflect.DeepEqual(sliceToSort, sclieExpected) {
		t.Error("Expected [4, 5, 5, 6, 7, 8, 10, 17, 22], got ", sliceToSort)
	}
}

func TestSecond(t *testing.T) {
	sliceToSort := []int{100, 5, 22, 4, 6, 7, 8, 10, 17, 1}
	Sort(sliceToSort)

	sclieExpected := []int{1, 4, 5, 6, 7, 8, 10, 17, 22, 100}
	if !reflect.DeepEqual(sliceToSort, sclieExpected) {
		t.Error("Expected [1, 4, 5, 6, 7, 8, 10, 17, 22, 100], got ", sliceToSort)
	}
}

func TestThird(t *testing.T) {
	sliceToSort := utils.GenerateRandomSlice(1_000)
	var sliceToSort1 []int
	sliceToSort1 = append(sliceToSort1, sliceToSort...)

	startQuick := time.Now()
	sort.Slice(sliceToSort, func(i, j int) bool { return sliceToSort[i] < sliceToSort[j] })
	fmt.Printf("%s took %v\n", "Quick", time.Since(startQuick))

	startTim := time.Now()
	Sort(sliceToSort1)
	fmt.Printf("%s took %v\n", "Quick our", time.Since(startTim))

	if !reflect.DeepEqual(sliceToSort1, sliceToSort) {
		t.Error("sliceSorted not equal  sliceToSort", sliceToSort, sliceToSort1)
	}
}
func TestFourth(t *testing.T) {
	sliceToSort := utils.GenerateRandomSlice(10_000)
	var sliceToSort1 []int
	sliceToSort1 = append(sliceToSort1, sliceToSort...)

	startQuick := time.Now()
	sort.Slice(sliceToSort, func(i, j int) bool { return sliceToSort[i] < sliceToSort[j] })
	fmt.Printf("%s took %v\n", "Quick", time.Since(startQuick))

	startTim := time.Now()
	Sort(sliceToSort1)
	fmt.Printf("%s took %v\n", "Quick our", time.Since(startTim))

	if !reflect.DeepEqual(sliceToSort1, sliceToSort) {
		t.Error("sliceSorted not equal  sliceToSort", sliceToSort, sliceToSort1)
	}
}

func TestFifth(t *testing.T) {
	sliceToSort := utils.GenerateRandomSlice(100_000)
	var sliceToSort1 []int
	sliceToSort1 = append(sliceToSort1, sliceToSort...)

	startQuick := time.Now()
	sort.Slice(sliceToSort, func(i, j int) bool { return sliceToSort[i] < sliceToSort[j] })
	fmt.Printf("%s took %v\n", "Quick", time.Since(startQuick))

	startTim := time.Now()
	Sort(sliceToSort1)
	fmt.Printf("%s took %v\n", "Quick our", time.Since(startTim))

	if !reflect.DeepEqual(sliceToSort1, sliceToSort) {
		t.Error("sliceSorted not equal  sliceToSort", sliceToSort, sliceToSort1)
	}
}

func TestSixth(t *testing.T) {
	sliceToSort := utils.GenerateRandomSlice(1_000_000)
	var sliceToSort1 []int
	sliceToSort1 = append(sliceToSort1, sliceToSort...)

	startQuick := time.Now()
	sort.Slice(sliceToSort, func(i, j int) bool { return sliceToSort[i] < sliceToSort[j] })
	fmt.Printf("%s took %v\n", "Quick", time.Since(startQuick))

	startTim := time.Now()
	Sort(sliceToSort1)
	fmt.Printf("%s took %v\n", "Quick our", time.Since(startTim))

	if !reflect.DeepEqual(sliceToSort1, sliceToSort) {
		t.Error("sliceSorted not equal  sliceToSort", sliceToSort, sliceToSort1)
	}
}

func TestSeventh(t *testing.T) {
	sliceToSort := utils.GenerateOrderedSlice(1_000_000)
	var sliceToSort1 []int
	sliceToSort1 = append(sliceToSort1, sliceToSort...)

	startQuick := time.Now()
	sort.Slice(sliceToSort, func(i, j int) bool { return sliceToSort[i] < sliceToSort[j] })
	fmt.Printf("%s took %v\n", "Quick", time.Since(startQuick))

	startTim := time.Now()
	Sort(sliceToSort1)
	fmt.Printf("%s took %v\n", "Quick our", time.Since(startTim))

	if !reflect.DeepEqual(sliceToSort1, sliceToSort) {
		t.Error("sliceSorted not equal  sliceToSort", sliceToSort, sliceToSort1)
	}
}
