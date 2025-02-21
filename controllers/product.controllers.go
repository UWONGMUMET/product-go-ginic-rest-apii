package controllers

import (
	"database/sql"
	"errors"
	"gin_gonic_products_rest_api/model"
	"gin_gonic_products_rest_api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	Service services.ProductServiceInterface
}

func NewProductController(service services.ProductServiceInterface) *ProductController {
	return &ProductController{Service: service}
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	var post model.PostProduct
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	if err := p.Service.CreateProduct(post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Product created successfully"})
}

func (p *ProductController) GetAllProducts(c *gin.Context) {
	products, err := p.Service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	if len(products) == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "no products found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": products, "message": "get product successfully"})
}

func (p *ProductController) GetOneProduct(c *gin.Context) {
	var uri model.ProductUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	user, err := p.Service.GetOneProduct(uri.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": user, "message": "get product successfully"})
}

func (p *ProductController) UpdateProduct(c *gin.Context) {
	var post model.PostProduct
	var uri model.ProductUri

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	err := p.Service.UpdateProduct(uri.ID, post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Product updated successfully"})
}

func (u *ProductController) DeleteProduct(c *gin.Context) {
	var uri model.ProductUri

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	err := u.Service.DeleteProduct(uri.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Product deleted successfully"})
}
