package nations

import "sort"

// Nations holds mapping for id-nation
type Nations []Nation

// Get returns Nation by ID
func (n Nations) Get(id ID) Nation {
	return n[id]
}

// GetByCode returns ID of given nation
// Precondition: nations are sorted ascending
func (n Nations) GetByCode(code string) ID {
	ge := func(i int) bool {
		return n[i].code >= code
	}

	foundIndex := sort.Search(len(n), ge)

	if foundIndex < 0 || foundIndex >= len(n) || n[foundIndex].code != code {
		return NullID
	}

	return ID(foundIndex)
}
