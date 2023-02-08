package actions

import "skullpt-lib/actions/file"



func CollectActions() AvailableActions {
	avail := AvailableActions{
		"file": file.Register(),
	}
}
