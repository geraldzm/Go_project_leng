package Sort

import (
	"strconv"
	"time"
)

type BubbleSort struct {
	Array *[]int
	ch    chan Item
}

func (b *BubbleSort) Init() {
	b.ch = make(chan Item)
	go b.sort() //GORRUTINA
}

func (b BubbleSort) GetArray() *[]int {
	return b.Array
}

func (b BubbleSort) sort() {

	t0 := time.Now().UnixNano() / int64(time.Millisecond)

	ar := *b.Array
	len := len(*b.Array)
	comparisons := 0
	swaps := 0
	iterations := 1

	for i := 0; i < len-1; i++ {
		iterations++
		for j := 0; j < len-i-1; j++ {
			comparisons++
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
				swaps++
				b.ch <- Item{IndexFrom: j, IndexTo: j + 1}
			}
			iterations++
		}
	}

	t1 := time.Now().UnixNano() / int64(time.Millisecond)

	b.ch <- Item{
		Finished:   true,
		TimeEnd:    strconv.FormatInt(t0, 10),
		TimeStart:  strconv.FormatInt(t1, 10),
		TotalTime:  strconv.FormatInt(t1-t0, 10),
		TotalComp:  comparisons,
		TotalSwaps: swaps,
		TotalIter:  iterations,
	}

	close(b.ch)
}

func (b BubbleSort) Sort() (Item, bool) {
	itm, more := <-b.ch
	return itm, more
}
