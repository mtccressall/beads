package setup

import "github.com/steveyegge/beads/internal/templates/agents"

var opencodeIntegration = agentsIntegration{
	name:         "OpenCode",
	setupCommand: "bd setup opencode",
	readHint:     "OpenCode reads AGENTS.md at the start of each session. Restart OpenCode if it is already running.",
	profile:      agents.ProfileFull,
}

var opencodeEnvProvider = defaultAgentsEnv

// InstallOpenCode installs the OpenCode agent integration.
func InstallOpenCode() error {
	return installOpenCode(opencodeEnvProvider())
}

// installOpenCode installs the OpenCode agent integration.
func installOpenCode(env agentsEnv) error {
	return installAgents(env, opencodeIntegration)
}

// CheckOpenCode checks if the OpenCode agent integration is installed.
func CheckOpenCode() error {
	return checkOpenCode(opencodeEnvProvider())
}

// checkOpenCode checks if the OpenCode agent integration is installed.
func checkOpenCode(env agentsEnv) error {
	return checkAgents(env, opencodeIntegration)
}

// RemoveOpenCode removes the OpenCode agent integration.
func RemoveOpenCode() error {
	return removeOpenCode(opencodeEnvProvider())
}

// removeOpenCode removes the OpenCode agent integration.
func removeOpenCode(env agentsEnv) error {
	return removeAgents(env, opencodeIntegration)
}
