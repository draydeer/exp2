package process

import "massivep/pkg/runner"

type ProcessRunner struct {
	runner.BaseRunner
}

func (runner *ProcessRunner) SetConfig(config map[string]interface{})  {
	runner.Config = config
}
