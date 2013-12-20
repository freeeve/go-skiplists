package main

type Set struct {
	comp  func(a, b interface{}) bool
	first *setElement
}

type setElement struct {
	key  interface{}
	next *setElement
}

func NewSet(f func(a, b interface{}) bool) *Set {
	return &Set{f, nil}
}

func (set *Set) Add(k interface{}) {
	// adding the first element
	if set.first == nil {
		set.first = &setElement{k, nil}
		return
	}
	// adding something less than the first element, insert
	if set.comp(k, set.first.key) {
		set.first = &setElement{k, set.first}
		return
	}
	var prev *setElement = nil
	e := set.first
	for e != nil {
		// if they are equal, do nothing
		if set.comp(k, e.key) == set.comp(e.key, k) {
			return
		}
		// if inspected val is greater than k, insert
		if set.comp(k, e.key) {
			prev.next = &setElement{k, e}
			return
		}
		// if we hit the end, insert
		if e.next == nil {
			e.next = &setElement{k, nil}
			return
		}
		prev = e
		e = e.next
	}
}

func (set *Set) Len() int {
	count := 0
	e := set.first
	for e != nil {
		e = e.next
		count++
	}
	return count
}

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
