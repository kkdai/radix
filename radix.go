package radix

import (
	"fmt"
	"strings"
)

type radixTree struct {
	root node
}

func contrainPrefix(str1, str2 string) bool {
	if sub, find := stringSubsetPrefix(str1, str2); find {
		return sub == str2
	}
	//In case "" != ""
	return str1 == str2
}

func stringSubsetPrefix(str1, str2 string) (string, bool) {
	findSubset := false
	for i := 0; i < len(str1) && i < len(str2); i++ {
		if str1[i] != str2[i] {
			retStr := str1[:i]
			return retStr, findSubset
		}
		findSubset = true
	}

	if len(str1) > len(str2) {
		return str2, findSubset
	}

	return str1, findSubset
}

func NewRadixTree() *radixTree {
	return &radixTree{}
}

func (t *radixTree) recursivePrintTree(currentNode *node, treeLevel int) {
	indentStr := ""
	for i := 1; i < treeLevel; i++ {
		indentStr = indentStr + "\t"
	}

	if currentNode.isLeafNode() {
		//Reach leaf, the end point
		fmt.Printf("%s[%d]Leaf key:%s value:%v\n", indentStr, treeLevel, currentNode.leaf.key, currentNode.leaf.value)
		return
	}

	fmt.Printf("%s[%d]Node has %d edges \n", indentStr, treeLevel, len(currentNode.edges))
	for _, edgeObj := range currentNode.edges {
		fmt.Printf("%s[%d]Edge[%s]\n", indentStr, treeLevel, string(edgeObj.containKey))
		t.recursivePrintTree(edgeObj.targetNode, treeLevel+1)
	}
}

func (t *radixTree) PrintTree() {
	t.recursivePrintTree(&t.root, 1)
}

func (t *radixTree) recursiveInsertTree(currentNode *node, containKey string, targetKey string, targetValue interface{}) {

	//Reach leaf the end point, refer this case https://goo.gl/mqXzB1
	if currentNode.isLeafNode() {
		//Insert key value as new child node of currentNode
		currentNode.insertLeafNote(containKey, targetKey, targetValue)
		//Original leaf node, become another leaf of currentNode
		currentNode.insertLeafNote("", currentNode.leaf.key, currentNode.leaf.value)
		// currentNode become not leaf node
		currentNode.leaf = nil
		return //Insert complete
	}

	for edgeIndex, _ := range currentNode.edges {
		subStr, find := stringSubsetPrefix(containKey, currentNode.edges[edgeIndex].containKey)
		if find {
			if subStr == currentNode.edges[edgeIndex].containKey {
				//trim edgeObj.containKey from targetKey
				nextTargetKey := strings.TrimPrefix(containKey, currentNode.edges[edgeIndex].containKey)
				t.recursiveInsertTree(currentNode.edges[edgeIndex].targetNode, nextTargetKey, targetKey, targetValue)
				return
			} else {
				//contain case
				//using subStr to create new node and separate two edges
				//Refer: https://goo.gl/j2MDBI
				splitNode := currentNode.insertSplitNote(subStr, currentNode.edges[edgeIndex].containKey)
				if splitNode == nil {
					panic("Unexpect error on split node")
				}

				splitContainKey := strings.TrimPrefix(containKey, subStr)
				splitNode.insertLeafNote(splitContainKey, targetKey, targetValue)
				return
			}
		}
	}

	//New edge with new key on leaf node
	//Ref: https://goo.gl/nSLTJr
	currentNode.insertLeafNote(containKey, targetKey, targetValue)
	return
}

func (t *radixTree) Insert(searchKey string, value interface{}) {
	t.recursiveInsertTree(&t.root, searchKey, searchKey, value)
}

func (t *radixTree) recursiveLoopup(searchNode *node, searchKey string) (interface{}, bool) {
	if searchNode.isLeafNode() {
		return searchNode.leaf.value, true
	}

	for _, edgeObj := range searchNode.edges {
		if contrainPrefix(searchKey, edgeObj.containKey) {
			nextSearchKey := strings.TrimPrefix(searchKey, edgeObj.containKey)
			return t.recursiveLoopup(edgeObj.targetNode, nextSearchKey)
		}
	}

	return nil, false
}

func (t *radixTree) Lookup(searchKey string) (interface{}, bool) {
	return t.recursiveLoopup(&t.root, searchKey)
}
