package runner

type Runner interface {
	Restart()
	Start()
	Terminate()
}

type BaseRunner struct {
	Config map[string]interface{}
}
