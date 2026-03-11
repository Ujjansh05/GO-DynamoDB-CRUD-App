package rules


import(
	"io"
	"github.com/aws/aws-sdk-go/service/dyanmodb"
)

type Interface interface{

	ConvertIoReaderToStruct(data io.Reader, model interface{})(body interface{}, err error)
		GetMock() interface{}
		Migrate (connection *dynamodb.DynamoDB)error
		Validate(model interface{}) error
}
