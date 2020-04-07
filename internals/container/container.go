package container

import (
	"fmt"
	"go.uber.org/dig"
	"product-service/internals/config"
	"product-service/internals/controller"
	"product-service/internals/infrastructure/http"
)

type Container struct {
	container *dig.Container
}

func (c *Container) Configure() error {
	if err := c.container.Provide(controller.NewPingController); err != nil {
		return err
	}
	if err := c.container.Provide(controller.NewProductController); err != nil {
		return err
	}
	if err := c.container.Provide(http.NewHTTPServer); err != nil {
		return err
	}
	if err := c.container.Provide(config.NewConfiguration); err != nil {
		return err
	}
	return nil
}

func (c *Container) Start() error {
	fmt.Println("Start Container")
	if err := c.container.Invoke(func(s *http.Server) {
		s.Start()
	}); err != nil {
		return err
	}
	return nil
}


func NewContainer() (*Container, error) {
	d := dig.New()
	container := &Container{
		container: d,
	}
	if err := container.Configure(); err != nil {
		return nil, err
	}
	return container, nil
}