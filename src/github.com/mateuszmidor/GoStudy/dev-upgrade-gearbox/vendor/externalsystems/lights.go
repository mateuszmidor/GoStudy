package externalsystems

// Lights represents data related to car lights
type Lights struct {
	position *int
}

// NewLights is constructor
func NewLights(position *int) Lights {
	return Lights{position}
}

// GetPosition getter
// null - brak opcji w samochodzie
// 1-3 - w dół
// 7-10 - w górę
func (l Lights) GetPosition() *int {
	return l.position
}
