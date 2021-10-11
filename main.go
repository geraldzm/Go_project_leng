package main

import (
	"proyecto/funciones"
)

func main() {
	//Experimentos

	n := []int{200} // 400, 600, 800, 1000

	for i := 0; i < 1; i++ {

		ran := n[i]

		//experimento a
		A := funciones.Funcion1(ran)

		//experimento b
		Distr := make([]int, 54)

		for j := 0; j < ran; j++ {
			Distr[A[j]] = Distr[A[j]] + 1
		}

		//experimento c
		TS := make([]int, 0)
		StatTS := make([]int, ran)

		for index, val := range A {
			StatTS[index] = funciones.Funcion3(val, &TS)
		}

		//experimento d
		TOS := make([]int, ran)
		copy(TOS, A)
		TOS = funciones.Funcion4(TOS, ran)

		//experimento e
		TOQ := make([]int, ran)
		copy(TOQ, A)
		funciones.Funcion5(&TOQ)

		//experimento f
		Abb := funciones.Tree{Root: nil}
		StatAbb := make([]int, ran)
		for index, val := range A {
			StatAbb[index] = Abb.Funcion9(val)
		}

		//experimento g

	}

}
