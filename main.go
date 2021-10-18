package main

import (
	"fmt"
	"io"
	"os"
	"proyecto/funciones"
	"strconv"
)

func main() {
	//Experimentos

	n := []int{200, 400, 600, 800, 1000}
	seeds := []int{11, 12, 13, 14, 15}

	resultados := ""

	for i := 0; i < len(n); i++ {
		ran := n[i]
		resultados += "---------------------------------------\n"
		resultados += "N = " + strconv.Itoa(ran) + "\n"

		//experimento a
		resultados += "Experimento A: \n"
		resultados += fmt.Sprintf("n = %d y seed = %d\n", ran, seeds[i])

		A := funciones.Funcion1(ran, seeds[i]) //seed 19
		resultados += fmt.Sprintf("\n%v\n\n", A)

		//experimento b
		resultados += "Experimento B: \n"
		Distr := make([]int, 54)
		resultados += fmt.Sprintf("Set de entrada:\n%v\n", Distr)

		for j := 0; j < ran; j++ {
			Distr[A[j]] = Distr[A[j]] + 1
		}

		resultados += fmt.Sprintf("Set de resultado:\n%v\n\n", Distr)

		//experimento c
		resultados += "Experimento C: \n"
		TS := make([]int, 0)
		StatTS := make([]int, ran)

		resultados += fmt.Sprintf("Set de entrada:\n%v\n", TS)
		for index, val := range A {
			StatTS[index] = funciones.Funcion3(val, &TS)
		}
		resultados += fmt.Sprintf("Set de resultado:\n%v\n\n", StatTS)

		//experimento d
		resultados += "Experimento D: \n"

		TOS := make([]int, ran)
		resultados += fmt.Sprintf("Set de entrada:\n%v\n", TOS)
		copy(TOS, A)
		TOS = funciones.Funcion4(TOS, ran)
		resultados += fmt.Sprintf("Set de resultado:\n%v\n\n", TOS)

		//experimento e
		resultados += "Experimento E: \n"

		TOQ := make([]int, ran)
		resultados += fmt.Sprintf("Set de entrada:\n%v\n", TOQ)
		copy(TOQ, A)
		funciones.Funcion5(&TOQ)
		resultados += fmt.Sprintf("Set de resultado:\n%v\n\n", TOQ)

		//experimento f
		resultados += "Experimento F: \n"
		Abb := funciones.Tree{Root: nil}
		StatAbb := make([]int, ran)

		for index, val := range A {
			StatAbb[index] = Abb.Funcion9(val)
		}

		resultados += fmt.Sprintf("Set de entrada:\n%v\n", A)
		resultados += fmt.Sprintf("Set de resultado:\n%v\n\n", ArbolALista(Abb.Root))
		resultados += fmt.Sprintf("Comparaciones: \n%v\n\n", StatAbb)

		//experimento g
		resultados += "Experimento G: \n"
		AbbDSW := funciones.Tree{Root: nil}
		for _, val := range A {
			AbbDSW.Funcion9(val)
		}
		AbbDSW = *funciones.Funcion11(&AbbDSW)
		resultados += fmt.Sprintf("Set de resultado:\n%v\n\n", ArbolALista(AbbDSW.Root))

		//experimento h
		funciones.Funcion2(Distr)

		//experimento i
		resultados += "Experimento I: Semilla = 17 \n"
		cant := 10000
		arr := funciones.Funcion1(cant, 17)

		mapita := make(map[int]map[int]int)
		mapita[0] = make(map[int]int)
		mapita[1] = make(map[int]int)
		mapita[2] = make(map[int]int)
		mapita[3] = make(map[int]int)
		mapita[4] = make(map[int]int)

		for _, val := range arr {
			temp := 0
			_, temp = funciones.Funcion6(val, &TS)
			mapita[0][val] += temp
			_, temp = funciones.Funcion7(val, TOS)
			mapita[1][val] += temp
			_, temp = funciones.Funcion7(val, TOQ)
			mapita[2][val] += temp
			_, temp = funciones.Funcion10(val, Abb)
			mapita[3][val] += temp
			_, temp = funciones.Funcion10(val, AbbDSW)
			mapita[4][val] += temp
		}

		tipos := []string{"TS", "TOS", "TOQ", "ABB", "ABBDSW"}
		ind := 0
		for _, key_l := range mapita {
			lista_X := make([]int, 0)
			lista_Y := make([]int, 0)

			for key, val := range key_l {
				lista_X = append(lista_X, key)
				lista_Y = append(lista_Y, val)
			}

			resultados += fmt.Sprintf("[%s] Set X:\n%v\n\n", tipos[ind], lista_X)
			resultados += fmt.Sprintf("[%s] Set Y:\n%v\n\n", tipos[ind], lista_Y)
			ind++
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

	WriteFile("resultados.txt", resultados)
}

func ArbolALista(nodo *funciones.Nodo) string {
	if nodo == nil {
		return "[]"
	}

	nodo_str := "[" + strconv.Itoa(nodo.Value) + ", "

	nodo_str += ArbolALista(nodo.LeftChild)
	nodo_str += ", "
	nodo_str += ArbolALista(nodo.RightChild)
	nodo_str += "]"

	return nodo_str
}

func WriteFile(filename, data string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	newLine := data

	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("file appended successfully")
}

func print(w io.Writer, node *funciones.Nodo, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.Value)
	print(w, node.LeftChild, ns+2, 'L')
	print(w, node.RightChild, ns+2, 'R')
}
