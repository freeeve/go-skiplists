package skiplist

import (
	"math/rand"
	. "gopkg.in/check.v1"
)

type MapBenchSuite struct{}

var _ = Suite(&MapBenchSuite{})

func (s *MapBenchSuite) BenchmarkPut(c *C) {
	m := NewMap(compareInts)
	i := 0
	for ; i < c.N; i++ {
		m.Put(i, i)
	}
	c.Assert(m.Len(), Equals, i)
}

func fillMap(n int) *Map {
	m := NewMap(compareInts)
	for j := 0; j < n; j++ {
		m.Put(j, j*2)
	}
	return m
}

func fillMapRand(n int) *Map {
	m := NewMap(compareInts)
	r := rand.New(rand.NewSource(123123123))
	for j := 0; j < n; j++ {
		k := r.Int()
		m.Put(k, k*2)
	}
	return m
}

func benchmarkGetN(n int, c *C) {
	c.StopTimer()
	m := fillMap(n)
	c.StartTimer()
	i := 0
	for ; i < c.N; i++ {
		x, ok := m.Get(i)
		if i < n {
			c.Assert(ok, Equals, true)
			c.Assert(x, Equals, i*2)
		} else {
			c.Assert(ok, Equals, false)
			c.Assert(x, IsNil)
		}
	}
}

func (s *MapBenchSuite) BenchmarkGet1000(c *C) {
	benchmarkGetN(1000, c)
}
func (s *MapBenchSuite) BenchmarkGet10000(c *C) {
	benchmarkGetN(10000, c)
}
func (s *MapBenchSuite) BenchmarkGet100000(c *C) {
	benchmarkGetN(100000, c)
}
func (s *MapBenchSuite) BenchmarkGet1000000(c *C) {
	benchmarkGetN(1000000, c)
}

/*
func (s *MapBenchSuite) BenchmarkGet10000000(c *C) {
	benchmarkGetN(10000000, c)
}*/

func benchmarkGetNRand(n int, c *C) {
	c.StopTimer()
	m := fillMap(n)
	r := rand.New(rand.NewSource(1423123432))
	c.StartTimer()
	i := 0
	for ; i < c.N; i++ {
		k := r.Int()
		x, ok := m.Get(k)
		if ok {
			c.Assert(x, Equals, k*2)
		} else {
			c.Assert(x, IsNil)
		}
	}
}

func (s *MapBenchSuite) BenchmarkGetRand100000(c *C) {
	benchmarkGetNRand(100000, c)
}

func (s *MapBenchSuite) BenchmarkGetRand1000000(c *C) {
	benchmarkGetNRand(1000000, c)
}
func (s *MapBenchSuite) BenchmarkGetRand10000000(c *C) {
	benchmarkGetNRand(10000000, c)
}
