package model

import (
	"github.com/lontten/ldb"
	"time"
)

const (
	TaskStatusPending = 1 // 未执行
	TaskStatusRunning = 2 // 执行中
	TaskStatusSuccess = 3 // 执行成功
	TaskStatusFailed  = 4 // 执行失败

	FailTypeNone         = 0 // 无
	FailTypeException    = 1 // 异常
	FailTypeTimeout      = 2 // 超时
	FailTypeManualCancel = 3 // 手动取消

	ExecTypeShell = 1 // shell
	ExecTypeHTTP  = 2 // http

	HTTPMethodGet    = 1 // GET
	HTTPMethodPost   = 2 // POST
	HTTPMethodPut    = 3 // PUT
	HTTPMethodDelete = 4 // DELETE
	HTTPMethodPatch  = 5 // PATCH
)

type Task struct {
	TaskID      int64     `db:"task_id" json:"taskId"`           // task 主键
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`     // 创建时间
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`     // 更新时间
	TaskName    string    `db:"task_name" json:"taskName"`       // task 名字
	ExecContent string    `db:"exec_content" json:"execContent"` // 执行内容
	TaskStatus  int       `db:"task_status" json:"taskStatus"`   // task 状态：1未执行；2执行中；3执行成功；4执行失败
	FailType    int       `db:"fail_type" json:"failType"`       // task 失败类型：0无；1异常；2超时；3手动取消
	ExecType    int       `db:"exec_type" json:"execType"`       // 执行方式：1shell；2http
	HTTPMethod  int       `db:"http_method" json:"httpMethod"`   // http请求方法：1get；2post；3put；4delete;5patch
	TimeOut     int64     `db:"time_out" json:"timeOut"`         // 超时时间；单位秒，默认0 不限制
	Remark      string    `db:"remark" json:"remark"`            // 备注
}

func (Task) TableConf() *ldb.TableConf {
	return new(ldb.TableConf).
		Table("ltask_task").
		PrimaryKeys("task_id").
		AutoIncrements("task_id")
}
