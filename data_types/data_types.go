package data_types

type Gallery struct {
	Name      string            `json: "name"`
	Endpoints []string          `json: "endpoints"`
	Actions   map[string]Action `json: "actions"`
}

// Action holds the action config
type Action struct {
	ActionType string    `json: "action_type"`
	Specifics  Specifics `json: "specifics"`
}

// Specifics is the config for the action
type Specifics []byte

// ActionCommand is the function signature all actions needs to do
type ActionCommand func(cfg Specifics) error

// AvailableCommands will hold all registerd commands for an action
// e.g. create = fn for creating
type ActionCatalouge func(cfg Specifics) (ActionCommand, error)

// AvailableActions holds all the currently registered actions
type AvailableActions map[string]ActionCatalouge

// QueuedCommands is a list of commands ready to be applied
type QueuedCommands struct {
	Endpoint string
	Commands map[string]ActionCommand
}
