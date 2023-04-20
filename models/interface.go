package models

type ToDoListService interface {
	CreateList(reqBody *RequestBodyList, todos map[string]*ToDo, key string)
	PatchList(list *ToDoList) error
	GetList(id string) (ResponseBodyList, error)
	DeleteList(key string) error
	GetAllLists() ([]ResponseBodyList, error)

	CreateToDoInList(listId string, content string) error
	PatchToDoInList(content string, id string) error
	GetToDoInList(key string) (*ToDo, error)
	DeleteToDoInList(key string) error
	GetAllToDosInList(listId string) ([]ToDo, error)
}