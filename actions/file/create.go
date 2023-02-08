package file

import (
	"skullpt-lib/actions"
)

var Create actions.ActionCommand = returnAction(actions.Specifics{})

func returnAction(cfg actions.Specifics) actions.ActionCommand {
	return actionFileCreate
}

func actionFileCreate(actions.Specifics) error {
	// do stuff
	return nil
}
