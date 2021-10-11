package funciones

//https://www.golangprograms.com/golang-program-for-implementation-of-binary-search.html
func Funcion7(val int, arr []int) (bool, int) {

	counter := 0
	low := 0
	high := len(arr) - 1

	for low <= high {
		counter++
		median := (low + high) / 2

		if arr[median] < val {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) || arr[low] != val {
		return false, counter
	}

	return true, counter
}
