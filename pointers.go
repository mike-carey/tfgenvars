package tfgenvars

import (
	"sort"
)

type Pointer struct {
	Name string
	Position int
}

type Pointers []Pointer

func (p Pointers) Sort() {
	ps := &pointerSorter{
		Pointers: p,
	}

	sort.Sort(ps)
}

func (ps Pointers) Map() map[string]int {
	m := make(map[string]int, len(ps))
	for _, p := range ps {
		m[p.Name] = p.Position
	}
	return m
}

type pointerSorter struct {
	Pointers []Pointer
}

func (s *pointerSorter) Len() int {
	return len(s.Pointers)
}

func (s *pointerSorter) Swap(i, j int) {
	s.Pointers[i], s.Pointers[j] = s.Pointers[j], s.Pointers[i]
}

func (s *pointerSorter) Less(i, j int) bool {
	return (&s.Pointers[i]).Position < (&s.Pointers[j]).Position
}
