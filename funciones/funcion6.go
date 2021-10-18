package funciones

func Funcion6(key int, arry *[]int) (bool, int) {
	comparations := 1

	for i := 0; i < len(*arry); i++ {
		comparations++
		if (*arry)[i] == key {
			return true, comparations
		}

		comparations++
	}

	return false, comparations
}
