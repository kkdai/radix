package radix

import (
	"fmt"
	"testing"
)

func TeststringSubsetPrefix(t *testing.T) {
	sub, _ := stringSubsetPrefix([]byte("playground"), []byte("yield"))
	if string(sub) != "" {
		t.Errorf("None string subset failed, expect:%s but get:%s\n", "", sub)
	}

	sub, _ = stringSubsetPrefix([]byte("playground"), []byte("playground"))
	if string(sub) != "playground" {
		t.Errorf("full subset failed, expect:%s but get:%s\n", "playground", sub)
	}

	sub, _ = stringSubsetPrefix([]byte("playground"), []byte("playboy"))
	if string(sub) != "play" {
		t.Errorf("Sub string subset failed, expect:%s but get:%s\n", "play", sub)
	}

}

func TestPrintTree(t *testing.T) {

	rootNode := node{nodeIndex: 0, leaf: nil}

	cNode := node{nodeIndex: 1, leaf: nil}
	lNode := node{nodeIndex: 2, leaf: &leafNode{key: "company", value: 1}}
	rNode := node{nodeIndex: 3, leaf: &leafNode{key: "comflict", value: 2}}

	rootEdge := edge{containKey: []byte("com")}
	rootEdge.sourceNode = &rootNode
	rootEdge.targetNote = &cNode
	rootNode.edges = append(rootNode.edges, rootEdge)

	lEdge := edge{containKey: []byte("pany")}
	lEdge.sourceNode = &cNode
	lEdge.targetNote = &lNode

	rEdge := edge{containKey: []byte("flict")}
	rEdge.sourceNode = &cNode
	rEdge.targetNote = &rNode

	cNode.edges = append(cNode.edges, lEdge)
	cNode.edges = append(cNode.edges, rEdge)
	fmt.Println("enges:", cNode.edges)
	rTree := radixTree{}
	rTree.root = &rootNode

	rTree.PrintTree()
}
