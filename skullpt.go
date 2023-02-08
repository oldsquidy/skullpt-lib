package skullpt_lib

import (
	"encoding/json"
	"fmt"
)

type Gallery struct {
	Name      string   `json: "name"`
	Endpoints []string `json: "endpoints"`
	Actions   []Action `json: "actions"`
}

// Action holds the action config
type Action struct {
	actionType string    `json: "action_type"`
	specifics  Specifics `json: "specifics"`
}

// Specifics is the config for the action
type Specifics map[string]any

// ActionCommand is the function signature all actions needs to do
type ActionCommand func(cfg Specifics) error

// AvailableActions will hold all registerd commands for an action
// e.g. create = fn for creating
type AvailableCommands map[string]ActionCommand

type AvailableActions map[string]AvailableCommands

func ParseGallery(galleryBytes []byte) ([]Gallery, error) {

	var gallery []Gallery

	if err := json.Unmarshal(galleryBytes, &gallery); err != nil {
		return []Gallery{}, fmt.Errorf("error reading gallery config: %s", err)
	}
	return gallery, nil
}

func BuildActions(gallery Gallery)
