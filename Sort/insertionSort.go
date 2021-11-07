package Sort

type InsertionSort struct {
	Array *[]int
	ch    chan Item
}

func (in *InsertionSort) Init() {
	in.ch = make(chan Item)
	go in.sort()
}

func (in *InsertionSort) GetArray() *[]int {
	return in.Array
}

func (in *InsertionSort) sort() {
	i := 0
	key := 0
	j := 0

	for i = 0; i < len(*in.Array); i++ {
		temp_i := i
		key = (*in.Array)[i]
		j = i - 1

		for {
			if j < 0 || (*in.Array)[j] <= key {
				break
			}

			in.ch <- Item{IndexFrom: j, IndexTo: j + 1}
			(*in.Array)[j+1] = (*in.Array)[j]
			j = j - 1
		}

		in.ch <- Item{IndexFrom: temp_i, IndexTo: j + 1}
		(*in.Array)[j+1] = key
	}

	close(in.ch)
}

func (in *InsertionSort) Sort() (Item, bool) {
	itm, more := <-in.ch
	return itm, more
}
