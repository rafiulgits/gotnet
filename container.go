package gotnet

import "go.uber.org/dig"

type Container struct {
	*dig.Container
}

func newContainer() *Container {
	return &Container{
		dig.New(),
	}
}
