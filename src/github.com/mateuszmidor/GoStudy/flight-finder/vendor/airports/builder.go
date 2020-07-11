package airports

import "sort"

// Builder creates sorted collection of airports
type Builder struct {
	airports Airports
}

// Append adds new airport at the collection end
func (b *Builder) Append(code, name string) {
	b.airports = append(b.airports, Airport{code, name})
}

// Build returns sorted collection of airports
func (b *Builder) Build() Airports {
	less := func(i, j int) bool {
		return b.airports[i].code < b.airports[j].code
	}

	sort.Slice(b.airports, less)
	return b.airports
}
