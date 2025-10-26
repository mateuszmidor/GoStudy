package domain

type UiState struct {
	StationList        []string
	SubscriptionActive bool
}

func NewUiState() UiState {
	return UiState{}
}
