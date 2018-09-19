package docker_container

import "massivep/pkg/runner"

type DockerContainerRunner struct {
	runner.BaseRunner
}

func (runner *DockerContainerRunner) SetConfig(config map[string]interface{})  {
	runner.Config = config
}
