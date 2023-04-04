package product

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
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

func ParseDynamoAttributteToStruct(response map[string]*dynamodb.AttributeValue) (p Product, err error) {
	if response == nil || (response != nil && len(response) == 0) {
		return p, errors.New("Item not found")
	}
	for key, value := range response {
		if key == "_id" {
			p.Id, err = uuid.Parse(*value.S)
			if p.Id == uuid.Nil {
				err = errors.New("Item not found")
			}
		}
		if key == "name" {
			p.Name = *value.S
		}
		if key == "createdAt" {
			p.CreateAt, err = time.Parse(entities.GetTimeFormat(), *value.S)
		}
		if key == "updatedAt" {
			p.UpdateAt, err = time.Parse(entities.GetTimeFormat(), *value.S)
		}
		if err != nil {
			return p, err
		}
	}

	return p, nil
}
