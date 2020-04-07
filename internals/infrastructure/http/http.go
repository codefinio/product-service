package http

import (
	"github.com/gin-gonic/gin"
	ctrl "product-service/internals/controller"
)
// infrastructure -> Controller -> Service -> Entity
// Service -> Repository -> DB
// Server ...
type Server struct {
	route *gin.Engine
	pingCtrl *ctrl.PingController

	productCtrl *ctrl.ProductController
}

func (h *Server) Configure() {
	r := h.route
	r.GET("/ping", h.pingCtrl.GetPing)

	g := r.Group("/product")
	{
		g.GET("/getAll",h.productCtrl.GetProducts)
		g.POST("/add",h.productCtrl.AddProduct)
		g.PUT("/update/:id", h.productCtrl.UpdateProduct)
	}
}
// Start ...
func (h *Server) Start() {
	h.Configure()
	if err := h.route.Run(":3000"); err != nil {
		panic(err)
	}
}

func NewHTTPServer() *Server {
	ctrl.NewPingController()
	return &Server{
		route: gin.Default(),
		pingCtrl: ctrl.NewPingController(),
		productCtrl: ctrl.NewProductController(),
	}
}