type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	cloneOfNode := map[*Node]*Node{}
	return cloneGraphHelper(node, cloneOfNode)
}

func cloneGraphHelper(node *Node, cloneOfNode map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if _, ok := cloneOfNode[node]; ok {
		return cloneOfNode[node]
	}

	cloneNode := &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}
	cloneOfNode[node] = cloneNode

	for _, neighbor := range node.Neighbors {
		cloneNode.Neighbors = append(cloneNode.Neighbors, cloneGraphHelper(neighbor))
	}
	return cloneNode
}
