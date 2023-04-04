package rules

import (
	"io"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Interface interface {
	ConvertIoReaderToStruct(data io.Reader, model interface{}) (body interface{}, err error)
	GetMock() interface{}
	Migrate(conn *dynamodb.DynamoDB) error
	Validate(model interface{}) error
}
