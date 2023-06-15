package products

import "github.com/josiasbarretob/go-web/cmd/server/handler"

type IService interface {
	ListProducts() ([]handler.Product, error)
	CreateProduct(productReq handler.ProductRequest) (handler.Product, error)
	DescribeProduct(id int) (handler.Product, error)
	UpdateProduct(id int,  productReq handler.ProductRequest) (*handler.Product, error)
	DeleteProduct(id int) error
	PatchProduct(id int, productReq handler.ProductRequest)(*handler.Product, error)
	FindByNameProduct(nameProduct string)(*handler.Product, error)
}

type productService struct {
	repository IRepository
}

func NewProductService(repo IRepository) *productService {
	return &productService{
		repository: repo,
	}
}

func (p *productService) CreateProduct(productRequest handler.ProductRequest) (handler.Product, error) {
	product, err := p.repository.Save(productRequest)
	return product, err
}

func (p *productService) ListProducts() ([]handler.Product, error) {
	products, err := p.repository.GetAll()
	return products, err
}

func (p *productService) DescribeProduct(id int) (handler.Product, error) {
	product, err := p.repository.FindByID(id)
	return product, err
}
func (p *productService) UpdateProduct(id int,  productReq handler.ProductRequest)(*handler.Product, error){
	product, err := p.repository.Update(id, productReq)
	return product, err
}
func (p *productService) DeleteProduct(id int) error{
	err := p.repository.Del(id)
	return err
}
func (p *productService) PatchProduct(id int, productReq handler.ProductRequest)(*handler.Product, error){
	product, err := p.repository.Patch(id, productReq)
	return product, err
}
func (p *productService) FindByNameProduct(nameProduct string)(*handler.Product, error){
	product, err := p.repository.FindByName(nameProduct)
	return product, err
}

