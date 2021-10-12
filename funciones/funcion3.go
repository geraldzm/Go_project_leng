package funciones

func Funcion3(key int, arry *[]int) int {

	comparations := 1

	for i := 0; i < len(*arry); i++ {

		comparations++
		if (*arry)[i] == key {
			return comparations
		}

		comparations++
	}

	(*arry) = append((*arry), key)

	return comparations
}
