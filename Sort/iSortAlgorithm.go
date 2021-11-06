package Sort


type Item struct {
    IndexFrom int
    IndexTo int
}

type ISortAlgorithm interface {
    Sort() (Item, bool)
    Init()
    GetArray() *[]int
}