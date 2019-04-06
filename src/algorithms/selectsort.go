package algorithms

func SelectSort(numbers []int64, ascending bool) []int64 {
	comparator := getComparator(ascending)
	length := len(numbers)
	for idx := range numbers {
		minIdx := idx
		for i := idx + 1; i < length; i++ {
			if comparator(numbers[i], numbers[minIdx]) {
				minIdx = i
			}
		}

		//swap
		swap(&numbers[idx], &numbers[minIdx])
	}

	return numbers
}
