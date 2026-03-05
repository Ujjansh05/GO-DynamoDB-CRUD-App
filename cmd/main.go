package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/respository/adapter"
	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/respository/instance"
	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/routes"
	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/respository/rules"
	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/rules/product"
	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/utlis/logger"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)


func main(){
	configs := Config.GetConfig()
	connection := instance.GetConnection()
	respository := adapter.NewAdapter(connection)


	logger.INFO("waiting for the service to start....", nil)
	errors := Migrate(connection)
	if len(errors) > 0 {
		for _, err := range errors {
			logger.PANIC("Errors on migrations.......", err )

		} 
	}
	logger.PANIC("", checkTables(connection))
	port := fmt.Sprintf(":%v", configs.Port)
	router := routes.NewRouter().SetRouter(respository)
	logger.INFO("service is running on port", port)
	server := http.ListenAndServe(port, router)
	log.Fatal(server)
}


func Migrate(connection *dynamodb.DynamoDB) []error{
	var errors []error
	callMigrateAndAppendError (&errors, connection, &RulesProduct.Rules{})

}


func checkTables(){
	
}