package map_test

import (
	"github.com/JnBrymN/GoSkipList"
	. "launchpad.net/gocheck"
	"testing"
)

type MapBenchSuite struct{}

var _ = Suite(&MapBenchSuite{})

// gocheck link to go test - only needs doing once for whole lib
func Test(t *testing.T) {
	TestingT(t)
}

func compareInts(a, b int) bool { return a < b }

func (s *MapBenchSuite) BenchmarkPut(c *C) {
	m := NewMap(compareInts)
	i := 0
	for ; i < c.N; i++ {
		m.put(i, i)
	}
	c.Assert(m.Len(), Equals, i+1)
}

func fillMap(n int) *Map {
	m := NewMap(compareInts)
	for j := 0; j < n; j++ {
		m.Put(j, j)
	}
	return &m
}

func (s *MapBenchSuite) BenchmarkGet1000(c *C) {
	m := fillMap(1000)
	i := 0
	for ; i < c.N; i++ {
		c.Assert(m.Get(i), Equals, i)
	}
	c.Assert(m.Len(), Equals, i+1)
}

func (s *MapBenchSuite) BenchmarkGet10000(c *C) {
	m := fillMap(10000)
	i := 0
	for ; i < c.N; i++ {
		c.Assert(m.Get(i), Equals, i)
	}
	c.Assert(m.Len(), Equals, i+1)
}

func (s *MapBenchSuite) BenchmarkGet100000(c *C) {
	m := fillMap(100000)
	i := 0
	for ; i < c.N; i++ {
		c.Assert(m.Get(i), Equals, i)
	}
	c.Assert(m.Len(), Equals, i+1)
}
