package algorithms

type heap struct {
	elems []int64
	size  int
}

func (h *heap) heapify(i int, ascending bool) {
	var best int

	compare := getComparator(!ascending)

	leftChild := i*2 + 1
	rightChild := i*2 + 2

	if leftChild < h.size && compare(h.elems[leftChild], h.elems[i]) {
		best = leftChild
	} else {
		best = i
	}
	if rightChild < h.size && compare(h.elems[rightChild], h.elems[best]) {
		best = rightChild
	}

	if best != i {
		swap(&h.elems[i], &h.elems[best])
		h.heapify(best, ascending)
	}
}

func buildHeap(numbers []int64, ascending bool) *heap {
	h := new(heap)
	h.elems = numbers
	h.size = len(numbers)

	for i := (h.size / 2) - 1; i >= 0; i-- {
		h.heapify(i, ascending)
	}

	return h
}

func HeapSort(numbers []int64, ascending bool) []int64 {
	h := buildHeap(numbers, ascending)
	for i := h.size - 1; i > 0; i-- {
		swap(&h.elems[0], &h.elems[i])
		h.size = h.size - 1
		h.heapify(0, ascending)
	}

	return h.elems
}
