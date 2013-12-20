package main

import (
	. "launchpad.net/gocheck"
)

type MapBenchSuite struct{}

var _ = Suite(&MapBenchSuite{})

func (s *MapBenchSuite) BenchmarkPut(c *C) {
	m := NewMap(compareInts)
	i := 0
	for ; i < c.N; i++ {
		m.Put(i, i)
	}
	c.Assert(m.Len(), Equals, i+1)
}

func fillMap(n int) *Map {
	m := NewMap(compareInts)
	for j := 0; j < n; j++ {
		m.Put(j, j)
	}
	return m
}

func (s *MapBenchSuite) BenchmarkGet1000(c *C) {
	m := fillMap(1000)
	i := 0
	for ; i < c.N; i++ {
		x, ok := m.Get(i)
		c.Assert(ok, Equals, true)
		c.Assert(x, Equals, i*2)
	}
	c.Assert(m.Len(), Equals, i+1)
}

func (s *MapBenchSuite) BenchmarkGet10000(c *C) {
	m := fillMap(10000)
	i := 0
	for ; i < c.N; i++ {
		x, ok := m.Get(i)
		c.Assert(ok, Equals, true)
		c.Assert(x, Equals, i*2)
	}
	c.Assert(m.Len(), Equals, i+1)
}

func (s *MapBenchSuite) BenchmarkGet100000(c *C) {
	m := fillMap(100000)
	i := 0
	for ; i < c.N; i++ {
		x, ok := m.Get(i)
		c.Assert(ok, Equals, true)
		c.Assert(x, Equals, i*2)
	}
	c.Assert(m.Len(), Equals, i+1)
}
