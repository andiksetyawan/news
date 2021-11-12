package exception

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"news/model"
)

func HandlerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, model.WebResponse{
		Success: false,
		Message: err.Error(),
		Data:    gin.H{},
	})
	c.Abort()
}
