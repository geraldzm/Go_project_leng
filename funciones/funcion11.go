package funciones

import (
	"math"
)

func Funcion11(tree *Tree) *Tree {
	vine := CreateRightVine(tree.Root)
	return &Tree{Root: BalanceVine(vine)}
}

func BalanceVine(node *Nodo) *Nodo {
	nodeCount := GetNodeCount(node)
	times := int(math.Log2(float64(nodeCount)))
	newRoot := node
	for i := 0; i < times; i++ {
		newRoot = LeftRotate(newRoot)
		node = newRoot.RightChild
		for j := 0; j < nodeCount/2-1; j++ {
			node = LeftRotate(node)
			node = node.RightChild
		}
		nodeCount >>= 1
	}
	return newRoot
}

func LeftRotate(node *Nodo) *Nodo {
	if node.RightChild != nil {
		tmpNodoRight := node.RightChild
		node.RightChild = tmpNodoRight.RightChild
		tmpNodoRight.RightChild = tmpNodoRight.LeftChild
		tmpNodoRight.LeftChild = node.LeftChild
		node.LeftChild = tmpNodoRight

		temp := node.Value
		node.Value = tmpNodoRight.Value
		tmpNodoRight.Value = temp
	}
	return node
}

func RightRotate(node *Nodo) *Nodo {
	if node.LeftChild != nil {
		tmpNodoLeft := node.LeftChild
		node.LeftChild = tmpNodoLeft.LeftChild
		tmpNodoLeft.LeftChild = tmpNodoLeft.RightChild
		tmpNodoLeft.RightChild = node.RightChild
		node.RightChild = tmpNodoLeft

		temp := node.Value
		node.Value = tmpNodoLeft.Value
		tmpNodoLeft.Value = temp
	}
	return node
}

func CreateRightVine(node *Nodo) *Nodo {
	for node.LeftChild != nil {
		node = RightRotate(node)
	}
	if node.RightChild != nil {
		node.RightChild = CreateRightVine(node.RightChild)
	}
	return node
}

func GetNodeCount(node *Nodo) int {
	if node == nil {
		return 0
	}
	i := 1
	for node.RightChild != nil {
		node = node.RightChild
		i++
	}
	return i
}
