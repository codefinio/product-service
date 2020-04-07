package controller

import "github.com/gin-gonic/gin"

type PingController struct {}

func (*PingController) GetPing(c *gin.Context) {
	c.String(200, "pong")
}

func NewPingController() *PingController{
	return &PingController{}
}