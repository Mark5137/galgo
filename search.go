package galgo

func Linear(collection []int, number int) int {
	for i, value := range collection {
		if value == number {
			return i
		}
	}

	return -1
}

func Binary(collection []int, number int) int {
	a := 0
	b := len(collection) - 1

	for a <= b {
		i := (a + b) / 2

		if collection[i] == number {
			return i
		} else if collection[i] < number {
			a = i + 1
		} else if collection[i] > number {
			b = i - 1
		}
	}

	return -1
}

func Interpolation(collection []int, number int) int {
	a := 0
	b := len(collection) - 1

	for (a <= b) && (number >= collection[a]) && (number <= collection[b]) {
		x := a + (((b - a) / (collection[b] - collection[a])) * (number - collection[a]))

		if collection[x] == number {
			return x
		}

		if collection[x] < number {
			a = x + 1
		} else {
			b = x - 1
		}
	}

	return -1
}
