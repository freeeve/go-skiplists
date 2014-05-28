package skiplist

import (
	"math"
	"math/rand"
)

// SortedSet is the struct to hold the details of a SortedSet
// backed by a skiplist
type SortedSet struct {
	less      func(a, b interface{}) bool
	head      []*setElement
	length    int
	maxLevels int
	r         *rand.Rand
}

// setElement is the struct to hold elements of the map
type setElement struct {
	val  interface{}
	next []*setElement
}

// NewSortedSet creates a new empty set, it takes a
// comparison function that should implement Less
func NewSortedSet(less func(a, b interface{}) bool) *SortedSet {
	return &SortedSet{
		less:      less,
		maxLevels: 64,
		head:      make([]*setElement, 64),
		r:         rand.New(rand.NewSource(123123)),
	}
}

func newSetElement(v interface{}, levels int) *setElement {
	return &setElement{v, make([]*setElement, levels)}
}

func (ss *SortedSet) randomLevels() int {
	level := int(math.Log(1.0-ss.r.Float64()) / math.Log(1.0-0.5))
	if level >= ss.maxLevels {
		level = ss.maxLevels
	}
	if level == 0 {
		level++
	}
	return level
}

// Add puts a value in the set.
// returns true if it overwrites, false if it inserts a new value
func (ss *SortedSet) Add(v interface{}) bool {
	var backPointer = make([]*setElement, 64)
	// zeroing this causes the compiler to not allocate memory each time
	// for a 20-30% boost in speed
	for i := 0; i < 64; i++ {
		backPointer[i] = nil
	}
	for level := ss.maxLevels - 1; level >= 0; level-- {
		var e *setElement = nil
		if level+1 == ss.maxLevels || backPointer[level+1] == nil {
			e = ss.head[level]
		} else {
			e = backPointer[level+1]
		}
		for e != nil {
			// if they are equal, overwrite?
			if ss.less(v, e.val) == ss.less(e.val, v) {
				return true
			}
			// if inspected val is greater than k, go back and down a level
			if ss.less(v, e.val) {
				break
			}
			backPointer[level] = e
			e = e.next[level]
		}
	}
	// create new element
	e := newSetElement(v, ss.randomLevels())

	// connect new element up with backPointer
	for level := 0; level < len(e.next); level++ {
		if backPointer[level] == nil {
			e.next[level] = ss.head[level]
			ss.head[level] = e
		} else {
			e.next[level] = backPointer[level].next[level]
			backPointer[level].next[level] = e
		}
	}
	//log.Println(e, backPointer)

	ss.length++
	return false
}

// Cardinality returns the length of a SortedSet
func (ss *SortedSet) Cardinality() int {
	// TODO why is this busted
	//ret := ss.length
	e := ss.head[0]
	ret := 0
	for e != nil {
		ret++
		e = e.next[0]
	}
	return ret
}

// Contains returns true if it finds the value, false otherwise
func (ss *SortedSet) Contains(v interface{}) bool {
	var backPointer = make([]*setElement, 64)
	// zeroing this causes the compiler to not allocate memory each time
	// for a 20-30% boost in speed
	for i := 0; i < 64; i++ {
		backPointer[i] = nil
	}
	for level := ss.maxLevels - 1; level >= 0; level-- {
		var e *setElement = nil
		if level+1 == ss.maxLevels || backPointer[level+1] == nil {
			e = ss.head[level]
		} else {
			e = backPointer[level+1]
		}
		for e != nil {
			// if they are equal, return val
			if ss.less(v, e.val) == ss.less(e.val, v) {
				return true
			}
			// if inspected val is greater than v, go back and down a level
			if ss.less(v, e.val) {
				break
			}
			backPointer[level] = e
			e = e.next[level]
		}
	}
	return false
}

// Remove removes the element for a value,
// returns true if it found and removed, false otherwise
func (ss *SortedSet) Remove(v interface{}) bool {
	var backPointer = make([]*setElement, 64)
	// zeroing this causes the compiler to not allocate memory each time
	// for a 20-30% boost in speed
	for i := 0; i < 64; i++ {
		backPointer[i] = nil
	}
	for level := ss.maxLevels - 1; level >= 0; level-- {
		var e *setElement = nil
		if level+1 == ss.maxLevels || backPointer[level+1] == nil {
			e = ss.head[level]
		} else {
			e = backPointer[level+1]
		}
		for e != nil {
			// if they are equal, remove and return true
			if level == 0 && ss.less(v, e.val) == ss.less(e.val, v) {
				for level := 0; level < len(e.next); level++ {
					if backPointer[level] == nil {
						ss.head[level] = e.next[level]
					} else {
						backPointer[level].next[level] = e.next[level]
					}
				}

				ss.length--
				return true
			}
			if ss.less(v, e.val) == ss.less(e.val, v) {
				break
			}
			// if inspected val is greater than k, go back and down a level
			if ss.less(v, e.val) {
				break
			}
			backPointer[level] = e
			e = e.next[level]
		}
	}
	return false
}
