package Sort

type QuickSort struct {
	Array *[]int
	ch    chan Item
}

func (q *QuickSort) Init() {
	q.ch = make(chan Item)
	go q.sort()
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
	q.QuicksortAux(0, len(*q.Array)-1)

	q.ch <- Item{ Finished: true, TimeEnd: "10", TimeStart: "7", TotalTime: "3", TotalComp: 100}

	close(q.ch)
}

// Codigo tomado de: https://www.geeksforgeeks.org/quick-sort/
// Funcion de ordenamiento usando quicksort
func (q *QuickSort) QuicksortAux(low, high int) {
	if low < high {
		pi := q.Partition(low, high)

		q.QuicksortAux(low, pi-1)
		q.QuicksortAux(pi+1, high)
	}
}

// Codigo tomado de: https://www.geeksforgeeks.org/quick-sort/
/* This function takes last element as pivot, places
   the pivot element at its correct position in sorted
    array, and places all smaller (smaller than pivot)
   to left of pivot and all greater elements to right
   of pivot */
func (q *QuickSort) Partition(low, high int) int {
	pivot := (*q.Array)[high]

	i := low - 1

	for j := low; j <= high-1; j++ {
		if (*q.Array)[j] < pivot {
			i++

			q.ch <- Item{IndexFrom: i, IndexTo: j}
			Swap(&(*q.Array)[i], &(*q.Array)[j])
		}
	}

	q.ch <- Item{IndexFrom: i + 1, IndexTo: high}
	Swap(&(*q.Array)[i+1], &(*q.Array)[high])

	return i + 1
}

// Codigo tomado de: https://www.geeksforgeeks.org/quick-sort/
// Swap two int variables values
func Swap(a, b *int) {
	t := *a

	(*a) = (*b)
	(*b) = t
}
