package funciones

func Insert(key int, arry *[]int) int {

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

// func main() {

// 	r := []int{6, 4, 3, 7}

// 	rs := Insert(15, &r)

// 	fmt.Println(rs)
// 	fmt.Println(r)
// }
