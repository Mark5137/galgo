package galgo

import (
	"fmt"
	"testing"
)

type Test struct {
	number int
	arr    []int
	result int
}

var tests = []Test{
	{
		number: 15,
		arr:    []int{1, 5, 10, 15, 20},
		result: 3,
	},
	{
		number: 25,
		arr:    []int{1, 5, 10, 15, 20},
		result: -1,
	},
}

func TestLinear(t *testing.T) {
	for _, test := range tests {
		result := Linear(test.arr, test.number)
		if result != test.result {
			t.Error(fmt.Sprintf("Expected %d, got ", test.result), result)
		}
	}
}

func TestBinary(t *testing.T) {
	for _, test := range tests {
		result := Binary(test.arr, test.number)
		if result != test.result {
			t.Error(fmt.Sprintf("Expected %d, got ", test.result), result)
		}
	}
}

func TestInterpolation(t *testing.T) {
	for _, test := range tests {
		result := Interpolation(test.arr, test.number)
		if result != test.result {
			t.Error(fmt.Sprintf("Expected %d, got ", test.result), result)
		}
	}
}
