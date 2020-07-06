package airport

// Airports holds mapping for id-airport
type Airports []Airport

// Get returns Airport by AirportID
func (a Airports) Get(id AirportID) Airport {
	return a[id]
}

// GetByCode returns AirportID of given airport
// Precondition: airports are sorted ascending
func (a Airports) GetByCode(code string) AirportID {
	first := 0
	last := len(a)
	count := last - first

	for count > 0 {
		i := first
		step := count / 2
		i += step
		if a[i].code < code {
			first = i + 1
			count -= step + 1
		} else {
			count = step
		}
	}
	return AirportID(first)
}
