package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("request has empty parameters")
	}

	// Does a user with username already exist?
	userExists, err := api.dbStore.UserExists(event.Username)
	if err != nil {
		return fmt.Errorf("there was an error checking if the user exists %w", err)
	}

	if userExists {
		return fmt.Errorf("a user with that username already exists")
	}

	// We know the user doesn't exist, so let's insert it
	err = api.dbStore.InsertUser(event)
	if err != nil {
		return fmt.Errorf("there was an error inserting the user %w", err)
	}

	return nil
}
