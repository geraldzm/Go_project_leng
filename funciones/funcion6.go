package funciones

func SearchKey(key int, arry *[]int) (bool, int) {

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

// func main() {

// 	r := []int{6, 4, 3, 7, 15, 9}

// 	found, rs := SearchKey(1, &r)

// 	fmt.Println(found, rs)
// }
