package skiplist

import (
	"fmt"
	. "launchpad.net/gocheck"
)

type MapStoreSuite struct{}

var _ = Suite(&MapStoreSuite{})

func (s *MapStoreSuite) TestPersistInts(c *C) {
	m := NewMap(compareInt64s)
	m.Put(int64(1), int64(2))
	m.Put(int64(3), int64(4))
	m.Persist("storetest.slm", Int64Int64Record{})
}

func (s *MapStoreSuite) TestMergeEmptyInts(c *C) {
	m := NewMap(compareInt64s)
	m.Put(int64(1), int64(2))
	m.Put(int64(3), int64(4))
	err := m.Persist("storetest.slm", Int64Int64Record{})
	c.Assert(err, IsNil)
	m = NewMap(compareInt64s)
	err = m.Merge("storetest.slm", Int64Int64Record{})
	c.Assert(err, IsNil)
	x, ok := m.Get(int64(1))
	c.Assert(ok, Equals, true)
	c.Assert(x, Equals, int64(2))
	x, ok = m.Get(int64(3))
	c.Assert(ok, Equals, true)
	c.Assert(x, Equals, int64(4))
	c.Assert(m.Len(), Equals, 2)
}

func (s *MapStoreSuite) TestPersistStrings(c *C) {
	m := NewMap(compareStrings)
	m.Put("a", "bc")
	m.Put("d", "ef")
	m.Persist("storetest.slm", StringStringRecord{})
}

func (s *MapStoreSuite) TestMergeEmptyStrings(c *C) {
	m := NewMap(compareStrings)
	m.Put("a", "bc")
	m.Put("d", "ef")
	m.Persist("storetest.slm", StringStringRecord{})
	m = NewMap(compareStrings)
	m.Merge("storetest.slm", StringStringRecord{})
	x, ok := m.Get("a")
	c.Assert(ok, Equals, true)
	c.Assert(x, Equals, "bc")
	x, ok = m.Get("d")
	c.Assert(ok, Equals, true)
	c.Assert(x, Equals, "ef")
	c.Assert(m.Len(), Equals, 2)
}

func benchmarkMergeN(n int, c *C) {
	c.StopTimer()
	m := NewMap(compareStrings)
	for i := 0; i < n; i++ {
		m.Put(fmt.Sprintf("%d", i), fmt.Sprintf("%d%d%d%d", i, i, i, i))
	}
	err := m.Persist("storetest.slm", StringStringRecord{})
	c.Assert(err, IsNil)
	c.StartTimer()
	for i := 0; i < c.N; i++ {
		err := m.Merge("storetest.slm", StringStringRecord{})
		c.Assert(err, IsNil)
		c.Assert(m.Len(), Equals, n)
	}
}

func (s *MapStoreSuite) BenchmarkMerge1000(c *C) {
	benchmarkMergeN(1000, c)
}
func (s *MapStoreSuite) BenchmarkMerge10000(c *C) {
	benchmarkMergeN(10000, c)
}
func (s *MapStoreSuite) BenchmarkMerge100000(c *C) {
	benchmarkMergeN(100000, c)
}
func (s *MapStoreSuite) BenchmarkMerge1000000(c *C) {
	benchmarkMergeN(1000000, c)
}
