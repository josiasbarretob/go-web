package products

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/josiasbarretob/go-web/cmd/server/handler"
	"github.com/josiasbarretob/go-web/pkg/web"
)

type productController struct {
	service IService
}

func NewProductController(service IService) *productController {
	return &productController{
		service: service,
	}
}

func (s *productController) CreateProduct(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(400, nil, "Token inválido!"))
		return
	}

	var productRequest handler.ProductRequest
	if err := ctx.ShouldBindJSON(&productRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		//errRest := validation.ValidateUserError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	product, err := s.service.CreateProduct(productRequest)
	if err != nil {
		log.Printf("Erro ao criar produto!")
		ctx.JSON(http.StatusInternalServerError, nil)
	}
	ctx.JSON(http.StatusCreated, product)
}

func (p *productController) DescribeProduct(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(400, nil, "Token inválido!"))
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Erro ao converter o ID para inteiro.")
	}
	product, err := p.service.DescribeProduct(id)
	if err != nil {
		log.Printf("Erro ao buscar produto por ID!")
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (p *productController) ListProducts(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(400, nil, "Token inválido!"))
		return
	}

	products, err := p.service.ListProducts()
	if err != nil {
		log.Printf("Erro ao listar produtos!")
		ctx.JSON(http.StatusInternalServerError, nil)
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(400, nil, "Token inválido!"))
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var request handler.ProductRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if request.Nome == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo nome do produto é obrigatório!"})
		return
	}
	if request.Cor == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo COR do produto é obrigatório!"})
		return
	}
	if request.Preco == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo PRECO do produto é obrigatório!"})
		return
	}
	if request.Estoque == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo ESTOQUE do produto é obrigatório!"})
		return
	}
	if request.Codigo == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo CODIGO do produto é obrigatório!"})
		return
	}
	// if request.Publicacao != false || request.Publicacao != true {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo PUBLICACAO do produto é obrigatório!"})
	// 	return
	// }
	if request.DataCriacao == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo DATA DE CRIACAO do produto é obrigatório!"})
		return
	}
	updateProduct, err := p.service.UpdateProduct(id, request)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, updateProduct)
}

func (p *productController) DeleteProduct(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(400, nil, "Token inválido!"))
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	erro := p.service.DeleteProduct(id)
	if erro != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
func (p *productController) PatchProduct(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(400, nil, "Token inválido!"))
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	productReq := handler.ProductRequest{}
	ctx.ShouldBindJSON(&productReq)
	updatedProduct, err := p.service.PatchProduct(id, productReq)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"erro": err})
	}
	ctx.JSON(http.StatusOK, updatedProduct)

}

func (p *productController) DescribeNameProduct(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(400, nil, "Token inválido!"))
		return
	}

	nameProduct := ctx.Query("nome")
	if nameProduct == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Parametros inválidos"})
		return
	}

	product, err := p.service.FindByNameProduct(nameProduct)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, product)
}
