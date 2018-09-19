package docker_container

import "massivep/pkg/runner"

type DockerContainerRunner struct {
	runner.BaseRunner
}

func NewDockerContainerRunner() *DockerContainerRunner {
	return &DockerContainerRunner{}
}
