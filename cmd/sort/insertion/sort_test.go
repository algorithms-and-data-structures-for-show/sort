package insertion

import (
	"fmt"
	"github.com/algorithms-and-data-structures-for-show/sort/cmd/sort/utils"
	"reflect"
	"testing"
	"time"
)

func TestFirst(t *testing.T) {
	sliceToSort := []int{5, 22, 4, 6, 7, 8, 10, 17, 5}
	sliceSorted := sortBadVariant(sliceToSort)

	sclieExpected := []int{4, 5, 5, 6, 7, 8, 10, 17, 22}
	if !reflect.DeepEqual(sliceSorted, sclieExpected) {
		t.Error("Expected [4, 5, 5, 6, 7, 8, 10, 17, 22], got ", sliceSorted)
	}
}

func TestSecond(t *testing.T) {
	sliceToSort := []int{100, 5, 22, 4, 6, 7, 8, 10, 17, 1}
	sliceSorted := sortBadVariant(sliceToSort)

	sclieExpected := []int{1, 4, 5, 6, 7, 8, 10, 17, 22, 100}
	if !reflect.DeepEqual(sliceSorted, sclieExpected) {
		t.Error("Expected [1, 4, 5, 6, 7, 8, 10, 17, 22, 100], got ", sliceSorted)
	}
}

func TestThird(t *testing.T) {
	sliceToSort := utils.GenerateRandomSlice(10_000)
	var sliceToSort1 []int
	sliceToSort1 = append(sliceToSort1, sliceToSort...)

	startSort1 := time.Now()
	Sort(sliceToSort)
	fmt.Printf("%s took %v\n", "Sort", time.Since(startSort1))

	startWithSwap := time.Now()
	sliceSorted := sortBadVariant(sliceToSort1)
	fmt.Printf("%s took %v\n", "sortBadVariant ", time.Since(startWithSwap))

	if !reflect.DeepEqual(sliceSorted, sliceToSort) {
		t.Error("sliceSorted not equal  sliceToSort", sliceToSort, sliceSorted)
	}
}

func TestFourth(t *testing.T) {
	sliceToSort := utils.GenerateRandomSlice(100)
	var sliceToSort1 []int
	sliceToSort1 = append(sliceToSort1, sliceToSort...)

	startSort1 := time.Now()
	Sort(sliceToSort)
	fmt.Printf("%s took %v\n", "Sort", time.Since(startSort1))

	startWithSwap := time.Now()
	sliceSorted := sortBadVariant(sliceToSort1)
	fmt.Printf("%s took %v\n", "sortBadVariant ", time.Since(startWithSwap))

	if !reflect.DeepEqual(sliceSorted, sliceToSort) {
		t.Error("sliceSorted not equal  sliceToSort", sliceToSort, sliceSorted)
	}
}

//func TestFive(t *testing.T) {
//	sliceToSort := utils.GenerateRandomSlice(100_000)
//	var sliceToSort1 []int
//	sliceToSort1 = append(sliceToSort1, sliceToSort...)
//
//	startSort1 := time.Now()
//	Sort(sliceToSort)
//	fmt.Printf("%s took %v\n", "Sort", time.Since(startSort1))
//
//	startWithSwap := time.Now()
//	sliceSorted := sortBadVariant(sliceToSort1)
//	fmt.Printf("%s took %v\n", "sortBadVariant ", time.Since(startWithSwap))
//
//	if !reflect.DeepEqual(sliceSorted, sliceToSort) {
//		t.Error("sliceSorted not equal  sliceToSort", sliceToSort, sliceSorted)
//	}
//}

func TestSix(t *testing.T) {
	sliceToSort := utils.GenerateOrderedSlice(100_000)
	var sliceToSort1 []int
	sliceToSort1 = append(sliceToSort1, sliceToSort...)

	startSort1 := time.Now()
	Sort(sliceToSort)
	fmt.Printf("%s took %v\n", "Sort", time.Since(startSort1))

	startWithSwap := time.Now()
	sliceSorted := sortBadVariant(sliceToSort1)
	fmt.Printf("%s took %v\n", "sortBadVariant ", time.Since(startWithSwap))

	if !reflect.DeepEqual(sliceSorted, sliceToSort) {
		t.Error("sliceSorted not equal  sliceToSort", sliceToSort, sliceSorted)
	}
}

func TestSeven(t *testing.T) {
	sliceToSort := utils.GenerateOrderedSlice(1_000_000)
	var sliceToSort1 []int
	sliceToSort1 = append(sliceToSort1, sliceToSort...)

	startSort1 := time.Now()
	Sort(sliceToSort)
	fmt.Printf("%s took %v\n", "Sort", time.Since(startSort1))

	startWithSwap := time.Now()
	sliceSorted := sortBadVariant(sliceToSort1)
	fmt.Printf("%s took %v\n", "sortBadVariant ", time.Since(startWithSwap))

	if !reflect.DeepEqual(sliceSorted, sliceToSort) {
		t.Error("sliceSorted not equal  sliceToSort", sliceToSort, sliceSorted)
	}
}
