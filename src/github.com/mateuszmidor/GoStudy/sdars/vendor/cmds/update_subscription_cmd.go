package cmds

import "sdars"

type UpdateSubscriptionCmd struct {
	active bool
}

func NewUpdateSubscriptionCmd(active bool) *UpdateSubscriptionCmd {
	return &UpdateSubscriptionCmd{active}
}

func (cmd UpdateSubscriptionCmd) Execute(tuner *sdars.Tuner) {
	tuner.State.ActiveSubscription = cmd.active
	tuner.ClusterPort.UpdateSubscription(cmd.active)
}