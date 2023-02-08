package file

import "skullpt-lib/data_types"

var Create data_types.ActionCommand = returnAction(data_types.Specifics{})

func returnAction(cfg data_types.Specifics) data_types.ActionCommand {
	return actionFileCreate
}

func actionFileCreate(data_types.Specifics) error {
	// do stuff
	return nil
}
