package nations

import (
	"sort"
)

// Builder creates sorted collection of nations
type Builder struct {
	nations Nations
}

// Append adds new airport at the collection end
func (b *Builder) Append(code, iso, currency, name string) {
	b.nations = append(b.nations, NewNation(code, iso, currency, name))
}

// Build returns sorted collection of nations
func (b *Builder) Build() Nations {
	less := func(i, j int) bool {
		return b.nations[i].code < b.nations[j].code
	}

	sort.Slice(b.nations, less)
	return b.nations
}
