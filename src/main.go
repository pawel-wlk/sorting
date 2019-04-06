package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"time"

	"./algorithms"
)

var sortingTime float64

func sort(sortType string, ascending bool, numbers []int64) []int64 {
	switch sortType {
	case "select":
		return algorithms.SelectSort(numbers, ascending)
	case "insert":
		return algorithms.InsertionSort(numbers, ascending)
	case "quick":
		return algorithms.QuickSort(numbers, ascending)
	case "heap":
		return algorithms.HeapSort(numbers, ascending)
	case "mquick":
		return algorithms.ModifiedQuickSort(numbers, ascending)
	}

	return nil
}

func isSorted(numbers []int64, ascending bool) bool {
	for i := 1; i < len(numbers); i++ {
		if ascending && numbers[i] < numbers[i-1] {
			return false
		}
		if !ascending && numbers[i] > numbers[i-1] {
			return false
		}
	}

	return true
}

func analyzeSort(sortType string, ascending bool, numbers []int64) []int64 {
	algorithms.Comparisons = 0
	algorithms.Swaps = 0
	startTime := time.Now()
	output := sort(sortType, ascending, numbers)
	endTime := time.Now()

	sortingTime = endTime.Sub(startTime).Seconds()

	return output
}

func main() {
	// parse args
	sortTypePtr := flag.String("type", "select", "sorting algorithm")
	ascendingFlagPtr := flag.Bool("asc", false, "sort ascending, default")
	descendingFlagPtr := flag.Bool("desc", true, "sort descending")
	statFlagPtr := flag.Bool("stat", false, "generate statistics for sorting algorithms")
	flag.Parse()

	if !*statFlagPtr {
		// read stdio
		var numsCount int
		_, err := fmt.Scanln(&numsCount)

		if err != nil {
			fmt.Println("bad input")
			return
		}

		numbers := make([]int64, numsCount)
		for i := 0; i < numsCount; i++ {
			_, err := fmt.Scan(&numbers[i])
			if err != nil {
				fmt.Println("bad input")
				return
			}
		}

		ascending := *ascendingFlagPtr || *descendingFlagPtr
		output := analyzeSort(*sortTypePtr, ascending, numbers)
		// print stats
		fmt.Fprintln(os.Stderr, "Comparisons:", algorithms.Comparisons)
		fmt.Fprintln(os.Stderr, "Swaps:", algorithms.Swaps)
		fmt.Fprintln(os.Stderr, "Time:", sortingTime, "seconds")

		// check if numbers are sorted
		if isSorted(output, ascending) {
			fmt.Println("sorted")
		} else {
			fmt.Println("not sorted")
		}

		// print out numbers
		fmt.Println(output, len(output))

	} else {
		if len(flag.Args()) != 2 {
			return
		}

		path := flag.Args()[0]
		k, _ := strconv.Atoi(flag.Args()[1])

		file, err := os.Create(path)

		if err != nil {
			fmt.Println("file error")
			return
		}

		for n := 100; n <= 10000; n += 100 {
			numbers := make([]int64, n)
			for i := 0; i < k; i++ {
				for j := 0; j < n; j++ {
					randNum, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
					numbers[j] = randNum.Int64()
				}

				analyzeSort("select", true, append([]int64(nil), numbers...))
				fmt.Fprintf(file, "%s, %d, %d, %d, %f\n", "select", n, algorithms.Comparisons, algorithms.Swaps, sortingTime)

				analyzeSort("insert", true, append([]int64(nil), numbers...))
				fmt.Fprintf(file, "%s, %d, %d, %d, %f\n", "insert", n, algorithms.Comparisons, algorithms.Swaps, sortingTime)

				analyzeSort("heap", true, append([]int64(nil), numbers...))
				fmt.Fprintf(file, "%s, %d, %d, %d, %f\n", "heap", n, algorithms.Comparisons, algorithms.Swaps, sortingTime)

				analyzeSort("quick", true, append([]int64(nil), numbers...))
				fmt.Fprintf(file, "%s, %d, %d, %d, %f\n", "quick", n, algorithms.Comparisons, algorithms.Swaps, sortingTime)

				analyzeSort("mquick", true, append([]int64(nil), numbers...))
				fmt.Fprintf(file, "%s, %d, %d, %d, %f\n", "mquick", n, algorithms.Comparisons, algorithms.Swaps, sortingTime)
			}
		}
	}

}
