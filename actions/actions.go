package actions

import (
	"skullpt-lib/actions/file"
	"skullpt-lib/data_types"
)

func CollectActions() data_types.AvailableActions {
	return data_types.AvailableActions{
		"file": file.Register(),
	}
}
