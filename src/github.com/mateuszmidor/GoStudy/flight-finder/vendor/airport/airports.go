package airport

import "sort"

// Airports holds mapping for id-airport
type Airports []Airport

// Get returns Airport by ID
func (a Airports) Get(id ID) Airport {
	return a[id]
}

// GetByCode returns ID of given airport
// Precondition: airports are sorted ascending
func (a Airports) GetByCode(code string) ID {
	ge := func(i int) bool {
		return a[i].code >= code
	}

	foundIndex := sort.Search(len(a), ge)

	if a[foundIndex].code != code {
		return NullID
	}

	return ID(foundIndex)
}
