package clients

// Sail belongs to the public interface of Sailworks.
type Sail struct{}

type Sailworks interface {
	GetSails(count int) ([]Sail, error)
}
