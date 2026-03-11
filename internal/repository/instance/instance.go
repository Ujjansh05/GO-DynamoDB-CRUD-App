package instance 

import(
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetConnection() *dynamodb.DynamoDB{
	//Initialse a session that the SDK will use to load
	//creaditials from the shared crendiatials file ~/.aws/credientals
	//and the region from the shared configration file ~/.aws/config.
	
	sess:=session.Must(session.NewSesionwithOptions(session.Options{
		sessionConfigState:session.SharedConfigEnable,
	}))
	return dynamodb.New(sess)
}
