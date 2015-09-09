package radix

import "fmt"

type radixTree struct {
	root *node
}

func stringSubsetPrefix(byt1, byt2 []byte) ([]byte, bool) {
	var biggerLen int
	if len(byt1) > len(byt2) {
		biggerLen = len(byt1)
	} else {
		biggerLen = len(byt2)
	}

	workByte := make([]byte, biggerLen)
	findSubset := false
	for i := 0; i < len(byt1) && i < len(byt2); i++ {
		if byt1[i] != byt2[i] {
			retByte := make([]byte, i+1)
			retByte = workByte[:i]
			return retByte, findSubset
		}
		findSubset = true
		workByte[i] = byt1[i]
	}
	return workByte, findSubset
}

func (t *radixTree) recursivePrintTree(currentNode *node, treeLevel int) {
	if currentNode.leaf != nil {
		fmt.Printf("Leaf[%d] key:%s value:%v\n", currentNode.nodeIndex, currentNode.leaf.key, currentNode.leaf.value)
		return
	}

	fmt.Printf("\n[%d/%d] node\n", treeLevel, currentNode.nodeIndex)
	for _, edgeObj := range currentNode.edges {
		fmt.Printf("edge[%s]-> ", string(edgeObj.containKey))
		if edgeObj.targetNote != nil {
			fmt.Printf("[%d]\n", edgeObj.targetNote.nodeIndex)
		} else {
			fmt.Printf("[nil]\n")
			continue
		}

		currentNode = edgeObj.targetNote
		t.recursivePrintTree(currentNode, treeLevel+1)
	}
}

func (t *radixTree) PrintTree() {
	currentNode := t.root
	fmt.Println("root node:", t.root, " leaf:", t.root.leaf)
	t.recursivePrintTree(currentNode, 1)
}

func (t *radixTree) Insert(searchKey []byte, value interface{}) {
	currentNode := t.root
	for currentNode.leaf != nil {

	}
}

func (t *radixTree) Lookup(searchKey []byte) (interface{}, bool) {
	return nil, false
}
