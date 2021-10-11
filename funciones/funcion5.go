package funciones

// Funcion para ordenar un arreglo de enteros usando quick sort
func Funcion5(array *[]int) {
	QuicksortAux(array, 0, len(*array)-1)
}

// Codigo tomado de: https://www.geeksforgeeks.org/quick-sort/
// Funcion de ordenamiento usando quicksort
func QuicksortAux(array *[]int, low, high int) {
	if low < high {
		pi := Partition(array, low, high)

		QuicksortAux(array, low, pi-1)
		QuicksortAux(array, pi+1, high)
	}
}

// Codigo tomado de: https://www.geeksforgeeks.org/quick-sort/
/* This function takes last element as pivot, places
   the pivot element at its correct position in sorted
    array, and places all smaller (smaller than pivot)
   to left of pivot and all greater elements to right
   of pivot */
func Partition(array *[]int, low, high int) int {
	pivot := (*array)[high]

	i := low - 1

	for j := low; j <= high-1; j++ {
		if (*array)[j] < pivot {
			i++

			Swap(&(*array)[i], &(*array)[j])
		}
	}

	Swap(&(*array)[i+1], &(*array)[high])

	return i + 1
}

// Codigo tomado de: https://www.geeksforgeeks.org/quick-sort/
// Swap two int variables values
func Swap(a, b *int) {
	t := *a

	(*a) = (*b)
	(*b) = t
}
