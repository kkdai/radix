package radix

import "strings"

type edge struct {
	containKey string
	sourceNode *node
	targetNode *node
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
	return n.leaf != nil && len(n.edges) == 0
}

func (n *node) insertLeafNote(containKey, totalKey string, value interface{}) {
	newNode := &node{leaf: &leafNode{key: totalKey, value: value}}
	newEdge := edge{containKey: containKey, sourceNode: n, targetNode: newNode}
	n.edges = append(n.edges, newEdge)
}

func (n *node) insertSplitNote(splitKey string, edgeKey string) *node {

	if n.isLeafNode() {
		//node is leaf node could not split, return nil
		return nil
	}

	for edgeIndex, _ := range n.edges {
		if n.edges[edgeIndex].containKey == edgeKey {
			//backup for split
			originalTargetNode := n.edges[edgeIndex].targetNode

			//insert split node
			splitNode := &node{}
			n.edges[edgeIndex] = edge{containKey: splitKey, sourceNode: n, targetNode: splitNode}

			//connect to original node
			remainKey := strings.TrimPrefix(edgeKey, splitKey)
			edgeFromSplitToOri := edge{containKey: remainKey, sourceNode: splitNode, targetNode: originalTargetNode}
			splitNode.edges = append(splitNode.edges, edgeFromSplitToOri)
			return splitNode
		}
	}
	return nil
}
