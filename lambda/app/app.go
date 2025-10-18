package app

import (
	"lambda-func/api"
	"lambda-func/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	// here we actually initialize of db store
	db := database.NewDynamoDBClient()
	apiHandler := api.NewApiHandler(db)
	// and then it is passed into api handler

	return App{
		ApiHandler: apiHandler,
	}
}
