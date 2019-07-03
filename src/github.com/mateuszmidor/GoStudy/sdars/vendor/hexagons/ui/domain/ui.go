package domain

type Ui struct {
	StationList []string
	SubscriptionActive bool
}

func NewUi() Ui {
	return Ui{}
}