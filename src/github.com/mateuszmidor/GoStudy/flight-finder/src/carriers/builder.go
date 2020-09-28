package carriers

import "sort"

// Builder creates sorted collection of carriers
type Builder struct {
	carriers Carriers
}

// Append adds new carrier at the collection end
func (b *Builder) Append(code string) {
	b.carriers = append(b.carriers, Carrier{code})
}

// Build returns sorted collection of carriers
func (b *Builder) Build() Carriers {
	less := func(i, j int) bool {
		return b.carriers[i].code < b.carriers[j].code
	}

	sort.Slice(b.carriers, less)
	return b.carriers
}
