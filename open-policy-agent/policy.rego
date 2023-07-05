package age_authorization

import future.keywords

default can_drink := false # can_drink will be available as policy evaluation output

can_drink if {
	input.age >= 18
}
