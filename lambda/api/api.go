package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.UserStore
}

func NewApiHandler(dbStore database.UserStore) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("payload had empty parameters")
	}

	// does user already exists
	userExists, err := api.dbStore.DoesUserExists(event.Username)

	if err != nil {
		return fmt.Errorf("error occured while checking if the user exists %w", err)
	}

	if userExists {
		return fmt.Errorf("user with the given username already exists")
	}

	// if user doesnt exist already then push it to DB
	err = api.dbStore.InsertUser(event)

	if err != nil {
		return fmt.Errorf("error occured while inserting a user %w", err)
	}

	return nil
}
