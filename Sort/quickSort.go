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
	go q.sort() //GORRUTINA EJECUTADA
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

	t0 := time.Now().UnixNano() / int64(time.Millisecond) //MUESTRA DE TIEMPO INICIAL

	comparisons := 0 //VARIABLES PARA CONTAR OPERACIONES ELEMENTALES
	swaps := 0
	iterations := 1

	q.QuicksortAux(0, len(*q.Array)-1, &comparisons, &swaps, &iterations)
	t1 := time.Now().UnixNano() / int64(time.Millisecond) //MUESTRA DE TIEMPO FINAL

	q.ch <- Item{ //COMUNICACIÓN FINAL CON LA GRAFICADORA POR MEDIO DEL CANAL
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
	*comp++ //SE INCREMENTA CONTEO DE COMPARACIONES ENTRE VALORES
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

	*iter++ //SE INCREMENTA CONTEO DE EVALUACIONES DE UN CICLO REPETITIVO
	for j := low; j <= high-1; j++ {
		*comp++ //SE INCREMENTA CONTEO DE COMPARACIONES ENTRE VALORES
		if (*q.Array)[j] < pivot {
			i++

			q.ch <- Item{IndexFrom: i, IndexTo: j}
			Swap(&(*q.Array)[i], &(*q.Array)[j])
			*swaps++ //SE INCREMENTA CONTEO DE INTERCAMBIOS DE VALORES ENTRE DOS
		}
		*iter++ //SE INCREMENTA CONTEO DE EVALUACIONES DE UN CICLO REPETITIVO
	}

	q.ch <- Item{IndexFrom: i + 1, IndexTo: high} //COMUNICACIÓN CON LA GRAFICADORA POR MEDIO DEL CANAL

	Swap(&(*q.Array)[i+1], &(*q.Array)[high])
	*swaps++ //SE INCREMENTA CONTEO DE INTERCAMBIOS DE VALORES ENTRE DOS

	return i + 1
}

// Codigo tomado de: https://www.geeksforgeeks.org/quick-sort/
// Swap two int variables values
func Swap(a, b *int) {
	t := *a

	(*a) = (*b)
	(*b) = t
}
