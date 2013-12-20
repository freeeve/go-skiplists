package main

import (
	. "launchpad.net/gocheck"
	"testing"
)

type MapSuite struct{}

var _ = Suite(&MapSuite{})

// gocheck link to go test - only needs doing once for whole lib
func Test(t *testing.T) {
	TestingT(t)
}

func compareInts(a, b interface{}) bool { return a.(int) < b.(int) }

func (s *MapSuite) TestNewMap(c *C) {
	m := NewMap(compareInts)
	if m == nil {
		c.Fatal("m is nil")
	}
}

func (s *MapSuite) TestPut(c *C) {
	m := NewMap(compareInts)
	m.Put(1, 1)
}

func (s *MapSuite) TestLen(c *C) {
	m := NewMap(compareInts)
	c.Assert(m.Len(), Equals, 0)
	m.Put(1, 1)
	c.Assert(m.Len(), Equals, 1)
	m.Put(2, 2)
	c.Assert(m.Len(), Equals, 2)
	m.Put(2, 3)
	c.Assert(m.Len(), Equals, 2)
}

func (s *MapSuite) TestGetEmpty(c *C) {
	m := NewMap(compareInts)
	i, ok := m.Get(1)
	c.Assert(i, IsNil)
	c.Assert(ok, Equals, false)
}

func (s *MapSuite) TestGetNotEmpty(c *C) {
	m := NewMap(compareInts)
	m.Put(1, 37)
	i, ok := m.Get(1)
	c.Assert(i, Equals, 37)
	c.Assert(ok, Equals, true)
}

func (s *MapSuite) TestPutOverwrite(c *C) {
	m := NewMap(compareInts)
	m.Put(1, 1)
	m.Put(1, 37)
	i, ok := m.Get(1)
	c.Assert(i, Equals, 37)
	c.Assert(ok, Equals, true)
}
