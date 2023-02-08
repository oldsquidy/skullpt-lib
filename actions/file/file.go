package file

import (
	"encoding/json"
	"fmt"
	"skullpt-lib/data_types"
)

type FileConfig struct {
	Action string `json:"action"`
	Mode   string `json:"mode"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Target string `json:"target"`
}

func Register() data_types.ActionCatalouge {
	return CommandLookup
}

func CommandLookup(cfg data_types.Specifics) (data_types.ActionCommand, error) {
	var filecfg FileConfig
	if err := json.Unmarshal(cfg, &filecfg); err != nil {
		return nil, err
	}

	switch filecfg.Action {
	case "create":
		return Create, nil
	default:
		return nil, fmt.Errorf("unsupported file action: %s", filecfg.Action)
	}
}
