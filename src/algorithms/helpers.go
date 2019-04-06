package algorithms

var Comparisons int = 0
var Swaps int = 0

func getComparator(flag bool) func(int64, int64) bool {
	if flag {
		return func(a int64, b int64) bool {
			//fmt.Fprintln(os.Stderr, a, "<", b, "?")
			Comparisons++
			return a < b
		}
	}

	return func(a int64, b int64) bool {
		//fmt.Fprintln(os.Stderr, a, ">", b, "?")
		Comparisons++
		return a > b
	}
}

func swap(a *int64, b *int64) {
	//fmt.Fprintln(os.Stderr, "swapping", *a, "and", *b)
	Swaps++
	*a, *b = *b, *a
}
