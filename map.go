package main

type Map struct {
	comp  func(a, b interface{}) bool
	first *mapElement
}

type mapElement struct {
	key  interface{}
	val  interface{}
	next *mapElement
}

func NewMap(f func(a, b interface{}) bool) *Map {
	return &Map{f, nil}
}

func (m *Map) Put(k interface{}, v interface{}) {
	// adding the first element
	if m.first == nil {
		m.first = &mapElement{k, v, nil}
		return
	}
	// adding something less than the first element, insert
	if m.comp(k, m.first.key) {
		m.first = &mapElement{k, v, m.first}
		return
	}
	var prev *mapElement = nil
	e := m.first
	for e != nil {
		// if they are equal, overwrite
		if m.comp(k, e.key) == m.comp(e.key, k) {
			e.val = v
			return
		}
		// if inspected val is greater than k, insert
		if m.comp(k, e.key) {
			prev.next = &mapElement{k, v, e}
			return
		}
		// if we hit the end, insert
		if e.next == nil {
			e.next = &mapElement{k, v, nil}
			return
		}
		prev = e
		e = e.next
	}
}

func (m *Map) Len() int {
	count := 0
	e := m.first
	for e != nil {
		e = e.next
		count++
	}
	return count
}

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
