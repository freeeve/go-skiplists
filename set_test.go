package main

import (
	. "launchpad.net/gocheck"
)

type SetSuite struct{}

var _ = Suite(&SetSuite{})

func (s *SetSuite) TestAdd(c *C) {
	set := NewSet(compareInts)
	set.Add(1)
}

func (s *SetSuite) TestLen(c *C) {
	set := NewSet(compareInts)
	c.Assert(set.Len(), Equals, 0)
	set.Add(1)
	c.Assert(set.Len(), Equals, 1)
	set.Add(2)
	c.Assert(set.Len(), Equals, 2)
	set.Add(2)
	c.Assert(set.Len(), Equals, 2)
}

func (s *SetSuite) TestContains(c *C) {
	set := NewSet(compareInts)
	c.Assert(set.Contains(1), Equals, false)
	set.Add(1)
	c.Assert(set.Contains(1), Equals, true)
}
