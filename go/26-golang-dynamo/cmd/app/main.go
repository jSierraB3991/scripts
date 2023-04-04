package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jsierrab3991/scripts/26-golang-dynamo/config"
	"github.com/jsierrab3991/scripts/26-golang-dynamo/internals/repositories/adapter"
	"github.com/jsierrab3991/scripts/26-golang-dynamo/internals/repositories/instance"
	"github.com/jsierrab3991/scripts/26-golang-dynamo/internals/routes"

	//"github.com/jsierrab3991/scripts/26-golang-dynamo/internals/routes"
	"github.com/jsierrab3991/scripts/26-golang-dynamo/internals/rules"
	RulesProduct "github.com/jsierrab3991/scripts/26-golang-dynamo/internals/rules/product"
	"github.com/jsierrab3991/scripts/26-golang-dynamo/utils/logger"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func callMigraetAndAppendError(errs *[]error, connection *dynamodb.DynamoDB, rules rules.Interface) {
	err := rules.Migrate(connection)
	if err != nil {
		*errs = append(*errs, err)
	}
}

func Migrate(connection *dynamodb.DynamoDB) []error {
	var errs []error
	callMigraetAndAppendError(&errs, connection, &RulesProduct.RulesProduct{})
	return errs
}

func checkTables(connection *dynamodb.DynamoDB) error {
	response, err := connection.ListTables(&dynamodb.ListTablesInput{})
	if response != nil {
		if len(response.TableNames) == 0 {
			logger.Info("Tables not found: ", nil)
		}

		for _, tableName := range response.TableNames {
			logger.Info("table found: ", *tableName)
		}
	}
	return err
}

func main() {
	configs := config.GetConfig()
	connection := instance.GetConnection()
	repository := adapter.NewAdapter(connection)

	logger.Info("waiting for the service to start...", nil)
	errs := Migrate(connection)
	if len(errs) > 0 {
		for _, err := range errs {
			logger.Panic("Error on migration...", err)
		}

	}
	logger.Panic("", checkTables(connection))
	port := fmt.Sprintf(":%v", configs.Port)
	router := routes.NewRouter().SetRouter((repository))
	logger.Info("Service is running on port", port)

	server := http.ListenAndServe(port, router)
	log.Fatal(server)
}
