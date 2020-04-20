package merge

import (
	"reflect"
	"testing"
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
