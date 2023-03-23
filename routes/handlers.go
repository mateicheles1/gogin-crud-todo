package routes

import (
	"gogin-api/logs"
	"gogin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


var logger = logs.Logger()

func check(err error, c *gin.Context) {
	logger.Panic().
		Str("Method", c.Request.Method).
		Str("Path", c.Request.URL.Path).
		Int("Status code", http.StatusBadRequest).
        Msgf("Couldn't unmarshal the request body into the requestBody struct due to: %s", err)
}

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
	// un map de todo-uri care va primi pe key todo struct-ul din todolist struct-ul requestBody
	requestBodyTodos := make(map[string]*models.ToDo)
	// key din models.Data si id-ul struct-ului de todolist
	toDoListKey := uuid.New().String()
	
	
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		check(err, c)
	}
	
	if requestBody.Todos == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "todos can't be empty"})
		return
	}
	

	for k := range requestBody.Todos {
		toDosKey := uuid.New().String()
		requestBody.Todos[k].Id = toDosKey
		requestBody.Todos[k].Listid = toDoListKey
		// map-ul map[toDosKey]*ToDo primeste struct-urile de todo-uri din request body
		requestBodyTodos[toDosKey] = requestBody.Todos[k]
	}

	requestBody.Id = toDoListKey
	models.Data[toDoListKey] = requestBody
	models.Data[toDoListKey].Todos = requestBodyTodos
	
	logger.Info().
	Str("method", c.Request.Method).
	Int("status code", http.StatusOK).
	Str("path", c.Request.URL.Path).
	Msg("list successfully created")

	c.IndentedJSON(http.StatusOK, models.Data)

}


func updateList(c *gin.Context) {
	requestBody := new(models.ToDoList)

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		check(err, c)
	}

	models.Data[c.Param("listid")].Owner = requestBody.Owner

	logger.Info().
	Str("method", c.Request.Method).
	Int("status code", http.StatusOK).
	Msg("list successfully updated")

	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")])
}

func deleteList(c *gin.Context) {
	delete(models.Data, c.Param("listid"))

	logger.Info().
	Str("method", c.Request.Method).
	Int("status code", http.StatusOK).
	Str("path", c.Request.URL.Path).
	Msg("list successfully deleted")

	c.IndentedJSON(http.StatusOK, models.Data)
}


func getToDo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK,models.Data[c.Param("listid")].Todos[c.Param("todoid")])
}

func deleteToDo(c *gin.Context) {
	delete(models.Data[c.Param("listid")].Todos, c.Param("todoid"))

	logger.Info().
	Str("method", c.Request.Method).
	Int("status code", http.StatusOK).
	Msg("todo successfully deleted")

	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos)
}

func updateToDo(c *gin.Context) {
	requestBody := new(models.ToDo)

		if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		check(err, c)
	}


	models.Data[c.Param("listid")].Todos[c.Param("todoid")].Content = requestBody.Content

	logger.Info().
	Str("method", c.Request.Method).
	Int("status code", http.StatusOK).
	Msg("todo successfully updated")

	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos[c.Param("todoid")])
}

func createToDo(c *gin.Context) {
	requestBody := new(models.ToDo)

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		check(err, c)
	}

	key := uuid.New().String()
	requestBody.Listid = c.Param("listid")
	requestBody.Id = key
	models.Data[c.Param("listid")].Todos[key] = requestBody

	logger.Info().
	Str("method", c.Request.Method).
	Int("status code", http.StatusOK).
	Msg("todo successfully created")

	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos[key])
}