package Sort

import (
	"strconv"
	"time"
)

type InsertionSort struct {
	Array *[]int
	ch    chan Item
}

func (in *InsertionSort) Init() {
	in.ch = make(chan Item)
	go in.sort() //GORRUTINA EJECUTADA
}

func (in *InsertionSort) GetArray() *[]int {
	return in.Array
}

func (in *InsertionSort) Sort() (Item, bool) {
	itm, more := <-in.ch
	return itm, more
}

func (in *InsertionSort) sort() {
	i := 0
	key := 0
	j := 0

	t0 := time.Now().UnixNano() / int64(time.Millisecond) //MUESTRA DE TIEMPO INICIAL

	comparisons := 0 //VARIABLES PARA CONTAR OPERACIONES ELEMENTALES
	swaps := 0
	iterations := 1

	for i = 0; i < len(*in.Array); i++ {
		iterations++ //SE INCREMENTA CONTEO DE EVALUACIONES DE UN CICLO REPETITIVO
		key = (*in.Array)[i]
		j = i - 1

		for {
			iterations++  //SE INCREMENTA CONTEO DE EVALUACIONES DE UN CICLO REPETITIVO
			comparisons++ //SE INCREMENTA CONTEO DE COMPARACIONES ENTRE VALORES
			if j < 0 || (*in.Array)[j] <= key {
				break
			}

			swaps++                                     //SE INCREMENTA CONTEO DE INTERCAMBIOS DE VALORES ENTRE DOS
			in.ch <- Item{IndexFrom: j, IndexTo: j + 1} //COMUNICACIÓN CON LA GRAFICADORA POR MEDIO DEL CANAL
			(*in.Array)[j+1] = (*in.Array)[j]
			j = j - 1
		}

		(*in.Array)[j+1] = key
	}
	t1 := time.Now().UnixNano() / int64(time.Millisecond) //MUESTRA DE TIEMPO FINAL

	in.ch <- Item{ //COMUNICACIÓN FINAL CON LA GRAFICADORA POR MEDIO DEL CANAL
		Finished:   true,
		TimeEnd:    strconv.FormatInt(t0, 10),
		TimeStart:  strconv.FormatInt(t1, 10),
		TotalTime:  strconv.FormatInt(t1-t0, 10),
		TotalComp:  comparisons,
		TotalSwaps: swaps,
		TotalIter:  iterations,
	}

	close(in.ch) //SE CIERRA EL CANAL POR DONDE SE COMUNICA CON LA GRAFIACADORA

}
