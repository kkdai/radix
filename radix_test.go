package radix

import (
	"fmt"
	"testing"
)

func TeststringSubsetPrefix(t *testing.T) {
	sub, _ := stringSubsetPrefix("playground", "yield")
	if sub != "" {
		t.Errorf("None string subset failed, expect:%s but get:%s\n", "", sub)
	}

	sub, _ = stringSubsetPrefix("playground", "playground")
	if sub != "playground" {
		t.Errorf("full subset failed, expect:%s but get:%s\n", "playground", sub)
	}

	sub, _ = stringSubsetPrefix("playground", "playboy")
	if sub != "play" {
		t.Errorf("Sub string subset failed, expect:%s but get:%s\n", "play", sub)
	}

}

func TestPrintTree(t *testing.T) {

	rootNode := node{leaf: nil}

	cNode := node{leaf: nil}
	lNode := node{leaf: &leafNode{key: "company", value: 1}}
	rNode := node{leaf: &leafNode{key: "comflict", value: 2}}

	rootEdge := edge{containKey: "com"}
	rootEdge.sourceNode = &rootNode
	rootEdge.targetNode = &cNode
	rootNode.edges = append(rootNode.edges, rootEdge)

	lEdge := edge{containKey: "pany"}
	lEdge.sourceNode = &cNode
	lEdge.targetNode = &lNode

	rEdge := edge{containKey: "flict"}
	rEdge.sourceNode = &cNode
	rEdge.targetNode = &rNode

	cNode.edges = append(cNode.edges, lEdge)
	cNode.edges = append(cNode.edges, rEdge)
	fmt.Println("enges:", cNode.edges)
	rTree := radixTree{}
	rTree.root = rootNode

	rTree.PrintTree()
}

func TestNodeInsert(t *testing.T) {
	rTree := radixTree{root: node{}}
	rTree.root.insertLeafNote("keyAll", "keyAll", 1)
	rTree.root.insertLeafNote("open", "open", 2)
	rTree.PrintTree()
}

func TestTreeInsert(t *testing.T) {
	rTree := radixTree{}
	rTree.Insert("test", 1)
	rTree.PrintTree()
	rTree.Insert("team", 2)
	rTree.PrintTree()

	if rTree.root.edges[0].containKey != "te" {
		t.Errorf("TreeInsert: Simple case failed, expect `te`, but get %s\n", rTree.root.edges[0].containKey)
	}

}
