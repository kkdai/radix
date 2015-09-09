package radix

type edge struct {
	containKey []byte
	sourceNode *node
	targetNote *node
}

type leafNode struct {
	key   string
	value interface{}
}

type node struct {
	leaf      *leafNode
	edges     []edge
	nodeIndex int
}
