package routes

import (
	"database/sql"
	"gin_gonic_products_rest_api/controllers"
	"gin_gonic_products_rest_api/db"
	"gin_gonic_products_rest_api/services"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	DB     *sql.DB
	Router *gin.Engine
}

func (r *Routes) CreateConnection() {
	db := db.Connectdb()
	r.DB = db
}

func (r *Routes) Routes() {
	router := gin.Default()
	productService := services.NewProductService(r.DB)
	controller := controllers.NewProductController(productService)

	router.POST("/product", controller.CreateProduct)
	router.GET("/product", controller.GetAllProducts)
	router.GET("/product/:id", controller.GetOneProduct)
	router.PUT("/product/:id", controller.UpdateProduct)
	router.DELETE("/product/:id", controller.DeleteProduct)

	r.Router = router
}

func (r *Routes) Run() {
	r.Router.Run(":8080")
}
