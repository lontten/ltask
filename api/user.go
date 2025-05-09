package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserBo struct {
	Account  string `db:"account" json:"account"`   // 登录账号
	Password string `db:"password" json:"password"` // 密码
}

func login(c *gin.Context) {
	bo := UserBo{}
	if err := c.ShouldBindJSON(&bo); err == nil {
		c.String(http.StatusOK, `the body should be formA`)
	}

}
