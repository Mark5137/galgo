package galgo

import "testing"

var (
	collection         = []int{5, 7, 2, 9, 3}
	expectedCollection = []int{2, 3, 5, 7, 9}
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestBubbleSort(t *testing.T) {
	result := BubbleSort(collection)

	if !Equal(result, expectedCollection) {
		t.Error("Expected sorted collection")
	}
}

func TestInsertionSort(t *testing.T) {
	result := InsertionSort(collection)

	if !Equal(result, expectedCollection) {
		t.Error("Expected sorted collection")
	}
}

func TestSelectionSort(t *testing.T) {
	result := SelectionSort(collection)

	if !Equal(result, expectedCollection) {
		t.Error("Expected sorted collection")
	}
}
