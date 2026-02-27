package instance 

import(
	"github.com/aws/aws-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetConnection() *dynamodb.DynamoDB{


	sess:=session.Must(session.NewSesionwithOptions(session.Options{
		sessionConfigState:session.SharedConfigEnable,
	}))
	return dynamodb.New(sess)
}