package pkg

import (
	"github.com/gin-gonic/gin"
)

type Responder interface {
	OK(httpCode int, data interface{})

	OkWithPage(httpCode int, list interface{}, pagination *Pagination)

	Fail(err error)
}

type responderX struct {
	c *gin.Context
}

func NewResponse(c *gin.Context) Responder {
	return &responderX{c: c}
}

func (r *responderX) OK(httpCode int, data interface{}) {

	r.c.JSON(httpCode, gin.H{
		"status":  httpCode,
		"success": true,
		"data":    data,
	})
}

func (r *responderX) OkWithPage(httpCode int, list interface{}, pagination *Pagination) {

	r.c.JSON(httpCode, gin.H{
		"status":  httpCode,
		"success": true,
		"data": gin.H{
			"list":       list,
			"pagination": pagination,
		},
	})
}

func (r *responderX) Fail(err error) {

	r.c.JSON(200, gin.H{
		"status":  500,
		"success": false,
		"error":   err.Error(),
	})
}
