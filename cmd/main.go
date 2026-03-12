package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/config"
	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/repository/adapter"
	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/repository/instance"
	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/routes"
	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/rules"
	RulesProduct "github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/rules/product"
	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/utils/logger"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	configs := config.GetConfig()
	connection := instance.GetConnection()
	repository := adapter.NewAdapter(connection)

	logger.INFO("waiting for the service to start....", nil)
	errors := migrate(connection)
	if len(errors) > 0 {
		for _, err := range errors {
			logger.PANIC("Errors on migrate:", err)
		}
	}
	logger.PANIC("Error checking tables:", checkTables(connection))

	port := fmt.Sprintf(":%v", configs.Port)
	router := routes.NewRouter().SetRouter(repository)
	logger.INFO("service is running on port", port)

	server := http.ListenAndServe(port, router)
	log.Fatal(server)
}

func migrate(connection *dynamodb.DynamoDB) []error {
	var errors []error
	callMigrateAndAppendError(&errors, connection, &RulesProduct.Rules{})
	return errors
}

func callMigrateAndAppendError(errors *[]error, connection *dynamodb.DynamoDB, rule rules.Interface) {
	err := rule.Migrate(connection)
	if err != nil {
		*errors = append(*errors, err)
	}
}

func checkTables(connection *dynamodb.DynamoDB) error {
	response, err := connection.ListTables(&dynamodb.ListTablesInput{})
	if response != nil {
		if len(response.TableNames) == 0 {
			logger.INFO("Tables not found:", nil)
		}

		for _, tableName := range response.TableNames {
			logger.INFO("Table found", *tableName)
		}
	}
	return err
}
