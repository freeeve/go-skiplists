package set_test

import (
	"github.com/JnBrymN/GoSkipList"
	. "launchpad.net/gocheck"
	"testing"
)

type SetSuite struct{}

var _ = Suite(&DriverSuite{})

// gocheck link to go test - only needs doing once for package
func Test(t *testing.T) {
	TestingT(t)
}

func compareInts(a, b int) bool { return a < b }

func (s *CSLSuite) TestAdd(c *C) {
	set := NewSet(compareInts)
	set.Add(1)
}

func (s *CSLSuite) TestLen(c *C) {
	set := NewSet(compareInts)
	c.Assert(set.Len(), Equals, 0)
	set.Add(1)
	c.Assert(set.Len(), Equals, 1)
	et.Add(2)
	c.Assert(set.Len(), Equals, 2)
	set.Add(2)
	c.Assert(set.Len(), Equals, 2)
}

func (s *CSLSuite) TestContains(c *C) {
	et := New(compareInts)
	c.Assert(set.Contains(1), Equals, false)
	set.Add(1)
	c.Assert(set.Contains(1), Equals, true)
}
