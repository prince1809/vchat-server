package plugin

const (
	HEALTH_CHECK_INTERVAL         = 30 // seconds. How often the health check should run
	HEALTH_CHECK_DISABLE_DURATION = 60 // minutes. How long we wait for num fails to incur before disabling the plugin
	HEALTH_CHECK_PING_FAIL_LIMIT  = 3  // How many times we call RPC ping in a row before it is considered a failure
	HEALTH_CHECK_RESTART_LIMIT    = 3  // How many times we restart a plugin before we disable it
)

type PluginHealthCheckJob struct {
	cancel    chan struct{}
	cancelled chan struct{}
	env       *Environment
}
