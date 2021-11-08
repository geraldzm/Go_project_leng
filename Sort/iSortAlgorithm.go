package Sort


type Item struct {
    IndexFrom int
    IndexTo int
    Finished bool
    TimeStart string
    TimeEnd string
    TotalTime string
    TotalComp int
}

type ISortAlgorithm interface {
    Sort() (Item, bool)
    Init()
    GetArray() *[]int
}