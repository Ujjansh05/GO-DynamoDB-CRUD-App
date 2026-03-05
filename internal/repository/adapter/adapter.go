package adapter 


import(
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-g0/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sadk-go/dynamodb/expression"
)

type Data struct{
	connection *dynamodb.DynamoDB
	logMode bool
}

type Interface interface{
  Health () bool
  FindAll(condition expression.Expression, tableName string)(response *dynamodb.ScanOutput, err error)
  FindOne(condition map[string]interface{}, tableName string)(response *dynamodb.GetItemOutput, err error)
  CreateOrUpdate(entity interface{}, tableName string)(response *dynamodb.PutItemOutput, err error)
  Delete(condition map[string]interface{}, tableName string)(response *dynamodb.DeleteItemOutput, err error)

}

func Newadapter (con *dynamodb.ynamoDB) Interface{
	return &Database{
	connection: con,
	logMode : false,
	}
}


func (db * Database) Health() bool{
	_,err := db.connection.ListTables(&dynamo.ListTablesInput{})
	return err == nil 
}

func (db * Database) FindAll(condition expression.Expression, tableName string){
	input := &dynamodb.ScanInput{
		ExpressionAttributesNames: condition.Names(),
			ExpressionAttributesValues: condition.Values(),
			FilterExpressions: conditions.Filter(),
			ProjectionExpression: condition.Projection(),
	TableName : aws.String(tableName),
	}
}

func (db * Database) FindOne(condition map[string]interface{}, tableName)(response *dynamodb.GetItemOutput, err error){
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)

	if err !== nil {
		return nil , err
	}

	input := &dynamodb.GetItemInput{
		TableName : aws.String(tablename)
		key : conditionParsed,
	}
	return db.connectionGetItem(input)
}


func (db * Database) CreateOrUpdate (entity interface{}, tableName string)(response *dynamodb.PutItemOutput, err error){

	entityParsed, err := dynamodbattribute.MarshalMap(entity)

	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item : entityParsed, 
		TableName : aws.String(tableName),
	}
	return db.connection.PutItem(input)
}

func (db * Database) Delete(condition map[string] interface{} tableName string)(response *dynamodb.DeleteItemOutput, err error){
	

	input := &dynamodb.DeleteItemInput{
		if err != nil {
			return nil, err
		}

		Key : conditionParsed,
		TableName : aws.String(tableName),
	}
	return db.connection.DeleteItem(input)
}

