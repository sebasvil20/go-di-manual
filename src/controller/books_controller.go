package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebasvil20/go-di/src/service"
)

type IBooksController interface {
	GetBook(c *gin.Context)
}

type BooksController struct {
	BooksService service.IBookService
}

func (ctrl *BooksController) GetBook(c *gin.Context) {
	book, err := ctrl.BooksService.GetBook(c, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}
