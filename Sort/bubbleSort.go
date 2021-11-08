package Sort

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

	b.ch <- Item{ Finished: false, TimeEnd: "20", TimeStart: "8", TotalTime: "12", TotalComp: 1000}

	close(b.ch)
}

func (b BubbleSort) Sort() (Item, bool) {
	itm, more := <-b.ch
	return itm, more
}
