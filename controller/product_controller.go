package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pierrialexandervidmar/simple-go-mod/model"
	"github.com/pierrialexandervidmar/simple-go-mod/usecase"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts();

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	
	var product model.Product
	
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUsecase.CreateProducts(product)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct.ID)
}

func (p *productController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("productId")

	if(id == "") {
		response := model.Response {
			Message: "O ID do produto não pode ser vazia!",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if(err != nil) {
		response := model.Response {
			Message: "ID do produto precisa ser um número!",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUsecase.GetProductById(productId);

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	
	if product == nil {
		response := model.Response {
			Message: "Produto não foi encontrado na base de dados!",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

