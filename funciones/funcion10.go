package funciones

func Funcion10(num int, tree Tree) (bool, int) {
	node := tree.root
	counter := 0

	for node != nil {
		counter++
		if node.value > num {
			node = node.leftChild
		} else if node.value < num {
			node = node.rightChild
		} else {
			return true, counter
		}
	}
	return false, counter
}
