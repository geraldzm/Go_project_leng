package funciones

type Tree struct {
	Root *Nodo
}

type Nodo struct {
	Value      int
	LeftChild  *Nodo
	RightChild *Nodo
}
