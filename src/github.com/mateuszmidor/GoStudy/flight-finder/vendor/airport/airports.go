package airport

// Airports holds mapping for id-airport
type Airports []Airport

// Get returns Airport by AirportID
func (a Airports) Get(id AirportID) Airport {
	return a[id]
}

// GetByCode returns AirportID of given airport
func (a Airports) GetByCode(code string) AirportID {

	// todo: binary search assuming airports are sorted by code
	for i := 0; i < len(a); i++ {
		if a[i].code == code {
			return AirportID(i)
		}
	}
	return NilAiportID
}
