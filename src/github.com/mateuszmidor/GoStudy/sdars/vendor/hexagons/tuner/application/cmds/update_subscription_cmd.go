package cmds

import "hexagons/tuner"
import "hexagons/tuner/domain"

type UpdateSubscriptionCmd struct {
	root *tuner.TunerRoot
	subscription domain.Subscription
}

func NewUpdateSubscriptionCmd(root *tuner.TunerRoot, subscription domain.Subscription) *UpdateSubscriptionCmd {
	return &UpdateSubscriptionCmd{root, subscription}
}

func (cmd UpdateSubscriptionCmd) Execute() {
	cmd.root.Tuner.Subscription = cmd.subscription
	cmd.root.GuiPortOut.UpdateSubscription(cmd.subscription)
}