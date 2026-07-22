package getbalance

import (
	"bank-account-persistent/events"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/terraskye/eventsourcing"
)

func Test_GetBalance_Success(t *testing.T) {
	// given that we start with balance of 3 dollars and deposit 5 dollars
	state := accountState{Dollars: 3}
	env := &eventsourcing.Envelope{Event: &events.AccountFunded{Dollars: 5}}

	// when we evaluate current balance
	newState := evolve(state, env)

	// then we get current balance of 8 dollars
	assert.Equal(t, uint(8), newState.Dollars)
}
