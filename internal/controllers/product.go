package product

import "github.com/gofrs/uuid"

type Controller struct{
	respository adapter.Interface
}

type Interface interface{
	ListOne(id uuid.UUID) (entity product.Product, err error)
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

}

func (c *Controller) Create(entity *product.Product)(uuid.UUID, error){

} 

func (c *Controller) Update(id uuid.UUID, entity *product.Product){

} 

func (c *Controller) Remove(id uuid.UUID) error {

}

