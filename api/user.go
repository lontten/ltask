package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ltask/common/r"
)

type UserBo struct {
	Account  string `db:"account" json:"account"`   // 登录账号
	Password string `db:"password" json:"password"` // 密码
}

func Login(c *gin.Context) {
	bo := UserBo{}
	if err := c.ShouldBindJSON(&bo); err == nil {
		r.E(c, err)
		return
	}
	r.Success(c, bo)
	return
}

func Hello(c *gin.Context) {
	r.Success(c, "hello")
	return
}
