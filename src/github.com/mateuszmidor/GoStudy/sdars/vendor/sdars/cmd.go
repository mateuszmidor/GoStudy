package sdars

type Cmd interface {
	Execute(tuner *Tuner)
}