package freeagent

type TaskStatus string

const (
	TaskStatusActive    TaskStatus = "active"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusHidden    TaskStatus = "hidden"
)

type TaskBillingPeriod string

const (
	TaskBillingPeriodDay  TaskBillingPeriod = "day"
	TaskBillingPeriodHour TaskBillingPeriod = "hour"
)

type Task struct {
	URL           string            `json:"url,omitempty"`
	Name          string            `json:"name"`
	IsBillable    bool              `json:"is_billable"`
	Status        TaskStatus        `json:"status"`
	CreatedAt     string            `json:"created_at"`
	UpdatedAt     string            `json:"updated_at"`
	BillingRate   string            `json:"billing_rate"`
	BillingPeriod TaskBillingPeriod `json:"billing_period"`
}

type taskDTO struct {
	Task *Task `json:"task"`
}

func (c *FreeAgent) PostTask(task *Task) (*Task, error) {
	request := &taskDTO{task}
	response := &taskDTO{}
	err := c.post("/tasks", request, response)
	if err != nil {
		return nil, err
	}

	return response.Task, nil
}

func (c *FreeAgent) GetTask(id string) (*Task, error) {
	result := &taskDTO{}
	err := c.get("/tasks/"+id, result)
	if err != nil {
		return nil, err
	}

	return result.Task, nil
}

type TaskView string

const (
	TaskViewAll       TaskView = "all"
	TaskViewActive    TaskView = "active"
	TaskViewCompleted TaskView = "completed"
	TaskViewHidden    TaskView = "hidden"
)

type TaskSort string

const (
	TaskSortName        TaskSort = "name"
	TaskSortProject     TaskSort = "project"
	TaskSortBillingRate TaskSort = "billing_rate"
	TaskSortCreatedAt   TaskSort = "created_at"
	TaskSortUpdatedAt   TaskSort = "updated_at"
)

type TaskQuery struct {
	UpdatedSince string
	View         TaskView
	Sort         TaskSort
}

type tasksDTO struct {
	Tasks []*Task `json:"tasks"`
}

func (c *FreeAgent) GetTasks(q *TaskQuery) ([]*Task, error) {
	result := &tasksDTO{}
	err := c.get("/tasks", result)
	if err != nil {
		return nil, err
	}

	return result.Tasks, nil
}
