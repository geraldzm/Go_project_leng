package funciones

func Funcion10(num int, tree Tree) (bool, int) {
	node := tree.root
	counter := 1

	for node != nil {
		if node.value > num {
			counter++
			node = node.leftChild
		} else if node.value < num {
			counter += 2
			node = node.rightChild
		} else {
			counter += 2
			return true, counter
		}

		counter++
	}
	return false, counter
}
