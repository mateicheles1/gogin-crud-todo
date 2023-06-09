package service

import "gogin-api/models"

type ToDoListServiceInterface interface {
	CreateList(reqBody *models.RequestBodyList) string
	PatchList(list *models.ToDoList) error
	GetList(id string) (models.ResponseBodyList, error)
	DeleteList(key string) error
	GetAllLists() ([]models.ResponseBodyList, error)

	CreateToDoInList(listId string, content string) (string, error)
	PatchToDoInList(completed bool, id string) error
	GetToDoInList(key string) (*models.ToDo, error)
	DeleteToDoInList(key string) error
	GetAllToDosInList(listId string) ([]models.ToDo, error)

	GetDataStructure() (map[string]*models.ToDoList, error)
}
