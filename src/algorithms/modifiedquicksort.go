package algorithms

func modifiedPartition(numbers []int64, start int, end int, ascending bool) int {
	compare := getComparator(!ascending)

	possiblePivots := InsertionSort([]int64{numbers[start], numbers[(start+end)/2], numbers[end-1]}, true)

	x := possiblePivots[0]
	if len(possiblePivots) > 1 {
		x = possiblePivots[1]
	}

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

func modQuickSortInplace(numbers []int64, start int, end int, ascending bool) {
	if start < end {
		if end-start <= 16 {
			InsertionSort(numbers[start:end+1], ascending)
		} else {
			q := modifiedPartition(numbers, start, end, ascending)
			modQuickSortInplace(numbers, start, q-1, ascending)
			modQuickSortInplace(numbers, q+1, end, ascending)
		}
	}
}
func ModifiedQuickSort(numbers []int64, ascending bool) []int64 {
	modQuickSortInplace(numbers, 0, len(numbers)-1, ascending)
	return numbers
}
