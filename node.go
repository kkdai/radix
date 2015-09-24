package radix

import "strings"

type Edge struct {
	containKey string
	sourceNode *Node
	targetNode *Node
}

type leafNode struct {
	key   string
	value interface{}
}

type Node struct {
	leaf  *leafNode
	edges []Edge
}

func (n *Node) isLeafNode() bool {
	return n.leaf != nil && len(n.edges) == 0
}

func (n *Node) insertLeafNote(containKey, totalKey string, value interface{}) {
	newNode := &Node{leaf: &leafNode{key: totalKey, value: value}}
	newEdge := Edge{containKey: containKey, sourceNode: n, targetNode: newNode}
	n.edges = append(n.edges, newEdge)
}

func (n *Node) insertSplitNote(splitKey string, edgeKey string) *Node {

	if n.isLeafNode() {
		//node is leaf node could not split, return nil
		return nil
	}

	for edgeIndex, _ := range n.edges {
		if n.edges[edgeIndex].containKey == edgeKey {
			//backup for split
			originalTargetNode := n.edges[edgeIndex].targetNode

			//insert split node
			splitNode := &Node{}
			n.edges[edgeIndex] = Edge{containKey: splitKey, sourceNode: n, targetNode: splitNode}

			//connect to original node
			remainKey := strings.TrimPrefix(edgeKey, splitKey)
			edgeFromSplitToOri := Edge{containKey: remainKey, sourceNode: splitNode, targetNode: originalTargetNode}
			splitNode.edges = append(splitNode.edges, edgeFromSplitToOri)
			return splitNode
		}
	}
	return nil
}
