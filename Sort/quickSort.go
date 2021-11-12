package Sort

import (
	"strconv"
	"time"
)

type QuickSort struct {
	Array *[]int
	ch    chan Item
}

func (q *QuickSort) Init() {
	q.ch = make(chan Item)
	go q.sort() //GORRUTINA
}

func (q *QuickSort) GetArray() *[]int {
	return q.Array
}

func (q *QuickSort) Sort() (Item, bool) {
	itm, more := <-q.ch
	return itm, more
}

// Funcion para ordenar un arreglo de enteros usando quick sort
func (q *QuickSort) sort() {

	t0 := time.Now().UnixNano() / int64(time.Millisecond)
	comparisons := 0
	swaps := 0
	iterations := 1

	q.QuicksortAux(0, len(*q.Array)-1, &comparisons, &swaps, &iterations)
	t1 := time.Now().UnixNano() / int64(time.Millisecond)

	q.ch <- Item{
		Finished:   true,
		TimeEnd:    strconv.FormatInt(t0, 10),
		TimeStart:  strconv.FormatInt(t1, 10),
		TotalTime:  strconv.FormatInt(t1-t0, 10),
		TotalComp:  comparisons,
		TotalSwaps: swaps,
		TotalIter:  iterations,
	}

	close(q.ch)
}

// Codigo tomado de: https://www.geeksforgeeks.org/quick-sort/
// Funcion de ordenamiento usando quicksort
func (q *QuickSort) QuicksortAux(low int, high int, comp *int, swaps *int, iter *int) {
	*comp++
	if low < high {
		pi := q.Partition(low, high, comp, swaps, iter)

		q.QuicksortAux(low, pi-1, comp, swaps, iter)
		q.QuicksortAux(pi+1, high, comp, swaps, iter)
	}
}

// Codigo tomado de: https://www.geeksforgeeks.org/quick-sort/
/* This function takes last element as pivot, places
   the pivot element at its correct position in sorted
    array, and places all smaller (smaller than pivot)
   to left of pivot and all greater elements to right
   of pivot */
func (q *QuickSort) Partition(low int, high int, comp *int, swaps *int, iter *int) int {
	pivot := (*q.Array)[high]

	i := low - 1

	*iter++
	for j := low; j <= high-1; j++ {
		*comp++
		if (*q.Array)[j] < pivot {
			i++

			q.ch <- Item{IndexFrom: i, IndexTo: j}
			Swap(&(*q.Array)[i], &(*q.Array)[j])
			*swaps++
		}
	}

	q.ch <- Item{IndexFrom: i + 1, IndexTo: high}
	Swap(&(*q.Array)[i+1], &(*q.Array)[high])
	*swaps++

	return i + 1
}

// Codigo tomado de: https://www.geeksforgeeks.org/quick-sort/
// Swap two int variables values
func Swap(a, b *int) {
	t := *a

	(*a) = (*b)
	(*b) = t
}
