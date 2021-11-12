package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"news/exception"
	"news/model"
	"news/service"
)

type newsCont struct {
	newsServ service.NewsService
}

func NewNewsCont(service service.NewsService) *newsCont {
	return &newsCont{newsServ: service}
}

func (controller newsCont) Gets(c *gin.Context) {
	topic := c.Query("topic")
	status := c.Query("status")
	tags, err := controller.newsServ.FindNews(c, topic, status)
	if err != nil {
		exception.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Success: true,
		Message: "OK",
		Data:    tags,
	})
}

func (controller newsCont) Get(c *gin.Context) {
	id := c.Param("id")
	tags, err := controller.newsServ.FindNewsByID(c, id)
	if err != nil {
		exception.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Success: true,
		Message: "OK",
		Data:    tags,
	})
}

func (controller newsCont) Create(c *gin.Context) {
	var payload model.NewsCreateRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		exception.HandlerError(c, err)
		return
	}

	err := controller.newsServ.CreateNews(c, payload)
	if err != nil {
		exception.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Success: true,
		Message: "OK",
		Data:    gin.H{},
	})
}

func (controller newsCont) UpdateByID(c *gin.Context) {
	id := c.Param("id")
	var payload model.NewsUpdateRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		exception.HandlerError(c, err)
		return
	}

	err := controller.newsServ.UpdateNewsByID(c, payload, id)
	if err != nil {
		exception.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Success: true,
		Message: "OK",
		Data:    gin.H{},
	})
}

func (controller newsCont) FindNewsByID(c *gin.Context) {
	id := c.Param("id")
	tags, err := controller.newsServ.FindNewsByID(c, id)
	if err != nil {
		exception.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Success: true,
		Message: "OK",
		Data:    tags,
	})
}

func (controller newsCont) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	err := controller.newsServ.DeleteNewsByID(c, id)
	if err != nil {
		exception.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Success: true,
		Message: "OK",
		Data:    gin.H{"id": id},
	})
}
