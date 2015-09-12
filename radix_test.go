package radix

import (
	"fmt"
	"testing"
)

func TeststringSubsetPrefix(t *testing.T) {
	sub, _ := getSubsetPrefix("playground", "yield")
	if sub != "" {
		t.Errorf("None string subset failed, expect:%s but get:%s\n", "", sub)
	}

	sub, _ = getSubsetPrefix("playground", "playground")
	if sub != "playground" {
		t.Errorf("full subset failed, expect:%s but get:%s\n", "playground", sub)
	}

	sub, _ = getSubsetPrefix("playground", "playboy")
	if sub != "play" {
		t.Errorf("Sub string subset failed, expect:%s but get:%s\n", "play", sub)
	}

	ss1 := "main"
	ss2 := "mainly"
	sub, _ = getSubsetPrefix(ss2, ss1)
	if sub != ss1 {
		t.Errorf("Sub string subset failed, expect:%s but get:%s\n", "main", sub)
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
	rTree := NewRadixTree()
	rTree.root.insertLeafNote("keyAll", "keyAll", 1)
	rTree.root.insertLeafNote("open", "open", 2)
	rTree.PrintTree()
}

func TestTreeInsert(t *testing.T) {
	rTree := NewRadixTree()
	rTree.Insert("test", 1)
	rTree.Insert("team", 2)

	if rTree.root.edges[0].containKey != "te" {
		t.Errorf("TreeInsert: Simple case failed, expect `te`, but get %s\n", rTree.root.edges[0].containKey)
	}

	rTree2 := NewRadixTree()
	rTree2.Insert("main", 1)
	rTree2.Insert("mainly", 2)

	if rTree2.root.edges[0].containKey != "main" {
		t.Errorf("TreeInsert: Simple case failed, expect `main`, but get %s\n", rTree.root.edges[0].containKey)
	}
}

func TestLookup(t *testing.T) {
	rTree := NewRadixTree()
	rTree.Insert("test", 1)
	rTree.Insert("team", 2)
	rTree.Insert("trobot", 3)
	rTree.Insert("apple", 4)
	rTree.Insert("app", 5)
	rTree.Insert("tesla", 6)

	ret, find := rTree.Lookup("team")
	if !find || ret != 2 {
		t.Errorf("Lookup failed, expect '2', but get %v", ret)
	}

	ret, find = rTree.Lookup("apple")
	if !find || ret != 4 {
		t.Errorf("Lookup failed, expect '4', but get %v", ret)
	}

	ret, find = rTree.Lookup("tesla")
	if !find || ret != 6 {
		t.Errorf("Lookup failed, expect '6', but get %v", ret)
	}

	ret, find = rTree.Lookup("app")
	if !find || ret != 5 {
		t.Errorf("Lookup failed, expect '5', but get %v", ret)
	}

	rTree.Insert("app", 7)
	rTree.PrintTree()
	ret, find = rTree.Lookup("app")
	fmt.Println(ret, find)
	if !find || ret != 7 {
		t.Errorf("Insert update lookup failed, expect '7', but get %v", ret)
	}

}

func TestLocateLeafNode(t *testing.T) {
	rTree := NewRadixTree()
	rTree.Insert("test", 1)
	rTree.Insert("team", 2)
	rTree.Insert("trobot", 3)
	rTree.Insert("apple", 4)
	rTree.Insert("app", 5)
	rTree.Insert("tesla", 6)

	cNode, pNode, find := rTree.locateLeafNode("trobot")
	fmt.Println(cNode, pNode, find)

	cNode, pNode, find = rTree.locateLeafNode("trobota")
	fmt.Println(cNode, pNode, find)
}

func TestFindParent(t *testing.T) {
	rTree := NewRadixTree()
	rTree.Insert("test", 1)
	rTree.Insert("team", 2)
	rTree.Insert("trobot", 3)
	rTree.Insert("apple", 4)
	rTree.Insert("app", 5)
	rTree.Insert("tesla", 6)

	cNode, pNode, find := rTree.locateLeafNode("trobot")
	fmt.Println(cNode, pNode, find)
	cParent, cFind := rTree.findParent(cNode)
	if cFind {
		fmt.Println(cParent.edges)
	} else {
		t.Errorf("Failed in find parentNode")
	}

	nextParent, ccFind := rTree.findParent(cParent)
	if ccFind {
		fmt.Println(nextParent.edges)
	} else {
		t.Errorf("Failed in find parentNode")
	}

	pRoot, fRoot := rTree.findParent(&rTree.root)
	if fRoot {
		if pRoot != &rTree.root {
			t.Errorf("Failed on find parent on root")
		}
		fmt.Println(pRoot.edges)
	} else {
		t.Errorf("Failed on find parent on root, cannot find it.")
	}
}

func TestDelete(t *testing.T) {
	rTree := NewRadixTree()
	rTree.Insert("test", 1)
	rTree.Insert("team", 2)
	rTree.Insert("trobot", 3)
	rTree.Insert("apple", 4)
	rTree.Insert("app", 5)
	rTree.Insert("tesla", 6)
	rTree.PrintTree()

	ret := rTree.Delete("team")
	if !ret {
		t.Errorf("Basic delete failed")
	}
	rTree.PrintTree()

	ret = rTree.Delete("app")
	rTree.PrintTree()

	ret = rTree.Delete("apple")
	if !ret {
		t.Errorf("Recursive delete on parent node, failed")
	}
	fmt.Println("---------------")
	rTree.PrintTree()
	rTree.Delete("trobot")
	ret = rTree.Delete("test")
	ret = rTree.Delete("tesla")
	fmt.Println("---------------")
	rTree.PrintTree()

	if rTree.Delete("anyWrongKey") {
		t.Errorf("Delete empty should not return true")
	}
}
