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
	leaf      *leafNode
	edges     []edge
	nodeIndex int
}
