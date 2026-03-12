package product

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/entities"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
)

type Product struct {
	entities.Base
	Name string `json:"name"`
}

func InterfaceToModel(data interface{}) (instance *Product, err error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	instance = &Product{}
	if err := json.Unmarshal(bytes, instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (p *Product) GetFilterId() map[string]interface{} {
	return map[string]interface{}{"_id": p.ID.String()}
}

func (p *Product) TableName() string {
	return "Products"
}

func (p *Product) Bytes() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Product) GetMap() map[string]interface{} {
	return map[string]interface{}{
		"_id":       p.ID.String(),
		"name":      p.Name,
		"CreatedAt": p.CreatedAt.Format(entities.GetTimeFormat()),
		"UpdatedAt": p.UpdatedAt.Format(entities.GetTimeFormat()),
	}
}

func ParseDynamoAttributeToStruct(response map[string]*dynamodb.AttributeValue) (p Product, err error) {
	if response == nil || len(response) == 0 {
		return p, errors.New("Item not found")
	}

	for key, value := range response {
		switch key {
		case "_id":
			if value == nil || value.S == nil {
				return p, errors.New("Item not found")
			}

			p.ID, err = uuid.Parse(*value.S)
			if p.ID == uuid.Nil {
				err = errors.New("Item not found")
			}
		case "name":
			if value != nil && value.S != nil {
				p.Name = *value.S
			}
		case "CreatedAt", "createdAt":
			if value != nil && value.S != nil {
				p.CreatedAt, err = time.Parse(entities.GetTimeFormat(), *value.S)
			}
		case "UpdatedAt", "updatedAt":
			if value != nil && value.S != nil {
				p.UpdatedAt, err = time.Parse(entities.GetTimeFormat(), *value.S)
			}
		}

		if err != nil {
			return p, err
		}
	}

	return p, nil
}
