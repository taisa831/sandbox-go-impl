package car

type common struct {
	handle string
}

type Car interface {
	Run() string
}

func (c *common) Run() string {
	return "run..."
}
