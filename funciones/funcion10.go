package funciones

func Funcion10(num int, tree Tree) (bool, int) {
	node := tree.root
	counter := 1

	for node != nil {

		if node.value > num {
			node = node.leftChild
		} else if node.value < num {
			node = node.rightChild
		} else {
			counter += 2
			return true, counter
		}

		counter += 3
	}

	return false, counter
}
