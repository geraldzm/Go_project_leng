package Sort

import (
	"strconv"
	"time"
)

type SelectionSort struct {
	Array *[]int
	ch    chan Item
}

func (q *SelectionSort) Init() {
	q.ch = make(chan Item)
	go q.sort() //GORRUTINA
}

func (q *SelectionSort) GetArray() *[]int {
	return q.Array
}

func (q *SelectionSort) Sort() (Item, bool) {
	itm, more := <-q.ch
	return itm, more
}

//Código adaptado de https://www.tutorialspoint.com/selection-sort-in-go-lang
func (s *SelectionSort) sort() {

	t0 := time.Now().UnixNano() / int64(time.Millisecond)

	arr := *s.Array
	len := len(*s.Array)
	var min_index int

	comparisons := 0
	swaps := 0
	iterations := 1

	for i := 0; i < len-1; i++ {
		min_index = i
		iterations++
		// Find index of minimum element
		for j := i + 1; j < len; j++ {
			comparisons++
			if arr[j] < arr[min_index] {
				min_index = j
			}
			iterations++
		}
		iterations++
		arr[i], arr[min_index] = arr[min_index], arr[i]
		s.ch <- Item{IndexFrom: min_index, IndexTo: i}
		swaps++
	}

	t1 := time.Now().UnixNano() / int64(time.Millisecond)

	s.ch <- Item{
		Finished:   true,
		TimeEnd:    strconv.FormatInt(t0, 10),
		TimeStart:  strconv.FormatInt(t1, 10),
		TotalTime:  strconv.FormatInt(t1-t0, 10),
		TotalComp:  comparisons,
		TotalIter:  iterations,
		TotalSwaps: swaps,
	}

	close(s.ch)
}
