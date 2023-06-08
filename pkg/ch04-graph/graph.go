package ch04

import (
	"math"
)

type Edge struct {
	Target string
	Weight float64
}

func NewEdge(target string, weight float64) Edge {
	return Edge{target, weight}
}

type Node struct {
	Tag    string
	Edges  []Edge
	Cost   float64
	Parent string
}

func NewNode(tag string, edges []Edge) *Node {
	return &Node{Tag: tag, Edges: edges, Cost: math.Inf(1)}
}

type Graph struct {
	Nodes     map[string]*Node
	processed []string
}

func NewGraph(nodes map[string]*Node) *Graph {
	graph := &Graph{Nodes: nodes}
	graph.init()
	return graph
}

func (g *Graph) init() {
	g.Nodes["S"].Cost = 0
}

func (g *Graph) Reset() {
	g.processed = nil
	for _, node := range g.Nodes {
		node.Cost = math.Inf(1)
		node.Parent = ""
	}
	g.init()
}

func (g *Graph) Cheapest() *Node {
	cost := math.Inf(1)
	var cheapest *Node
	for tag, node := range g.Nodes {
		if g.Processed(tag) {
			continue
		} else if node.Cost < cost {
			cost = node.Cost
			cheapest = node
		}
	}
	return cheapest
}

func (g *Graph) Processed(tag string) bool {
	for _, tg := range g.processed {
		if tg == tag {
			return true
		}
	}
	return false
}

func (g *Graph) Mark(tag string) {
	g.processed = append(g.processed, tag)
}
