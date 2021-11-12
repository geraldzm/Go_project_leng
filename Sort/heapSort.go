package Sort

import (
	"strconv"
	"time"
)

type HeapSort struct {
	Array *[]int
	ch    chan Item
}

func (h *HeapSort) Init() {
	h.ch = make(chan Item)
	go h.sort() //GORRUTINA EJECUTADA
}

func (h *HeapSort) GetArray() *[]int {
	return h.Array
}

func (h *HeapSort) Sort() (Item, bool) {
	itm, more := <-h.ch
	return itm, more
}

//Código adaptado de https://golangbyexample.com/heapsort-in-golang/
func (h *HeapSort) sort() {

	t0 := time.Now().UnixNano() / int64(time.Millisecond) //MUESTRA DE TIEMPO INICIAL

	// arr := *h.Array
	len := len(*h.Array)

	comparisons := 0 //VARIABLES PARA CONTAR OPERACIONES ELEMENTALES
	swaps := 0
	iterations := 1

	h.buildMaxHeap(len, &comparisons, &swaps, &iterations)
	for i := len - 1; i > 0; i-- {
		// Move current root to end
		h.swap(0, i)
		h.ch <- Item{IndexFrom: i, IndexTo: 0} //COMUNICACIÓN CON LA GRAFICADORA POR MEDIO DEL CANAL
		swaps++                                //SE INCREMENTA CONTEO DE INTERCAMBIOS DE VALORES ENTRE DOS
		h.downHeapify(0, i, &comparisons, &swaps, &iterations)
		iterations++ //SE INCREMENTA CONTEO DE EVALUACIONES DE UN CICLO REPETITIVO
	}

	t1 := time.Now().UnixNano() / int64(time.Millisecond) //MUESTRA DE TIEMPO FINAL

	h.ch <- Item{ //COMUNICACIÓN FINAL CON LA GRAFICADORA POR MEDIO DEL CANAL
		Finished:   true,
		TimeEnd:    strconv.FormatInt(t0, 10),
		TimeStart:  strconv.FormatInt(t1, 10),
		TotalTime:  strconv.FormatInt(t1-t0, 10),
		TotalComp:  comparisons,
		TotalIter:  iterations,
		TotalSwaps: swaps,
	}

	close(h.ch) //SE CIERRA EL CANAL POR DONDE SE COMUNICA CON LA GRAFIACADORA
}

func (h *HeapSort) leftchildIndex(index int) int {
	return 2*index + 1
}

func (h *HeapSort) rightchildIndex(index int) int {
	return 2*index + 2
}

func (h *HeapSort) swap(first, second int) {
	temp := (*h.Array)[first]
	(*h.Array)[first] = (*h.Array)[second]
	(*h.Array)[second] = temp
}

func (h *HeapSort) leaf(index int, size int, comp *int) bool {
	*comp++ //SE INCREMENTA CONTEO DE COMPARACIONES ENTRE VALORES
	*comp++
	if index >= (size/2) && index <= size {
		return true
	}
	return false
}

func (h *HeapSort) downHeapify(current int, size int, comp *int, swaps *int, iter *int) {
	*iter++ //SE INCREMENTA CONTEO DE EVALUACIONES DE UN CICLO REPETITIVO
	if h.leaf(current, size, comp) {
		return
	}
	biggest := current
	leftChildIndex := h.leftchildIndex(current)
	rightRightIndex := h.rightchildIndex(current)
	*comp += 2 //SE INCREMENTA CONTEO DE COMPARACIONES ENTRE VALORES
	if leftChildIndex < size && (*h.Array)[leftChildIndex] > (*h.Array)[biggest] {
		biggest = leftChildIndex
	}
	*comp += 2 //SE INCREMENTA CONTEO DE COMPARACIONES ENTRE VALORES
	if rightRightIndex < size && (*h.Array)[rightRightIndex] > (*h.Array)[biggest] {
		biggest = rightRightIndex
	}
	*comp++ //SE INCREMENTA CONTEO DE COMPARACIONES ENTRE VALORES
	if biggest != current {
		h.swap(current, biggest)
		h.ch <- Item{IndexFrom: biggest, IndexTo: current} //COMUNICACIÓN CON LA GRAFICADORA POR MEDIO DEL CANAL
		*swaps++                                           //SE INCREMENTA CONTEO DE INTERCAMBIOS DE VALORES ENTRE DOS
		h.downHeapify(biggest, size, comp, swaps, iter)
	}
}

func (h *HeapSort) buildMaxHeap(size int, comp *int, swaps *int, iter *int) {
	for index := ((size / 2) - 1); index >= 0; index-- {
		h.downHeapify(index, size, comp, swaps, iter)
		*iter++ //SE INCREMENTA CONTEO DE EVALUACIONES DE UN CICLO REPETITIVO
	}
}
