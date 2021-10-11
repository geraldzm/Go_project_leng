package funciones

//https://www.tutorialspoint.com/selection-sort-in-go-lang
func Funcion4(arr []int, size int) []int {
	var min_index int
	var temp int
	for i := 0; i < size-1; i++ {
		min_index = i
		// Find index of minimum element
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[min_index] {
				min_index = j
			}
		}
		temp = arr[i]
		arr[i] = arr[min_index]
		arr[min_index] = temp
	}
	return arr
}
