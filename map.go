package skiplist

// Map is the struct to hold the details of a map
type Map struct {
	comp  func(a, b interface{}) bool
	first *mapElement
}

// mapElement is the struct to hold elements of the map
type mapElement struct {
	key  interface{}
	val  interface{}
	next *mapElement
}

// NewMap creates a new empty map, it takes a
// comparison function that should implement Less
func NewMap(less func(a, b interface{}) bool) *Map {
	return &Map{less, nil}
}

// Put takes a key and value, and puts the value
// in the map for the key, replacing an existing value.
// returns true if it overwrites, false if it inserts a new key/value pair
func (m *Map) Put(k interface{}, v interface{}) bool {
	// adding the first element
	if m.first == nil {
		m.first = &mapElement{k, v, nil}
		return false
	}
	// adding something less than the first element, insert
	if m.comp(k, m.first.key) {
		m.first = &mapElement{k, v, m.first}
		return false
	}
	var prev *mapElement = nil
	e := m.first
	for e != nil {
		// if they are equal, overwrite
		if m.comp(k, e.key) == m.comp(e.key, k) {
			e.val = v
			return true
		}
		// if inspected val is greater than k, insert
		if m.comp(k, e.key) {
			prev.next = &mapElement{k, v, e}
			return false
		}
		// if we hit the end, insert
		if e.next == nil {
			e.next = &mapElement{k, v, nil}
			return false
		}
		prev = e
		e = e.next
	}
	// this should never happen
	return false
}

// Len returns the length of a Map
func (m *Map) Len() int {
	count := 0
	e := m.first
	for e != nil {
		e = e.next
		count++
	}
	return count
}

// Get returns the value for a key, and true if it finds the key,
// false otherwise
func (m *Map) Get(k interface{}) (interface{}, bool) {
	e := m.first
	for e != nil {
		if m.comp(k, e.key) == m.comp(e.key, k) {
			return e.val, true
		}
		e = e.next
	}
	return nil, false
}
