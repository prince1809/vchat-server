package plugin

import "github.com/prince1809/vchat-server/model"

// The API can be used to retrieve data or perform actions on behalf of the plugin. Most methods
// have direct counterparts in the REST API and very similar behavior.
//
// Plugins obtain access to the API by embedding MattermostPlugin and accessing the API member
// directly.
type API interface {
	// LoadPluginConfiguration loads the plugin's configuration. dest should be a pointer to a
	// struct that the configuration JSON can be unmarshalled to.
	LoadPluginConfiguration(dest interface{}) error

	// RegisterCommand registers a custom slash command. When the command is triggered, your plugin
	// can fulfill it via the ExecuteCommand hook.
	RegisterCommand(command *model.Command) error

	// UnregisterCommand unregisters a command previously registered via RegisterCommand.
	UnRegisterCommand(teamId, trigger string) error

	// GetSession returns the session object for the Session ID
	GetSession(sessionId string) (*model.Session, *model.AppError)
}
