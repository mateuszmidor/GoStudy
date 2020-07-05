package airport

type AirportID int

const NilAiportID = AirportID(-1)

type Airport struct {
	code string // KRK
	name string // Balica Krakow Airport
}

func NewAirport(code, name string) Airport {
	return Airport{code, name}
}

func (a *Airport) Code() string {
	return a.code
}
