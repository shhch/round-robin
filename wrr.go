package roundrobin

import (
	"sync"
)

type Source interface {
	Weight() int64
}

type wrrUnit struct {
	sd        Source
	srcWeight int64
	curWeight int64
}

type WeightedRoundRobin struct {
	dataList []*wrrUnit
	m        sync.Mutex
}

func (r *WeightedRoundRobin) Add(data ...Source) {
	r.m.Lock()
	defer r.m.Unlock()
	for _, v := range data {
		u := &wrrUnit{
			sd:        v,
			srcWeight: v.Weight(),
		}
		r.dataList = append(r.dataList, u)
	}
	return
}

func (r *WeightedRoundRobin) IsNil() bool {
	if r == nil || len(r.dataList) == 0 {
		return true
	}
	return false
}

// Its time complexity is O(n), n is the length of the data.
// Too much data will seriously affect its performance.
func (r *WeightedRoundRobin) Next() Source {
	if r.IsNil() {
		return nil
	}

	r.m.Lock()
	defer r.m.Unlock()

	total, best := int64(0), 0
	for k, v := range r.dataList {
		total += v.srcWeight
		v.curWeight += v.srcWeight
		if v.curWeight > r.dataList[best].curWeight {
			best = k
		}
	}

	r.dataList[best].curWeight -= total
	return r.dataList[best].sd
}

// NewWRR initializes and returns a WeightedRoundRobin structure.
// The dataList is initially empty, use the Add
// function to add data to the dataList.
func NewWRR() *WeightedRoundRobin {
	r := &WeightedRoundRobin{}
	r.dataList = make([]*wrrUnit, 0)
	return r
}
