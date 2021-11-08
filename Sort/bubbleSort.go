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
	go b.sort()
}

func (b BubbleSort) GetArray() *[]int {
	return b.Array
}

func (b BubbleSort) sort() {

	t0 := time.Now().UnixNano() / int64(time.Millisecond)

	ar := *b.Array
	len := len(*b.Array)

	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
				b.ch <- Item{IndexFrom: j, IndexTo: j + 1}
			}
		}
	}

	t1 := time.Now().UnixNano() / int64(time.Millisecond)

	b.ch <- Item {
		Finished: true,
		TimeEnd: strconv.FormatInt(t0 / int64(time.Second), 10),
		TimeStart: strconv.FormatInt(t1 / int64(time.Second), 10),
		TotalTime: strconv.FormatInt((t1-t0)  / int64(time.Second), 10),
		TotalComp: 1000,
	}

	close(b.ch)
}

func (b BubbleSort) Sort() (Item, bool) {
	itm, more := <-b.ch
	return itm, more
}
