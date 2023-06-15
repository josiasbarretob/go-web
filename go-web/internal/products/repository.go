package products

import (
	"fmt"

	"github.com/josiasbarretob/go-web/cmd/server/handler"
)

// Definicao das Interfaces Repository
type IRepository interface {
	GetAll() ([]handler.Product, error)
	Save(productReq handler.ProductRequest) (handler.Product, error)
	FindByID(id int) (handler.Product, error)
	Update(id int, productReq handler.ProductRequest) (*handler.Product, error)
	Del(id int) error
	Patch(id int, productReq handler.ProductRequest) (*handler.Product, error)
	FindByName(nameProduct string)(*handler.Product, error)
}

type repositoryH2 struct {
	product map[int]handler.Product
}

func NewRepository() *repositoryH2 {
	var product = make(map[int]handler.Product)
	return &repositoryH2{
		product: product,
	}
}

func (r *repositoryH2) Save(productRequest handler.ProductRequest) (handler.Product, error) {
	id := len(r.product)
	newProduct := handler.Product{
		ID:          id,
		Nome:        productRequest.Nome,
		Cor:         productRequest.Cor,
		Preco:       productRequest.Preco,
		Estoque:     productRequest.Estoque,
		Codigo:      productRequest.Codigo,
		Publicacao:  productRequest.Publicacao,
		DataCriacao: productRequest.DataCriacao,
	}
	r.product[id] = newProduct
	return newProduct, nil
}

func (r *repositoryH2) GetAll() ([]handler.Product, error) {
	var listProduct = make([]handler.Product, 0)
	for _, product := range r.product {
		listProduct = append(listProduct, product)
	}
	return listProduct, nil
}

func (r *repositoryH2) FindByID(id int) (handler.Product, error) {
	product, ok := r.product[id]
	if ok {
		return product, nil
	}
	return product, fmt.Errorf("ID não encontrado")
}

func (r *repositoryH2) Update(id int, productReq handler.ProductRequest) (*handler.Product, error) {
	if _, ok := r.product[id]; ok {
		updatedProduct := &handler.Product{
			ID:          id,
			Nome:        productReq.Nome,
			Cor:         productReq.Cor,
			Preco:       productReq.Preco,
			Estoque:     productReq.Estoque,
			Codigo:      productReq.Codigo,
			Publicacao:  productReq.Publicacao,
			DataCriacao: productReq.DataCriacao,
		}
		r.product[id] = *updatedProduct
		return updatedProduct, nil
	}
	return nil, fmt.Errorf("ID não encontrado")
}

func (r *repositoryH2) Del(id int) error {
	if _, ok := r.product[id]; ok {
		delete(r.product, id)
		return nil
	}
	return fmt.Errorf("ID não encontrado")
}

func (r *repositoryH2) Patch(id int, productReq handler.ProductRequest) (*handler.Product, error) {
	//var patchProduct handler.ProductRequest
	if prod, ok := r.product[id]; ok {
		patchProduct := &handler.Product{
			ID:          id,
			Nome:        productReq.Nome,
			Cor:         prod.Cor,
			Preco:       productReq.Preco,
			Estoque:     prod.Estoque,
			Codigo:      prod.Codigo,
			Publicacao:  prod.Publicacao,
			DataCriacao: prod.DataCriacao,
		}
		r.product[id] = *patchProduct
		return patchProduct, nil
	}
	return nil, fmt.Errorf("ID não encontrado")
}
func (r *repositoryH2)FindByName(nameProduct string)(*handler.Product, error){
	for _, product := range r.product{
		if product.Nome == nameProduct{
			return &product, nil
		}
	}
	return nil, fmt.Errorf("produto não encontrado pelo nome")
}