package controllers

import "github.com/gin-gonic/gin"

type ProductControllerInterface interface {
	CreateProduct(*gin.Context)
	GetAllProducts(*gin.Context)
	GetOneProduct(*gin.Context)
	UpdateProduct(*gin.Context)
	DeleteProduct(*gin.Context)
}
