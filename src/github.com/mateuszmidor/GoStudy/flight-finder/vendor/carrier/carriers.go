package carrier

// Carriers holds mapping for id-carrier
type Carriers []Carrier

// Get returns Carrier by CarrierID
func (c Carriers) Get(id ID) Carrier {
	return c[id]
}

// GetByCode returns CarrierID of given carrier
// Precondition: carriers are sorted ascending
func (c Carriers) GetByCode(code string) ID {
	first := 0
	last := len(c)
	count := last - first

	for count > 0 {
		i := first
		step := count / 2
		i += step
		if c[i].code < code {
			first = i + 1
			count -= step + 1
		} else {
			count = step
		}
	}
	return ID(first)
}
