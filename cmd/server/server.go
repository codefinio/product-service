package main

import (
	httpServer "product-service/internals/infrastructure/http"
)

func main() {
	sv := httpServer.NewHTTPServer()
	sv.Start()
}
