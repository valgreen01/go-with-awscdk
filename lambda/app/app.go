package app

import (
	"lambda-func/api"
	"lambda-func/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	// We actually initialize our DB store
	// gets passed down into the api handler
	db := database.NewDynamoDBClient()
	apiHandler := api.NewApiHandler(db)

	return App{
		ApiHandler: apiHandler,
	}
}
