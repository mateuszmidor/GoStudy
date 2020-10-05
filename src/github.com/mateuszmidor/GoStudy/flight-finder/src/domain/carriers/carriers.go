package carriers

import "sort"

// Carriers holds mapping for id-carrier
type Carriers []Carrier

// Get returns Carrier by CarrierID
func (c Carriers) Get(id ID) Carrier {
	return c[id]
}

// GetByCode returns CarrierID of given carrier
// Precondition: carriers are sorted ascending
func (c Carriers) GetByCode(code string) ID {
	ge := func(i int) bool {
		return c[i].code >= code
	}

	foundIndex := sort.Search(len(c), ge)

	if c[foundIndex].code != code {
		return NullID
	}

	return ID(foundIndex)
}
