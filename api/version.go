package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ltask/common/r"
)

func Version(c *gin.Context) {
	r.Success(c, "0.0.1")
	return
}
