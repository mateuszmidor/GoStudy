package segment

import (
	"airport"
	"carrier"
	"sort"
)

// Builder creates sorted collection of segments
type Builder struct {
	segments Segments
	airports airport.Airports
	carriers carrier.Carriers
}

// NewBuilder is constructor
func NewBuilder(airports airport.Airports, carriers carrier.Carriers) Builder {
	return Builder{Segments{}, airports, carriers}
}

// Append adds new segment to the collection
func (b *Builder) Append(from, to, carrier string) {
	fromID := b.airports.GetByCode(from)
	toID := b.airports.GetByCode(to)
	carrierID := b.carriers.GetByCode(carrier)
	b.segments = append(b.segments, Segment{fromID, toID, carrierID})
}

// Build returns sorted collection of segments
func (b *Builder) Build() Segments {
	less := func(i, j int) bool {
		if b.segments[i].from != b.segments[j].from {
			return b.segments[i].from < b.segments[j].from
		}
		return b.segments[i].to < b.segments[j].to
	}

	sort.Slice(b.segments, less)
	return b.segments
}