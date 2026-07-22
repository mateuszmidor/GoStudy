package fundaccount

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"bank-account-persistent/events"
)

func Test_FundAccount_NotCreated_Error(t *testing.T) {
	// given we start with no account and want to fund it with 6 dollars
	state := accountState{created: false}
	cmd := FundAccount{Dollars: 6}

	// when we try to change state
	newEvents, err := decide(state, cmd)

	// then we receive error for funding non-created account and no new events
	assert.Error(t, err)
	assert.Empty(t, newEvents)
}

func Test_FundAccount_Created_Success(t *testing.T) {
	// given we start with created account and want to fund it with 6 dollars
	state := accountState{created: true}
	cmd := FundAccount{Dollars: 6}

	// when we try to change state
	newEvents, err := decide(state, cmd)

	// then we receive no error and AccountFundedEvent
	assert.NoError(t, err)
	assert.Contains(t, newEvents, &events.AccountFunded{Dollars: 6})
}
