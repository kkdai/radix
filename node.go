package radix

type edge struct {
	containKey []byte
	sourceNode *node
	targetNote *node
}

type node struct {
	isLeaf bool
	edges  []edge
}
