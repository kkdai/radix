package radix

import "fmt"

type radixTree struct {
	root *node
}

func stringSubsetPrefix(str1, str2 string) (string, bool) {
	byt1 := []byte(str1)
	byt2 := []byte(str2)
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
			return string(retByte), findSubset
		}
		findSubset = true
		workByte[i] = byt1[i]
	}
	return string(workByte), findSubset
}

func (t *radixTree) recursivePrintTree(currentNode *node, treeLevel int) {
	if currentNode.leaf != nil {
		//Reach leaf, the end point
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
	fmt.Println("root node:", t.root, " leaf:", t.root.leaf)
	t.recursivePrintTree(t.root, 1)
}

func (t *radixTree) recursiveInsertTree(currentNode *node, totalKey string, targetKey string, targetValue interface{}) {
	if currentNode.leaf != nil {
		//Reach leaf, the end point
		//newNode := &node{}
		return
	}

	for _, edgeObj := range currentNode.edges {
		//fmt.Printf("edge[%s]-> ", string(edgeObj.containKey))
		if edgeObj.targetNote != nil {
			fmt.Printf("[%d]\n", edgeObj.targetNote.nodeIndex)
		} else {
			fmt.Printf("[nil]\n")
			continue
		}

		currentNode = edgeObj.targetNote
		t.recursiveInsertTree(currentNode, totalKey, targetKey, targetValue)
	}
}

func (t *radixTree) Insert(searchKey string, value interface{}) {
	t.recursiveInsertTree(t.root, searchKey, searchKey, value)
}

func (t *radixTree) Lookup(searchKey string) (interface{}, bool) {
	return nil, false
}
