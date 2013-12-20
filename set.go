package skiplist

// Set is the struct that holds info about the set
type Set struct {
	comp  func(a, b interface{}) bool
	first *setElement
}

// setElement holds an element of the set
type setElement struct {
	key  interface{}
	next *setElement
}

// NewSet returns a new empty set
func NewSet(f func(a, b interface{}) bool) *Set {
	return &Set{f, nil}
}

// Add adds an element to the set, doing nothing if it exists
// returns true if it existed, or false if it didn't exist
func (set *Set) Add(k interface{}) bool {
	// adding the first element
	if set.first == nil {
		set.first = &setElement{k, nil}
		return false
	}
	// adding something less than the first element, insert
	if set.comp(k, set.first.key) {
		set.first = &setElement{k, set.first}
		return false
	}
	var prev *setElement = nil
	e := set.first
	for e != nil {
		// if they are equal, do nothing
		if set.comp(k, e.key) == set.comp(e.key, k) {
			return true
		}
		// if inspected val is greater than k, insert
		if set.comp(k, e.key) {
			prev.next = &setElement{k, e}
			return false
		}
		// if we hit the end, insert
		if e.next == nil {
			e.next = &setElement{k, nil}
			return false
		}
		prev = e
		e = e.next
	}
	// this should never happen
	return false
}

// Len returns the length of the set
func (set *Set) Len() int {
	count := 0
	e := set.first
	for e != nil {
		e = e.next
		count++
	}
	return count
}

// Contains returns true if the key exists
func (m *Set) Contains(k interface{}) bool {
	e := m.first
	for e != nil {
		if m.comp(k, e.key) == m.comp(e.key, k) {
			return true
		}
		e = e.next
	}
	return false
}
