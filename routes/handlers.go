package routes

import (
	"gogin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

func lists(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Data)
}

func todos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos)
}

func getList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")])
}

func createList(c *gin.Context) {
	// request body 
	requestBody := new(models.ToDoList)
	// un map de todolist-uri care va primi pe key todolist struct-ul requestBody
	mapCopyRequestBody := make(map[string]*models.ToDoList)
	// un map de todo-uri care va primi pe key todo struct-ul din todolist struct-ul requestBody
	requestBodyTodos := make(map[string]*models.ToDo)
	// key din models.Data si id-ul struct-ului de todolist
	requestBodyKey := uuid.New().String()
	
	
	if err := c.ShouldBindWith(&requestBody, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if requestBody.Todos == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "todos can't be empty"})
		return
	}

	for k := range requestBody.Todos {
		toDosKey := uuid.New().String()
		requestBody.Todos[k].Id = toDosKey
		requestBody.Todos[k].Listid = requestBodyKey
		// map-ul map[toDosKey]*ToDo primeste struct-urile de todo-uri din request body
		requestBodyTodos[toDosKey] = requestBody.Todos[k]
	}

	requestBody.Id = requestBodyKey
	// map-ul map[requestBodyKey]*ToDoList primeste atat struct ul de todolist cat si map-ul de todos
	mapCopyRequestBody[requestBodyKey] = requestBody
	mapCopyRequestBody[requestBodyKey].Todos = requestBodyTodos
	// varibila Data primeste pe key-ul de requestBodyKey struct-ul de todolist
	models.Data[requestBodyKey] = mapCopyRequestBody[requestBodyKey]
	c.IndentedJSON(http.StatusOK, models.Data)

}


func updateList(c *gin.Context) {
	requestBody := new(models.ToDoList)

	if err := c.ShouldBindWith(&requestBody, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.Data[c.Param("listid")].Owner = requestBody.Owner
	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")])
}

func deleteList(c *gin.Context) {
	delete(models.Data, c.Param("listid"))
	c.IndentedJSON(http.StatusOK, models.Data)
}


func getToDo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK,models.Data[c.Param("listid")].Todos[c.Param("todoid")])
}

func deleteToDo(c *gin.Context) {
	delete(models.Data[c.Param("listid")].Todos, c.Param("todoid"))
	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos)
}

func updateToDo(c *gin.Context) {
	requestBody := new(models.ToDo)

		if err := c.ShouldBindWith(&requestBody, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}


	models.Data[c.Param("listid")].Todos[c.Param("todoid")].Content = requestBody.Content
	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos[c.Param("todoid")])
}

func createToDo(c *gin.Context) {
	requestBody := new(models.ToDo)

	if err := c.ShouldBindWith(&requestBody, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	key := uuid.New().String()
	requestBody.Listid = c.Param("listid")
	requestBody.Id = key
	models.Data[c.Param("listid")].Todos[key] = requestBody
	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos[key])
}