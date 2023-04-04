package product

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/jsierrab3991/scripts/26-golang-dynamo/internals/entities"
	"github.com/jsierrab3991/scripts/26-golang-dynamo/internals/entities/product"
)

type RulesProduct struct{}

func NewRules() *RulesProduct {
	return &RulesProduct{}
}

func (rp *RulesProduct) Migrate(conn *dynamodb.DynamoDB) error {
	return rp.createTable(conn)
}

func (rp *RulesProduct) ConvertIoReaderToStruct(data io.Reader, model interface{}) (interface{}, error) {
	if data == nil {
		return nil, errors.New("Body is Invalid")
	}
	return model, json.NewDecoder(data).Decode(model)
}

func (rp *RulesProduct) Validate(model interface{}) error {
	product, err := product.InterfaceToModel(model)
	if err != nil {
		return err
	}
	return validation.ValidateStruct(model,
		validation.Field(&product.Id, validation.Required, is.UUIDv4),
		validation.Field(&product.Name, validation.Required, validation.Length(3, 50)),
	)
}

func (rp *RulesProduct) GetMock() interface{} {
	return product.Product{
		Base: entities.Base{
			Id: uuid.New(),
		},
		Name: uuid.NewString(),
	}
}

func (rp *RulesProduct) createTable(conn *dynamodb.DynamoDB) error {
	table := &product.Product{}
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("_id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("_id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(table.TableName()),
	}
	response, err := conn.CreateTable(input)
	if err != nil && strings.Contains(err.Error(), "Table already exists") {
		return nil
	}
	if response != nil && strings.Contains(response.GoString(), "TableStatus:\"CREATING\"") {
		time.Sleep(3 * time.Second)
		err := rp.createTable(conn)
		if err != nil {
			return err
		}
	}
	return err
}
