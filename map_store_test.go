package skiplist

import (
	"compress/gzip"
	"fmt"
	. "launchpad.net/gocheck"
	"os"
)

type MapStoreSuite struct{}

var _ = Suite(&MapStoreSuite{})
var filename = "storetest.slm"

func (s *MapStoreSuite) TestPersistInts(c *C) {
	m := NewMap(compareInt64s)
	m.Put(int64(1), int64(2))
	m.Put(int64(3), int64(4))
	f, err := os.Create(filename)
	c.Assert(err, IsNil)
	m.Persist(f, Int64Int64Record{})
}

func (s *MapStoreSuite) TestMergeEmptyInts(c *C) {
	m := NewMap(compareInt64s)
	m.Put(int64(1), int64(2))
	m.Put(int64(3), int64(4))
	f, err := os.Create(filename)
	c.Assert(err, IsNil)
	err = m.Persist(f, Int64Int64Record{})
	c.Assert(err, IsNil)
	f.Close()

	f, err = os.Open(filename)
	m = NewMap(compareInt64s)
	err = m.Merge(f, Int64Int64Record{})
	c.Assert(err, IsNil)
	f.Close()
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
	f, err := os.Create(filename)
	c.Assert(err, IsNil)
	m.Persist(f, StringStringRecord{})
}

func (s *MapStoreSuite) TestMergeEmptyStrings(c *C) {
	m := NewMap(compareStrings)
	m.Put("a", "bc")
	m.Put("d", "éf")
	f, err := os.Create(filename)
	c.Assert(err, IsNil)
	m.Persist(f, StringStringRecord{})
	m = NewMap(compareStrings)
	f, err = os.Open(filename)
	c.Assert(err, IsNil)
	m.Merge(f, StringStringRecord{})
	x, ok := m.Get("a")
	c.Assert(ok, Equals, true)
	c.Assert(x, Equals, "bc")
	x, ok = m.Get("d")
	c.Assert(ok, Equals, true)
	c.Assert(x, Equals, "éf")
	c.Assert(m.Len(), Equals, 2)
}

func benchmarkMergeN(n int, c *C) {
	c.StopTimer()
	m := NewMap(compareStrings)
	for i := 0; i < n; i++ {
		m.Put(fmt.Sprintf("%d", i), fmt.Sprintf("%d%d%d%d", i, i, i, i))
	}
	f, err := os.Create(filename)
	c.Assert(err, IsNil)
	err = m.Persist(f, StringStringRecord{})
	c.Assert(err, IsNil)
	f.Close()
	m = NewMap(compareStrings)
	c.StartTimer()
	for i := 0; i < c.N; i++ {
		f, err := os.Open(filename)
		c.Assert(err, IsNil)
		err = m.Merge(f, StringStringRecord{})
		c.Assert(err, IsNil)
		f.Close()
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

func benchmarkMergeZipN(n int, c *C) {
	c.StopTimer()
	m := NewMap(compareStrings)
	for i := 0; i < n; i++ {
		m.Put(fmt.Sprintf("%d", i), fmt.Sprintf("%d%d%d%d", i, i, i, i))
	}
	f, err := os.Create(filename)
	c.Assert(err, IsNil)
	zf := gzip.NewWriter(f)
	err = m.Persist(zf, StringStringRecord{})
	c.Assert(err, IsNil)
	zf.Close()
	f.Close()
	m = NewMap(compareStrings)
	c.StartTimer()
	for i := 0; i < c.N; i++ {
		f, err := os.Open(filename)
		c.Assert(err, IsNil)
		zf, err := gzip.NewReader(f)
		c.Assert(err, IsNil)
		err = m.Merge(zf, StringStringRecord{})
		c.Assert(err, IsNil)
		zf.Close()
		f.Close()
		c.Assert(m.Len(), Equals, n)
		/*
			for i := 0; i < n; i++ {
				x, ok := m.Get(fmt.Sprintf("%d", i))
				c.Assert(ok, Equals, true)
				c.Assert(x, Equals, fmt.Sprintf("%d%d%d%d", i, i, i, i))
			}
		*/
	}
}

func (s *MapStoreSuite) BenchmarkMergeZip1000(c *C) {
	benchmarkMergeZipN(1000, c)
}
func (s *MapStoreSuite) BenchmarkMergeZip10000(c *C) {
	benchmarkMergeZipN(10000, c)
}
func (s *MapStoreSuite) BenchmarkMergeZip100000(c *C) {
	benchmarkMergeZipN(100000, c)
}
func (s *MapStoreSuite) BenchmarkMergeZip1000000(c *C) {
	benchmarkMergeZipN(1000000, c)
}
