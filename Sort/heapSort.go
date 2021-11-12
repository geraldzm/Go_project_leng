package Sort

import (
	"strconv"
	"time"
)

type HeapSort struct {
	Array *[]int
	ch    chan Item
}

func (h *HeapSort) Init() {
	h.ch = make(chan Item)
	go h.sort() //GORRUTINA
}

func (h *HeapSort) GetArray() *[]int {
	return h.Array
}

func (h *HeapSort) Sort() (Item, bool) {
	itm, more := <-h.ch
	return itm, more
}

//Código adaptado de https://golangbyexample.com/heapsort-in-golang/
func (h *HeapSort) sort() {

	t0 := time.Now().UnixNano() / int64(time.Millisecond)

	// arr := *h.Array
	len := len(*h.Array)

	comparisons := 0
	swaps := 0
	iterations := 1

	h.buildMaxHeap(len, &comparisons, &swaps, &iterations)
	for i := len - 1; i > 0; i-- {
		// Move current root to end
		h.swap(0, i)
		h.ch <- Item{IndexFrom: i, IndexTo: 0}
		swaps++
		h.downHeapify(0, i, &comparisons, &swaps, &iterations)
		iterations++
	}

	t1 := time.Now().UnixNano() / int64(time.Millisecond)

	h.ch <- Item{
		Finished:   true,
		TimeEnd:    strconv.FormatInt(t0, 10),
		TimeStart:  strconv.FormatInt(t1, 10),
		TotalTime:  strconv.FormatInt(t1-t0, 10),
		TotalComp:  comparisons,
		TotalIter:  iterations,
		TotalSwaps: swaps,
	}

	close(h.ch)
}

func (h *HeapSort) leftchildIndex(index int) int {
	return 2*index + 1
}

func (h *HeapSort) rightchildIndex(index int) int {
	return 2*index + 2
}

func (h *HeapSort) swap(first, second int) {
	temp := (*h.Array)[first]
	(*h.Array)[first] = (*h.Array)[second]
	(*h.Array)[second] = temp
}

func (h *HeapSort) leaf(index int, size int, comp *int) bool {
	*comp++
	*comp++
	if index >= (size/2) && index <= size {
		return true
	}
	return false
}

func (h *HeapSort) downHeapify(current int, size int, comp *int, swaps *int, iter *int) {
	*iter++
	if h.leaf(current, size, comp) {
		return
	}
	biggest := current
	leftChildIndex := h.leftchildIndex(current)
	rightRightIndex := h.rightchildIndex(current)
	*comp += 2
	if leftChildIndex < size && (*h.Array)[leftChildIndex] > (*h.Array)[biggest] {
		biggest = leftChildIndex
	}
	*comp += 2
	if rightRightIndex < size && (*h.Array)[rightRightIndex] > (*h.Array)[biggest] {

		biggest = rightRightIndex
	}
	*comp++
	if biggest != current {
		h.swap(current, biggest)
		h.ch <- Item{IndexFrom: biggest, IndexTo: current}
		*swaps++
		h.downHeapify(biggest, size, comp, swaps, iter)
	}
}

func (h *HeapSort) buildMaxHeap(size int, comp *int, swaps *int, iter *int) {
	for index := ((size / 2) - 1); index >= 0; index-- {
		h.downHeapify(index, size, comp, swaps, iter)
		*iter++
	}
}
