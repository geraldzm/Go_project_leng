package funciones

type Tree struct {
	root *Nodo
}

type Nodo struct {
	value      int
	leftChild  *Nodo
	rightChild *Nodo
}
