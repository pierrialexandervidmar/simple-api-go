package usecase

import (
	"github.com/pierrialexandervidmar/simple-go-mod/model"
	"github.com/pierrialexandervidmar/simple-go-mod/repository"
)

// productUsecase define a estrutura da camada de caso de uso (lógica de negócio)
// Aqui normalmente você injetaria um repositório (ex: productRepository) para acessar o banco de dados
type ProductUsecase struct {
	// Repository
	repository repository.ProductRepository
}

// NewProductUseCase é uma função construtora que retorna uma instância de productUsecase
// Ideal para injetar dependências no futuro (como um repositório)
func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

// GetProducts é um método da camada de caso de uso que retorna uma lista de produtos
// No futuro, ele deve chamar um repositório real para buscar dados do banco
func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProducts(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProducts(product)
	
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) GetProductById(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id_product)

	if err != nil {
		return nil, err
	}

	return product, nil
}
