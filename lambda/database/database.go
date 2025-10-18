package database

type DynamoDBClient struct {
	databaseStore string
}

func DynamoDBClientFunction() DynamoDBClient {
	return DynamoDBClient{
		databaseStore: "database store",
	}
}
