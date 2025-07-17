package service

var Task = taskService{}

type taskService struct{}

type TaskResult struct {
	Status string `json:"status"`
}

func (s *taskService) Run(job_id string) (*TaskResult, error) {
	return &TaskResult{
		Status: "status",
	}, nil
}

func (s *taskService) Abort(job_id string) (*TaskResult, error) {
	return &TaskResult{
		Status: "status",
	}, nil
}
