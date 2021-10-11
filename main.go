package main

import (
	"proyecto/funciones"
)

func main() {
	//Experimentos

	n := []int{200, 400, 600, 800, 1000}

	for i := 0; i < 5; i++ {

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
		AbbDSW := funciones.Tree{Root: nil}
		for _, val := range A {
			AbbDSW.Funcion9(val)
		}
		AbbDSW = *funciones.Funcion11(&AbbDSW)

		//experimento h
		funciones.Funcion2(Distr)

		//experimento i
		arr := funciones.Funcion1(10000)
		StatArrTS := make([]int, 10000)
		StatArrTOS := make([]int, 10000)
		StatArrTOQ := make([]int, 10000)
		StatArrABB := make([]int, 10000)
		StatArrABBDSW := make([]int, 10000)

		for index, val := range arr {
			_, StatArrTS[index] = funciones.Funcion6(val, &TS)
			_, StatArrTOS[index] = funciones.Funcion7(val, TOS)
			_, StatArrTOQ[index] = funciones.Funcion7(val, TOQ)
			_, StatArrABB[index] = funciones.Funcion10(val, Abb)
			_, StatArrABBDSW[index] = funciones.Funcion10(val, AbbDSW)
		}

		//experimento j
		// alturaAbb := funciones.HeightOf(Abb.Root)
		// alturaAbbDSW := funciones.HeightOf(AbbDSW.Root)

		// densidadAbb := alturaAbb / Abb.Size()
		// densidadAbbDSW := alturaAbbDSW / AbbDSW.Size()

		//Arreglos con total de comparaciones realizadas en las inserciones
		//StatTS
		//StatAbb

		//Arreglos con total de comparaciones realizada en las busquedas
		// StatArrTS
		// StatArrTOS
		// StatArrTOQ
		// StatArrABB
		// StatArrABBDSW

	}

}
