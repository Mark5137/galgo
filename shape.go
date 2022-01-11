package galgo

type Shape interface {
	Area() int
}

type Square struct {
	width int
}

type Rectangle struct {
	width  int
	height int
}

func (s Square) Area() int {
	return s.width * s.width
}

func (s Rectangle) Area() int {
	return s.width * s.height
}
