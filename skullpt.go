package skullpt_lib

import (
	"encoding/json"
	"fmt"
	"os"
	"skullpt-lib/actions"
	"skullpt-lib/data_types"
)

var registerdActions = data_types.AvailableActions{}

func CreateActionList() (data_types.QueuedCommands, error) {
	file, err := os.ReadFile("test_data/example_config.json")
	if err != nil {
		return data_types.QueuedCommands{}, err
	}

	gallerys, err := parseGallery(file)
	if err != nil {
		return data_types.QueuedCommands{}, err
	}

	endpointCommands := make(map[string][]data_types.ActionCommand)

	for _, gallery := range gallerys {
		cmds, err := convertActionToCommands(gallery.Actions)
		if err != nil {
			return data_types.QueuedCommands{}, err
		}
		for _, endpoint := range gallery.Endpoints {
			endpointCommands[endpoint] = cmds
		}
	}
}

func convertActionToCommands(config []data_types.Action) ([]data_types.ActionCommand, error) {
	convertedCommands := []data_types.ActionCommand{}

	for _, action := range config {
		chosenAction := registerdActions[action.AtionType]
		chosenCmd, err := chosenAction(action.Specifics)
		if err != nil {
			return nil, err
		}
		convertedCommands = append(convertedCommands, chosenCmd)
	}

	return convertedCommands, nil
}

func parseGallery(galleryBytes []byte) ([]data_types.Gallery, error) {

	var gallery []data_types.Gallery

	if err := json.Unmarshal(galleryBytes, &gallery); err != nil {
		return []data_types.Gallery{}, fmt.Errorf("error reading gallery config: %s", err)
	}
	return gallery, nil
}

func registerActions() {
	registerdActions = actions.CollectActions()
}

func init() {
	registerActions()
}
