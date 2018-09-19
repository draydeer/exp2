package process

import "massivep/pkg/runner"

type ProcessRunner struct {
	runner.BaseRunner
}

func NewProcessRunner() *ProcessRunner {
	return &ProcessRunner{}
}
