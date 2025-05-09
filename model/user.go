package model

import (
	"github.com/lontten/ldb"
	"time"
)

// 用户状态常量
const (
	UserStatusEnabled  = 1 // 账号启用
	UserStatusDisabled = 2 // 账号停用
)

type User struct {
	UserID     int64     `db:"user_id" json:"userId"`         // 管理员主键ID
	CreatedAt  time.Time `db:"created_at" json:"createdAt"`   // 创建时间(数据库自动维护)
	UpdatedAt  time.Time `db:"updated_at" json:"updatedAt"`   // 更新时间(数据库自动维护)
	Account    string    `db:"account" json:"account"`        // 登录账号(唯一)
	Email      string    `db:"email" json:"email"`            // 绑定邮箱(唯一)
	UserStatus int       `db:"user_status" json:"userStatus"` // 账号状态:1-启用,2-停用
	Password   string    `db:"password" json:"-"`             // 密码(不序列化到JSON)
}

func (User) TableConf() *ldb.TableConf {
	return new(ldb.TableConf).
		Table("ltask_user").
		PrimaryKeys("user_id").
		AutoIncrements("user_id")
}
