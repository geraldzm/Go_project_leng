package funciones

func Funcion1(n int) []int {
	var seed int = 19
	var periodo int = 2048
	var a int = 21 //multiplicador
	var b int = 25 //incremento

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		seed = (a*seed + b) % periodo
		arr[i] = seed % 53
	}
	return arr
}
