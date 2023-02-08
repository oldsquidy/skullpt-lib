package skullpt_lib

import (
	"encoding/json"
	"fmt"
	"os"
	"skullpt-lib/actions"
	"skullpt-lib/data_types"
)

var registerdActions = data_types.AvailableActions{}

func CreateActionList() ([]data_types.QueuedCommands, error) {
	file, err := os.ReadFile("test_data/example_config.json")
	if err != nil {
		return []data_types.QueuedCommands{}, err
	}

	gallerys, err := parseGallery(file)
	if err != nil {
		return []data_types.QueuedCommands{}, err
	}

	queuedCommands := []data_types.QueuedCommands{}

	for _, gallery := range gallerys {
		convertedCommands, err := convertActionToCommands(gallery.Actions)
		if err != nil {
			return []data_types.QueuedCommands{}, err
		}
		for _, endpoint := range gallery.Endpoints {
			queuedCommands = append(queuedCommands, data_types.QueuedCommands{
				Endpoint: endpoint,
				Commands: convertedCommands,
			})
		}
	}
	return queuedCommands, nil
}

func convertActionToCommands(config map[string]data_types.Action) (map[string]data_types.ActionCommand, error) {
	convertedCommands := make(map[string]data_types.ActionCommand)

	for desc, action := range config {
		chosenAction := registerdActions[action.ActionType]
		chosenCmd, err := chosenAction(action.Specifics)
		if err != nil {
			return nil, err
		}
		convertedCommands[desc] = chosenCmd
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
