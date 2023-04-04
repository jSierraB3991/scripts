package product

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jsierrab3991/scripts/26-golang-dynamo/internals/entities"
)

type ProductBody struct {
	Name string `json:"name"`
}

type Product struct {
	entities.Base
	Name string `json:"name"`
}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) GetFilterId() map[string]interface{} {
	return nil
}

func InterfaceToModel(product interface{}) (instance Product, err error) {
	bytes, err := json.Marshal(product)
	if err != nil {
		return instance, nil
	}
	return instance, json.Unmarshal(bytes, &instance)
}

func (p *Product) Bytes() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Product) GetMap() map[string]interface{} {
	return nil
}

func ParseDynamoAttributteToStruct(item map[string]*dynamodb.AttributeValue) (entity Product, err error) {
	return Product{}, nil
}
