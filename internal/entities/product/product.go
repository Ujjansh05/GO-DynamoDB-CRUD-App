package product

import(
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"errors"
	"time"
)

type Product struct{
	entities.Base
	Name string `json:"name"`
}

func InterfaceToModel(data interface{})(instance *Product, err error){

}

func (p *Product) GetFilterId() map[string]interface{}{

}

func (p *Product) TableName() string{
		return "Products"
}

func (p *Product) Bytes() ([]byte, error){
	return json.Marshal(p)
}

func (p *Product) GetMap() map[string]interface{}{

}

func ParseDynamoAttributeToStruct()(){

}