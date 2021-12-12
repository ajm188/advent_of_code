// Package sets provides set-like collections based on struct maps.
//
// Each set type has the interface:
//	type TSet map[T]struct{}{}
//	type Set interface {
//		Insert(elems ...T)
//		Has(key T) bool
//		Len() int
//		Items() []T
//		Copy() *TSet
// 	}
package sets

type Strings map[string]struct{}

func NewStrings(elems ...string) *Strings {
	set := Strings(map[string]struct{}{})
	set.Insert(elems...)
	return &set
}

func (set *Strings) Insert(elems ...string) {
	for _, s := range elems {
		(*set)[s] = struct{}{}
	}
}

func (set *Strings) Has(k string) bool {
	_, ok := (*set)[k]
	return ok
}

func (set *Strings) Len() int { return len(*set) }

func (set *Strings) Items() []string {
	items := make([]string, 0, set.Len())
	for k := range *set {
		items = append(items, k)
	}

	return items
}

func (set *Strings) Copy() *Strings {
	return NewStrings(set.Items()...)
}
