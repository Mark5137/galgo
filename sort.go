package galgo

func BubbleSort(collection []int) []int {
	length := len(collection)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if collection[j] > collection[j+1] {
				collection[j], collection[j+1] = collection[j+1], collection[j]
			}
		}
	}

	return collection
}

func InsertionSort(collection []int) []int {
	length := len(collection)
	for i := 1; i < length; i++ {
		j := i
		for j > 0 {
			if collection[j-1] > collection[j] {
				collection[j-1], collection[j] = collection[j], collection[j-1]
			}
			j--
		}
	}

	return collection
}

func SelectionSort(collection []int) []int {
	length := len(collection)
	for i := 0; i < length; i++ {
		x := i
		for j := i; j < length; j++ {
			if collection[j] < collection[x] {
				x = j
			}
		}
		collection[i], collection[x] = collection[x], collection[i]
	}

	return collection
}
