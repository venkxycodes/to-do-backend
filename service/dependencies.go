package service

import (
	"to-do/appcontext"
	_ "to-do/appcontext"
	"to-do/repo"
)

type ServerDependencies struct {
	ToDoService ToDoService
}

func InstantiateServerDependencies() *ServerDependencies {
	dbClient := appcontext.GetDBClient()
	toDoRepo := repo.NewToDoRepo(dbClient)
	toDoServ := NewToDoService(toDoRepo)
	return &ServerDependencies{
		ToDoService: toDoServ,
	}
}
