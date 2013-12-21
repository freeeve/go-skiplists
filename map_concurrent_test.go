package skiplist

import (
	. "launchpad.net/gocheck"
	"log"
	"sync"
	"time"
)

type MapConcurrentSuite struct{}

var _ = Suite(&MapConcurrentSuite{})
var chunksize = 1000000

func writer(i int, m *Map, w *sync.WaitGroup) {
	for j := 0; j <= chunksize; j++ {
		m.Put(i, j)
		time.Sleep(1 * time.Nanosecond)
	}
	log.Println(i, "; writer done")
	w.Done()
}

func remover(i int, m *Map, w *sync.WaitGroup) {
	for j := 0; j < 10; j++ {
		m.Remove(i)
		//		log.Println("removed:", i)
		time.Sleep(1 * time.Nanosecond)
	}
	log.Println(i, "; remover done")
	w.Done()
}

// try to break the map...
// concurrently adding stuff to the same key
func (s *MapSuite) TestConcurPutOverwrite(c *C) {
	//c.Skip("only need to run when testing concurrency")
	log.Println("starting concurrent map test; writers should finish last else adjust parameters")
	m := NewMap(compareInts)
	m.Put(2, 3)
	m.Put(0, 1)
	var w sync.WaitGroup
	n := 3
	w.Add(n * 2)
	for i := 0; i < n; i++ {
		go remover(i, m, &w)
		go writer(i, m, &w)
	}
	w.Wait()
	x, _ := m.Get(0)
	c.Assert(x, Equals, chunksize)
	x, _ = m.Get(1)
	c.Assert(x, Equals, chunksize)
	x, _ = m.Get(2)
	c.Assert(x, Equals, chunksize)
	c.Assert(m.Len(), Equals, n)
}
