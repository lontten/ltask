package r

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, data any, msg ...string) {
	resp, _ := Data(0, data, msg...)
	c.JSON(http.StatusOK, resp)
}

func R(c *gin.Context, code int, data any, msg ...string) {
	resp, _ := Data(code, data, msg...)
	c.JSON(http.StatusOK, resp)
}

func E(c *gin.Context, err error) {
	resp, _ := Data(-1, "", err.Error())
	c.JSON(http.StatusOK, resp)
}
