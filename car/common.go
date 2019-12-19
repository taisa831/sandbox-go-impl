package car

type common struct {
	handle string
}

type car interface {
	run() string
}

func (c *common) run() string {
	return "run..."
}
