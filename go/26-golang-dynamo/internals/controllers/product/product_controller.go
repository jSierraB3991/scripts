package product

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
	"github.com/jsierrab3991/scripts/26-golang-dynamo/internals/entities/product"
	"github.com/jsierrab3991/scripts/26-golang-dynamo/internals/repositories/adapter"
)

type Interface interface {
	ListAll() ([]product.Product, error)
	ListOne(id uuid.UUID) (product.Product, error)
	Create(body product.Product) (uuid.UUID, error)
	Update(id uuid.UUID, body product.Product) error
	Remove(id uuid.UUID) error
}

type ProductController struct {
	repository adapter.Interface
}

func NewController(repository adapter.Interface) Interface {
	return &ProductController{
		repository: repository,
	}
}

func (pc *ProductController) Create(body product.Product) (uuid.UUID, error) {
	body.GenerateId()
	body.SetCreatedAt()
	_, err := pc.repository.CreateOrUpdate(body.GetMap(), body.TableName())
	return body.Id, err
}

func (pc *ProductController) Update(id uuid.UUID, body product.Product) error {
	found, err := pc.ListOne(id)
	if err != nil {
		return err
	}
	found.Id = id
	found.Name = body.Name
	found.SetUpdateAt()
	_, err = pc.repository.CreateOrUpdate(found.GetFilterId(), body.TableName())
	return err
}
func (pc *ProductController) Remove(id uuid.UUID) error {
	entity, err := pc.ListOne(id)
	if err != nil {
		return err
	}
	_, err = pc.repository.Delete(entity.GetFilterId(), entity.TableName())
	return err
}

func (pc *ProductController) ListOne(id uuid.UUID) (entity product.Product, err error) {
	entity.Id = id
	response, err := pc.repository.FindOne(entity.GetFilterId(), entity.TableName())
	if err != nil {
		return entity, err
	}
	return product.ParseDynamoAttributteToStruct(response.Item)
}

func (pc *ProductController) ListAll() ([]product.Product, error) {

	list := []product.Product{}
	var entity product.Product
	filter := expression.Name("name").NotEqual(expression.Value(""))
	condition, err := expression.NewBuilder().WithFilter(filter).Build()

	if err != nil {
		return list, err
	}

	response, err := pc.repository.FindAll(condition, entity.TableName())
	if err != nil {
		return list, err
	}
	if response != nil {
		for _, value := range response.Items {
			entity, err = product.ParseDynamoAttributteToStruct(value)
			if err != nil {
				return list, err
			}
			list = append(list, entity)
		}
	}
	return list, nil
}
