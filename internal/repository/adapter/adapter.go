package adapter 


import(
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Data struct{
	connection *dynamodb.DynamoDB
	logMode bool
}


type Interface interface{

}

func Newadapter () Interface{

}


func (db * Database) Health() bool{
	_,err := db.connection.ListTables(&dynamo.ListTablesInput{})
	return err == nil 
}

func (db * Database) FindAll{

}

func (db * Database) FindOne(condition map[string]Interface{}, tableName)(response *dynamodb.GetItemOutput, err error){
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

