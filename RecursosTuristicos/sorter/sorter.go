package sorter

import (
	//"fmt"
	"sort"
	m "RecursosTuristicos/reader"
)

type By func(p1, p2 *m.RecursoD) bool

func (by By) Sort(recursos []m.RecursoD) {
	ps := &recursoSorter{
		recursos: recursos,
		by:      by,
	}
	sort.Sort(ps)
}

type recursoSorter struct {
	recursos []m.RecursoD
	by      func(p1, p2 *m.RecursoD) bool
}

func (s *recursoSorter) Len() int {
	return len(s.recursos)
}

func (s *recursoSorter) Swap(i, j int) {
	s.recursos[i], s.recursos[j] = s.recursos[j], s.recursos[i]
}

func (s *recursoSorter) Less(i, j int) bool {
	return s.by(&s.recursos[i], &s.recursos[j])
}