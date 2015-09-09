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
	if currentNode.isLeafNode() {
		//Reach leaf, the end point
		fmt.Printf("Leaf key:%s value:%v\n", currentNode.leaf.key, currentNode.leaf.value)
		return
	}

	fmt.Printf("\n[%d] node\n", treeLevel)
	for _, edgeObj := range currentNode.edges {
		fmt.Printf("edge[%s]-> ", string(edgeObj.containKey))
		if edgeObj.targetNote != nil {
			fmt.Printf("[node]\n")
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

func (t *radixTree) recursiveInsertTree(currentNode *node, containKey string, targetKey string, targetValue interface{}) {

	//Reach leaf the end point, refer this case https://goo.gl/mqXzB1
	if currentNode.isLeafNode() {
		//Insert key value as new child node of currentNode
		currentNode.insertChildNote(containKey, targetKey, targetValue)
		//Original leaf node, become another leaf of currentNode
		currentNode.insertChildNote(containKey, currentNode.leaf.key, currentNode.leaf.value)
		// currentNode become not leaf node
		currentNode.leaf = nil
		return
	}

	hasInsert := false
	for _, edgeObj := range currentNode.edges {
		//fmt.Printf("edge[%s]-> ", string(edgeObj.containKey))
		if edgeObj.targetNote != nil {
			fmt.Printf("[node]\n")
		} else {
			fmt.Printf("[nil]\n")
			continue
		}

		currentNode = edgeObj.targetNote
		t.recursiveInsertTree(currentNode, containKey, targetKey, targetValue)
	}

	if !hasInsert {
		//New edge with new key on leaf node
		//Ref: https://goo.gl/nSLTJr
		currentNode.insertChildNote(containKey, targetKey, targetValue)
	}
}

func (t *radixTree) Insert(searchKey string, value interface{}) {
	t.recursiveInsertTree(t.root, searchKey, searchKey, value)
}

func (t *radixTree) Lookup(searchKey string) (interface{}, bool) {
	return nil, false
}
