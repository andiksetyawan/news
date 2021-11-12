package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"news/exception"
	"news/model"
	"news/service"
)

type tagCont struct {
	tagServ service.TagService
}

func NewTagCont(service service.TagService) *tagCont {
	return &tagCont{tagServ: service}
}

func (controller tagCont) Gets(c *gin.Context) {
	tags, err := controller.tagServ.FindTags(c, nil)
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

func (controller tagCont) Create(c *gin.Context) {
	var payload model.TagCreateRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		exception.HandlerError(c, err)
		return
	}

	err := controller.tagServ.CreateTag(c, payload)
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

func (controller tagCont) FindTagBySlug(c *gin.Context) {
	slug := c.Param("slug")
	tags, err := controller.tagServ.FindTagBySlug(c, slug)
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

func (controller tagCont) UpdateBySlug(c *gin.Context) {
	slug := c.Param("slug")
	var payload model.TagCreateRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		exception.HandlerError(c, err)
		return
	}

	err := controller.tagServ.UpdateTagBySlug(c, payload, slug)
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

func (controller tagCont) DeleteBySlug(c *gin.Context) {
	slug := c.Param("slug")
	err := controller.tagServ.DeleteTagBySlug(c, slug)
	if err != nil {
		exception.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Success: true,
		Message: "OK",
		Data:    gin.H{"tag_slug": slug},
	})
}
