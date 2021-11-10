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
	go in.sort()
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

	t0 := time.Now().UnixNano() / int64(time.Millisecond)
	comparisons := 0
	swaps := 0
	iterations := 1

	for i = 0; i < len(*in.Array); i++ {
		iterations++
		key = (*in.Array)[i]
		j = i - 1

		for {
			if j < 0 || (*in.Array)[j] <= key {
				comparisons++
				break
			}

			swaps++
			in.ch <- Item{IndexFrom: j, IndexTo: j + 1}
			(*in.Array)[j+1] = (*in.Array)[j]
			j = j - 1
		}

		(*in.Array)[j+1] = key
	}
	t1 := time.Now().UnixNano() / int64(time.Millisecond)

	in.ch <- Item{
		Finished:   true,
		TimeEnd:    strconv.FormatInt(t0, 10),
		TimeStart:  strconv.FormatInt(t1, 10),
		TotalTime:  strconv.FormatInt(t1-t0, 10),
		TotalComp:  comparisons,
		TotalSwaps: swaps,
		TotalIter:  iterations,
	}

	close(in.ch)
}
