package controller

import "C"
import (
	"errors"
	"github.com/gin-gonic/gin"
	"product-service/internals/entity"
	"product-service/internals/service"
	"strconv"
)

type ProductController struct {
	Products []*entity.Product
	ProductService *service.ProductService
}

func (p *ProductController) GetProducts(c *gin.Context) {
	products, err := p.ProductService.ListProducts()
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, products)
}

func (p *ProductController) AddProduct(c *gin.Context) {
	item := &entity.Product{}
	if err := c.BindJSON(item); err != nil {
		c.AbortWithError(500, err)
		return
	}
	if err := p.ProductService.ImportProduct(item); err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, item)
}


func (p *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	reqProduct := &entity.Product{}
	if err := c.BindJSON(reqProduct); err != nil {
		c.AbortWithError(400, err)
	}
	var sp *entity.Product
	for _,pd := range p.Products {
		if id == strconv.Itoa(pd.Id) {
			sp = pd
		}
		break
	}
	if sp == nil {
		c.AbortWithError(400, errors.New("ID Not Found"))
		return
	}
	sp.Name = reqProduct.Name
	sp.Price = reqProduct.Price
	c.JSON(200, sp)
}

// NewProductController ...
func NewProductController() *ProductController{
	return &ProductController{
		Products: []*entity.Product{},
		ProductService: service.NewProductService(),
	}
}