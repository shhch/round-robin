package main

import (
	"github.com/shhch/round-robin"
)

func main() {
	rr := roundrobin.NewRR()
	rr.Add(1)
	rr.Add(2)
	rr.Add(3)

	for i := 0; i < 6; i++ {
		data := rr.Next()
		val, ok := data.(int)
		if ok {
			println(val)
		} else {
			println("type error")
		}

	}
}
