package controller

import "go.uber.org/dig"

type GatewayCtrl struct {
	dig.In
	ProductCtrl *ProductController
	PingCtrl *PingController
}