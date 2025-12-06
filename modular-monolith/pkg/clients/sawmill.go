package clients

// Plank belongs to the public interface of Sawmill.
type Plank struct{}

// Sawmill is the public interface of sawmill module.
type Sawmill interface {
	Run()
	GetPlanks(count int) ([]Plank, error)
}
