package runner

type Runner interface {
	Stream(topic ...string)
}
