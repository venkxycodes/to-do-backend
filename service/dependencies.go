package service

import (
	"to-do/appcontext"
	_ "to-do/appcontext"
	"to-do/repo"
)

type ServerDependencies struct {
	ToDoService ToDoService
	UserService UserService
}

func InstantiateServerDependencies() *ServerDependencies {
	dbClient := appcontext.GetDBClient()
	toDoRepo := repo.NewToDoRepo(dbClient)
	userRepo := repo.NewUserRepository(dbClient)
	return &ServerDependencies{
		ToDoService: NewToDoService(toDoRepo),
		UserService: NewUserService(userRepo),
	}
}
