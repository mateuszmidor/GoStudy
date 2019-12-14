package cmds

import (
	"hexagons/tuner/domain"
	"hexagons/tuner/infrastructure"
	"sharedkernel"
)

type UpdateSubscriptionCmd struct {
	subscription sharedkernel.Subscription
}

func NewUpdateSubscriptionCmd(subscription sharedkernel.Subscription) UpdateSubscriptionCmd {
	return UpdateSubscriptionCmd{subscription}
}

func (cmd UpdateSubscriptionCmd) Execute(state *domain.TunerState, ports *infrastructure.OuterWorldPorts) {
	state.Subscription = cmd.subscription
	ports.UIPort.UpdateSubscription(cmd.subscription)
}
