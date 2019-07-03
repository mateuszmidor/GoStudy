package cmds

import "hexagons/tuner/domain"
import "hexagons/tuner/infrastructure"

type UpdateSubscriptionCmd struct {
	subscription domain.Subscription
}

func NewUpdateSubscriptionCmd(subscription domain.Subscription) *UpdateSubscriptionCmd {
	return &UpdateSubscriptionCmd{subscription}
}

func (cmd UpdateSubscriptionCmd) Execute(tuner *domain.Tuner, ports *infrastructure.Ports) {
	tuner.Subscription = cmd.subscription
	ports.UiPortOut.UpdateSubscription(cmd.subscription)
}