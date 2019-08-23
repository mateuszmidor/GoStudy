package cmds

import "hexagons/tuner/domain"
import "hexagons/tuner/infrastructure"

type UpdateSubscriptionCmd struct {
	subscription domain.Subscription
}

func NewUpdateSubscriptionCmd(subscription domain.Subscription) UpdateSubscriptionCmd {
	return UpdateSubscriptionCmd{subscription}
}

func (cmd UpdateSubscriptionCmd) Execute(state *domain.TunerState, ports *infrastructure.OuterWorldPorts) {
	state.Subscription = cmd.subscription
	ports.UIPort.UpdateSubscription(cmd.subscription)
}
