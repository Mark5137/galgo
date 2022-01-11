package galgo

import "testing"

func TestRectangle(t *testing.T) {
	rectangle := Rectangle{width: 3, height: 2}
	result := rectangle.Area()

	if result != 6 {
		t.Error("Expected 6, got ", result)
	}
}

func TestSquare(t *testing.T) {
	square := Square{width: 3}
	result := square.Area()

	if result != 9 {
		t.Error("Expected 9, got ", result)
	}
}
