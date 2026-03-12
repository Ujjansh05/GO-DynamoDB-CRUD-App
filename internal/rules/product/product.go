package product

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
	"time"

	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/entities"
	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/entities/product"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type Rules struct{}

func NewRules() *Rules {
	return &Rules{}
}

func (r *Rules) ConvertIoReaderToStruct(data io.Reader, model interface{}) (interface{}, error) {
	if data == nil {
		return nil, errors.New("body is invalid")
	}
	return model, json.NewDecoder(data).Decode(model)
}

func (r *Rules) Migrate(connection *dynamodb.DynamoDB) error {
	return r.CreateTable(connection)
}

func (r *Rules) GetMock() interface{} {
	now := time.Now()
	return product.Product{
		Base: entities.Base{
			ID:        uuid.New(),
			CreatedAt: now,
			UpdatedAt: now,
		},
		Name: uuid.New().String(),
	}
}

func (r *Rules) Validate(model interface{}) error {
	productModel, err := product.InterfaceToModel(model)
	if err != nil {
		return err
	}

	return validation.ValidateStruct(productModel,
		validation.Field(&productModel.ID, validation.Required, is.UUIDv4),
		validation.Field(&productModel.Name, validation.Required, validation.Length(3, 50)),
	)
}

func (r *Rules) CreateTable(connection *dynamodb.DynamoDB) error {
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

	response, err := connection.CreateTable(input)
	if err != nil {
		message := strings.ToLower(err.Error())
		if strings.Contains(message, "table already exists") || strings.Contains(message, "resourceinuseexception") {
			return nil
		}
		return err
	}

	if response != nil &&
		response.TableDescription != nil &&
		response.TableDescription.TableStatus != nil &&
		*response.TableDescription.TableStatus == dynamodb.TableStatusCreating {
		time.Sleep(3 * time.Second)
	}

	return nil
}
