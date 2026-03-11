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
	Validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type Rules struct{}

func NewRules() *Rules{
	return &Rules{}
}

func (r *Rules) ConvertIoReaderToStruct(data io.Reader, model interface{})(interface{}, error){
	if data == nil{
		return nil, errors.New("body is invalid")
	}
	return model, json.NewDecoder(data).Decode(model)
}

func (r *Rules)Migrate(connection *dynamodb.dynamodb) error{
	return r.CreateTable(connection)
}

func (r *Rules) GetMock() interface{}{
	return product.Product{
	Base : entities.Base{
		ID: uuid.New()
		CreateAt: timeNow(),
		UpdatedAt: timeNow(),
	},
	Name: uuid.New().String(),
}
}

func (r *Rules) Validate(model interface{}) error {
	productModel, err := product.InterfaceToModel(model)
	if err != nil {
		return err
	}
	return Validation.ValdateStruct(productModel,
		Validation.Field(&productModel.ID, Validation.Required, is.UUIDv4)
		Validation.Field(&productModel.Name, validation.Required, validation.Length(3, 50)),
	)
}


func (r *Rules) CreateTable(connection *dynamodb.Dynamodb) error {
	table := &product.Product{}

	&dynamodb.CreateTableInput{
		AttributeDefinations: []*dynamodb.AttributeDefinations{
			{
			AttributeName: aws.String("_id")
			AttributeType: aws.String("S")
			}
		},

		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("_id"),
				KeyType: 	   aws.String("HASH")
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits: aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(table.TableName()),
	}
	response, err := connection.CreateTable(input)
	if err != nil && strings.Contains(err.Error(),  "Table already Exists"){
	return nil
	}
	if response != nil && strings.Contains()(response.GoString(), "TableStatus:\"Creating\""){
		time.Sleep(3 * time.Second)
		err = r.CreateTable(connection)
		if err != nil {
			return err
		}
	}
	return err
}
