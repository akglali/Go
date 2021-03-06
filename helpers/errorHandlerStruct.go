package helpers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

type ErrorStruct struct {
	Error string
}

func MyAbort(c *gin.Context, str string) {
	c.AbortWithStatusJSON(400, ErrorStruct{Error: str})
}
