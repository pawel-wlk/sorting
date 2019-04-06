package algorithms

func InsertionSort(numbers []int64, ascending bool) []int64 {
	compare := getComparator(!ascending)

	for i := 1; i < len(numbers); i++ {
		key := numbers[i]
		j := i - 1
		for j >= 0 && compare(numbers[j], key) {
			//fmt.Fprintln(os.Stderr, "moving", numbers[j+1], " one position back")
			numbers[j+1] = numbers[j]
			Swaps++
			j--
		}
		numbers[j+1] = key
		//fmt.Fprintln(os.Stderr, "setting", j+1, "index to", key)
		Swaps++

	}
	return numbers
}
