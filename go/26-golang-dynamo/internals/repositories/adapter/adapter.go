package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Database struct {
	connection *dynamodb.DynamoDB
	logMode    bool
}

type Interface interface {
	Health() bool
	FindAll(condition expression.Expression, tableName string) (response *dynamodb.ScanOutput, err error)
	FindOne(condition map[string]interface{}, tableName string) (*dynamodb.GetItemOutput, error)
	CreateOrUpdate(entity interface{}, tableName string) (*dynamodb.PutItemOutput, error)
	Delete(condition map[string]interface{}, tableName string) (*dynamodb.DeleteItemOutput, error)
}

func NewAdapter(connection *dynamodb.DynamoDB) Interface {
	return &Database{
		connection: connection,
		logMode:    false,
	}
}

func (db *Database) Health() bool {
	_, err := db.connection.ListTables(&dynamodb.ListTablesInput{})
	return err == nil
}

func (db *Database) FindAll(condition expression.Expression, tableName string) (response *dynamodb.ScanOutput, err error) {
	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  condition.Names(),
		ExpressionAttributeValues: condition.Values(),
		FilterExpression:          condition.Filter(),
		ProjectionExpression:      condition.Projection(),
		TableName:                 aws.String(tableName),
	}
	return db.connection.Scan(input)
}

func (db *Database) FindOne(condition map[string]interface{}, tableName string) (*dynamodb.GetItemOutput, error) {
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       conditionParsed,
	}
	return db.connection.GetItem(input)
}

func (db *Database) CreateOrUpdate(entity interface{}, tableName string) (*dynamodb.PutItemOutput, error) {

	entityParsed, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(tableName),
	}
	return db.connection.PutItem(input)
}

func (db *Database) Delete(condition map[string]interface{}, tableName string) (*dynamodb.DeleteItemOutput, error) {
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.DeleteItemInput{
		Key:       conditionParsed,
		TableName: &tableName,
	}

	return db.connection.DeleteItem(input)
}
