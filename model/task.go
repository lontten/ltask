package model

type Task struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Script      string `json:"script"`
	Timeout     int64  `json:"timeout"`
	Status      int64  `json:"status"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}
type TaskLog struct {
	Id        int64  `json:"id"`
	TaskId    int64  `json:"task_id"`
	Log       string `json:"log"`
	CreatedAt int64  `json:"created_at"`
}
