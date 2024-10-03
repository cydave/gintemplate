package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExampleController struct{}

func (e *ExampleController) Get(c *gin.Context) {
	c.String(http.StatusOK, "42")
}

func (e *ExampleController) Post(c *gin.Context) {
	c.String(http.StatusOK, "42")
}

func (e *ExampleController) Put(c *gin.Context) {
	c.String(http.StatusOK, "42")
}
