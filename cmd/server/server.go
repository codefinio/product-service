package main

import "product-service/internals/container"

func main() {
	c, err := container.NewContainer()
	if err != nil {
		panic(err)
	}
	if err := c.Start(); err != nil {
		panic(err)
	}
}
