package main

import (
	"github.com/shhch/round-robin"
)

type Node struct {
	addr   string
	weight int64
}

func (n *Node) Weight() int64 {
	return n.weight
}

func (n *Node) address() string {
	return n.addr
}

func main() {
	node1 := &Node{
		addr:   "127.0.0.1:6501",
		weight: 3,
	}
	node2 := &Node{
		addr:   "127.0.0.1:6502",
		weight: 2,
	}
	node3 := &Node{
		addr:   "127.0.0.1:6503",
		weight: 1,
	}
	wrr := roundrobin.NewWRR()
	wrr.Add(node1, node2, node3)

	for i := 0; i < 6; i++ {
		n := wrr.Next()
		node, ok := n.(*Node)
		if ok {
			println(node.address(), node.Weight())
		} else {
			println("type error")
		}
	}

}
