package clients

// Rope belongs to the public interface of Ropeworks.
type Rope struct{}

// Ropeworks is the public interface of ropeworks module.
type Ropeworks interface {
	GetRopes(count int) ([]Rope, error)
}
