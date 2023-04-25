package controllers

import (
	"gogin-api/models"
	"gogin-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service service.ToDoListServiceInterface
}

func NewController(service service.ToDoListServiceInterface) *Controller {
	return &Controller{
		Service: service,
	}
}

func (c *Controller) GetAllListsController(ctx *gin.Context) {
	lists, err := c.Service.GetAllLists()

	if err != nil {
		ctx.Status(http.StatusNoContent)
		return
	}

	ctx.JSON(http.StatusOK, lists)
}

func (c *Controller) GetListController(ctx *gin.Context) {
	list, err := c.Service.GetList(ctx.Param("listid"))

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, list)
}

func (c *Controller) CreateListController(ctx *gin.Context) {
	requestBody := new(models.RequestBodyList)

	if err := ctx.BindJSON(requestBody); err != nil {
		return
	}

	c.Service.CreateList(requestBody)

	ctx.Status(http.StatusCreated)
}

func (c *Controller) PatchListController(ctx *gin.Context) {
	requestBody := new(models.ToDoList)

	if err := ctx.BindJSON(requestBody); err != nil {
		return
	}

	requestBody.Id = ctx.Param("listid")

	err := c.Service.PatchList(requestBody)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *Controller) DeleteListController(ctx *gin.Context) {
	err := c.Service.DeleteList(ctx.Param("listid"))

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *Controller) GetToDoController(ctx *gin.Context) {
	todo, err := c.Service.GetToDoInList(ctx.Param("todoid"))

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (c *Controller) DeleteToDoController(ctx *gin.Context) {
	err := c.Service.DeleteToDoInList(ctx.Param("todoid"))

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *Controller) PatchToDoController(ctx *gin.Context) {
	requestBody := new(models.ToDo)

	if err := ctx.BindJSON(requestBody); err != nil {
		return
	}

	err := c.Service.PatchToDoInList(requestBody.Completed, ctx.Param("todoid"))

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *Controller) CreateToDoController(ctx *gin.Context) {
	requestBody := new(models.ToDo)

	if err := ctx.BindJSON(requestBody); err != nil {
		return
	}

	if requestBody.Content == "" {
		ctx.JSON(http.StatusBadRequest, "content can't be empty")
		return
	}

	if err := c.Service.CreateToDoInList(ctx.Param("listid"), requestBody.Content); err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c *Controller) GetAllToDosController(ctx *gin.Context) {
	todos, err := c.Service.GetAllToDosInList(ctx.Param("listid"))

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

func (c *Controller) GetDataStructureController(ctx *gin.Context) {
	dataStructure, err := c.Service.GetDataStructure()

	if err != nil {
		ctx.JSON(http.StatusNoContent, err.Error())
		return
	}

	ctx.JSON(200, dataStructure)
}
