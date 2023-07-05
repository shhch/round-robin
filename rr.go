package roundrobin

import (
	"sync"
	"sync/atomic"
)

type RoundRobin struct {
	index    *atomic.Int64
	dataList []any
	m        sync.Mutex
}

func (r *RoundRobin) Add(data ...any) {
	r.m.Lock()
	r.dataList = append(r.dataList, data...)
	r.m.Unlock()
	return
}

func (r *RoundRobin) IsNil() bool {
	if r == nil || len(r.dataList) == 0 {
		return true
	}
	return false
}

func (r *RoundRobin) Next() any {
	if r.IsNil() {
		return nil
	}
	i := int(r.index.Add(1)-1) % len(r.dataList)
	return r.dataList[i]
}

// NewRR initializes and returns a RoundRobin structure.
// The dataList is initially empty, use the Add
// function to add data to the dataList.
func NewRR() *RoundRobin {
	r := &RoundRobin{}
	r.index = &atomic.Int64{}
	r.index.Store(0)
	r.dataList = make([]any, 0)
	return r
}
