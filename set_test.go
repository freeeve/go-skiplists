package skiplist

import (
	. "gopkg.in/check.v1"
)

type SetSuite struct{}

var _ = Suite(&SetSuite{})

func (s *SetSuite) TestAdd(c *C) {
	set := NewSortedSet(compareInts)
	set.Add(1)
}

func (s *SetSuite) TestCardinality(c *C) {
	set := NewSortedSet(compareInts)
	c.Assert(set.Cardinality(), Equals, 0)
	set.Add(1)
	c.Assert(set.Cardinality(), Equals, 1)
	set.Add(2)
	c.Assert(set.Cardinality(), Equals, 2)
	set.Add(2)
	c.Assert(set.Cardinality(), Equals, 2)
}

func (s *SetSuite) TestContains(c *C) {
	set := NewSortedSet(compareInts)
	c.Assert(set.Contains(1), Equals, false)
	set.Add(1)
	c.Assert(set.Contains(1), Equals, true)
}
