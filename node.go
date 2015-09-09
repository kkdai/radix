package radix

type edge struct {
	containKey string
	sourceNode *node
	targetNote *node
}

type leafNode struct {
	key   string
	value interface{}
}

type node struct {
	leaf  *leafNode
	edges []edge
}

func (n *node) isLeafNode() bool {
	return n.leaf != nil
}

func (n *node) insertChildNote(containKey, totalKey string, value interface{}) {
	newNode := &node{leaf: &leafNode{key: totalKey, value: value}}
	newEdge := edge{containKey: containKey, sourceNode: n, targetNote: newNode}
	n.edges = append(n.edges, newEdge)
}
