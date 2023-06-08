package ch04

func Dijkstra(g *Graph) []string {
	for {
		chnode := g.Cheapest()
		if chnode == nil {
			break
		} else if chnode.Edges == nil {
			g.Mark(chnode.Tag)
			continue
		}
		for _, edge := range chnode.Edges {
			node := g.Nodes[edge.Target]
			newCost := chnode.Cost + edge.Weight
			if newCost < node.Cost {
				node.Cost = newCost
				node.Parent = chnode.Tag
			}
		}
		g.Mark(chnode.Tag)
	}
	path := make([]string, 0, 10)
	node := g.Nodes["F"]
	path = append(path, node.Tag)
	for node.Parent != "" {
		path = append(path, node.Parent)
		node = g.Nodes[node.Parent]
	}
	n := len(path) - 1
	reversed := make([]string, n+1)
	for i, tag := range path {
		reversed[n-i] = tag
	}
	return reversed
}
