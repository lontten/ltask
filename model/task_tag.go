package model

import "time"

// 标签值类型常量
const (
	TagValueTypeString = 1 // 字符串类型
	TagValueTypeNumber = 2 // 数值类型
)

// TaskTag 任务标签模型
type TaskTag struct {
	TaskTagID    int64     `db:"task_tag_id" json:"taskTagId"`       // 标签自增ID
	TagID        int64     `db:"tag_id" json:"tagId"`                // 关联任务ID
	CreatedAt    time.Time `db:"created_at" json:"createdAt"`        // 创建时间(数据库自动维护)
	UpdatedAt    time.Time `db:"updated_at" json:"updatedAt"`        // 更新时间(数据库自动维护)
	TagKey       string    `db:"tag_key" json:"tagKey"`              // 标签键
	TagValue     string    `db:"tag_value" json:"tagValue"`          // 标签原始值
	TagValueType int       `db:"tag_value_type" json:"tagValueType"` // 值类型:1-字符串,2-数值
}

func (TaskTag) TableConf() *ldb.TableConf {
	return new(ldb.TableConf).
		Table("ltask_task_tag").
		PrimaryKeys("task_tag_id").
		AutoIncrements("task_tag_id")
}
