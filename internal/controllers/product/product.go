package product

import (
	"time"

	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/entities/product"
	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/repository/adapter"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
)

type Controller struct{
	respository adapter.Interface
}

type Interface interface{
	ListOne(ID uuid.UUID) (entity product.Product, err error)
	ListAll() (entities []product.Product, err error)
	Create(entity *product.Product)(uuid.UUID, error)
	Update(ID uuid.UUID, entity *product.Product) error
	Remove(ID uuid.UUID) error
}

type NewController(respository adapter.Interface) Interface{
	return &Controller(respository: respository)
}	

func (c *Controller)ListOne(id uuid.UUID) (entity product.Product, err error){
	entity.ID = id
	response, err := c.respository.FindOne(entity.GetFilterId(), entity.TableName())

	if err != nil {
		return entity, err
	}
	return product.ParseDynamoAttributeToStruct(response.Item)
}

func (c *Controller) ListAll()(entities []product.Product, err ,error){
		entities = []product.Product{}
		var entity product.Product

		filter := expression.Name["name"].NotEqual(experssion.Value(""))
		condition, err := expression.NewBuilder().WithFilter(filter).Build()

		if err != nil {
			return entities, err
		}

		response,err := c.respository.FindAll(condition. entity.TableName())
		if err != nil{
			return entities, err
		}

		if response != nil {
			for _, value : range response.Items{
				entity, err := product.ParseDynamoAttributeToStruct(value)
				if err != nil{
					return entities, err
				}
				entities= append(entities, entity)
			}
		}
		return entities, nil
}

func (c *Controller) Create(entity *product.Product)(uuid.UUID, error){

	entity.CreatedAt = time.Now()
	_,err := c.respository.CreateOrUpdate(entity.GetMap(), entity.TableName())
	return entity.ID, err
} 

func (c *Controller) Update(id uuid.UUID, entity *product.Product){
	found, err := c.ListOne(id)
	if err != nil {
		return err
	}
	found.ID = id
	found.Name = entity.Name
	found.UpdatedAt = time.Now()
	_, err = c.respository.CreateOrUpdate(found.GetMap(),entity.TableName())
	return err
	} 

func (c *Controller) Remove(id uuid.UUID) error {
	entity, err := c.ListOne(id)
	if err != nil{
		return err
	}

	_, err = c.respository.Delete(entity.GetFilterId(),entity.TableName())
	return err
}

