package routes

import (
	"gogin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func lists(c *gin.Context) {
	if len(models.Data) == 0 {
		c.IndentedJSON(http.StatusNoContent, "204 no content")
	} else {
		c.IndentedJSON(http.StatusOK, models.Data)
	}
}

func todos(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.IndentedJSON(http.StatusNotFound, "404 resource not found")
		} else {
			c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos)
		}
}

func getList(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.IndentedJSON(http.StatusNotFound, "404 resource not found")
		} else {
			c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")])
		}

}

func createList(c *gin.Context) {
	requestBody := new(models.RequestBodyList)
	requestBodyTodos := make(map[string]*models.ToDo)
	todoListKey := uuid.New().String()

	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		check(err, c)
	}

	for _, v := range requestBody.Todos {
		toDosKey := uuid.New().String()
		requestBodyTodos[toDosKey] = new(models.ToDo)
		requestBodyTodos[toDosKey].Content = v
		requestBodyTodos[toDosKey].Id = toDosKey
	}

	models.Data[todoListKey] = new(models.ToDoList)
	models.Data[todoListKey].Id = todoListKey
	models.Data[todoListKey].Owner = requestBody.Owner
	models.Data[todoListKey].Todos = requestBodyTodos

	c.IndentedJSON(http.StatusCreated, models.Data[todoListKey])
}

func updateList(c *gin.Context) {

	_, hasList := models.Data[c.Param("listid")]
	if !hasList {
		c.IndentedJSON(http.StatusNotFound, "404 resource not found")
		return
	}

	requestBody := new(models.ToDoList)

	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		check(err, c)
	}

	models.Data[c.Param("listid")].Owner = requestBody.Owner

	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")])
}


func deleteList(c *gin.Context) {

	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.IndentedJSON(http.StatusNotFound, "404 resource not found")
		} else {
			delete(models.Data, c.Param("listid"))
		}


}


func getToDo(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.IndentedJSON(http.StatusNotFound, "404 list not found")
		}

		if hasList {
			_, hasTodo := models.Data[c.Param("listid")].Todos[c.Param("todoid")]
			if !hasTodo {
				c.IndentedJSON(http.StatusNotFound, "404 todo not found")
			} else {
				c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos[c.Param("todoid")])
			}
		}

}


func deleteToDo(c *gin.Context) {
	
	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.IndentedJSON(http.StatusNotFound, "404 list not found")
		}

		if hasList {
			_, hasTodo := models.Data[c.Param("listid")].Todos[c.Param("todoid")]
			if !hasTodo {
				c.IndentedJSON(http.StatusNotFound, "404 todo not found")
			} else {
				delete(models.Data[c.Param("listid")].Todos, c.Param("todoid"))
			}
		}

}

func updateToDo(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]

		if !hasList {
			c.IndentedJSON(http.StatusNotFound, "404 list not found")
			return
		} else {
			_, hasTodo := models.Data[c.Param("listid")].Todos[c.Param("todoid")]
			if !hasTodo {
				c.IndentedJSON(http.StatusNotFound, "404 todo not found")
				return
			}
		}

	requestBody := new(models.ToDo)

		if err := c.ShouldBindJSON(requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		check(err, c)
	}


	models.Data[c.Param("listid")].Todos[c.Param("todoid")].Content = requestBody.Content


	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos[c.Param("todoid")])
}

func createToDo(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.IndentedJSON(http.StatusNotFound, "404 list not found")
			return
		}

	requestBody := new(models.ToDo)

	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		check(err, c)
	}

	key := uuid.New().String()
	requestBody.Id = key
	models.Data[c.Param("listid")].Todos[key] = requestBody


	c.IndentedJSON(http.StatusCreated, models.Data[c.Param("listid")].Todos[key])
}