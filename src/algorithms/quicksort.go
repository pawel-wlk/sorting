package algorithms

func partition(numbers []int64, start int, end int, ascending bool) int {
	compare := getComparator(!ascending)

	x := numbers[end]
	i := start - 1
	for j := start; j < end; j++ {
		if !compare(numbers[j], x) {
			i++
			swap(&numbers[i], &numbers[j])
		}
	}

	swap(&numbers[i+1], &numbers[end])

	return i + 1
}

func quickSortInplace(numbers []int64, start int, end int, ascending bool) {
	if start < end {
		q := partition(numbers, start, end, ascending)
		quickSortInplace(numbers, start, q-1, ascending)
		quickSortInplace(numbers, q+1, end, ascending)
	}
}

func QuickSort(numbers []int64, ascending bool) []int64 {
	quickSortInplace(numbers, 0, len(numbers)-1, ascending)
	return numbers
}
