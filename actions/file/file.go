package file

import (
	"skullpt-lib/actions"
)

func Register() actions.AvailableCommands {
	return actions.AvailableCommands{
		"create": Create,
	}
}
