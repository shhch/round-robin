package test

import (
	"github.com/shhch/round-robin"
	"testing"
)

// RoundRobin is not affected by the value of dataLen,
// but it has a greater impact on WeightedRoundRobin.
var dataLen = 100

func BenchmarkRR(b *testing.B) {
	rr := roundrobin.NewRR()
	for i := 0; i < dataLen; i++ {
		rr.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rr.Next()
	}
}

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

func BenchmarkWRR(b *testing.B) {
	wrr := roundrobin.NewWRR()
	for i := 0; i < dataLen; i++ {
		node := &Node{
			addr:   "",
			weight: int64(i),
		}
		wrr.Add(node)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wrr.Next()
	}
}
