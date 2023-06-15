package main

import (
	"log"
	// "net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	//"github.com/josiasbarretob/go-web/cmd/server/handler"
	"github.com/josiasbarretob/go-web/internal/products"
)

var repo = products.NewRepository()
var implementsService = products.NewProductService(repo)
var controller = products.NewProductController(implementsService)

func main() {
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error ao carregar as variáveis de ambiente:", err)
	}
	router := gin.Default()
	products := router.Group("/products")
	router.GET("/", controller.DescribeNameProduct)
	products.GET("/", controller.ListProducts)
	products.GET("/:id", controller.DescribeProduct)
	products.POST("/", controller.CreateProduct)
	products.PUT("/:id", controller.UpdateProduct)
	products.DELETE("/:id", controller.DeleteProduct)
	products.PATCH("/:id", controller.PatchProduct)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}