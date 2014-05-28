package skiplist

import (
	"runtime"
	"sync"
	"time"
	. "gopkg.in/check.v1"
)

type MapConcurrentSuite struct{}

var _ = Suite(&MapConcurrentSuite{})
var chunksize = 100000

func writer(i int, m *Map, w *sync.WaitGroup) {
	for j := 0; j <= chunksize; j++ {
		m.Put(i, j)
		time.Sleep(1 * time.Nanosecond)
	}
	//	log.Println(i, "; writer done;", m.Mutex())
	w.Done()
}

func remover(i int, m *Map, w *sync.WaitGroup) {
	for j := 0; j < 1000; j++ {
		m.Remove(i)
		//		log.Println("removed:", i)
		time.Sleep(1 * time.Nanosecond)
	}
	//	log.Println(i, "; remover done", m.Mutex())
	w.Done()
}

// try to break the map...
// concurrently adding stuff and removing stuff
func (s *MapConcurrentSuite) TestConcurPutOverwrite(c *C) {
	//c.Skip("only need to run when testing concurrency")
	runtime.GOMAXPROCS(2)
	//log.Println("starting concurrent map test; writers should finish last else adjust parameters")
	m := NewMap(compareInts)
	var w sync.WaitGroup
	n := 10
	w.Add(n * 2)
	for i := 0; i < n; i++ {
		go remover(i, m, &w)
		go writer(i, m, &w)
	}
	w.Wait()
	for i := 0; i < n; i++ {
		m.Put(i, chunksize)
	}
	for i := 0; i < n; i++ {
		x, _ := m.Get(i)
		c.Assert(x, Equals, chunksize)
	}
	c.Assert(m.Len(), Equals, n)
}
