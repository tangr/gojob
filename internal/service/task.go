package service

import (
	"context"
	"fmt"
	"gojob/internal/logic/agent"
	"gojob/internal/model"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
)

var Task = taskService{}

type taskService struct{}

type TaskResult struct {
	Status string `json:"status"`
	JobId  string `json:"jobid"`
}

type ScriptObj = model.Script

func (s *taskService) Run(job_id string) (*TaskResult, error) {
	ctx := context.Background()

	jobId, err := strconv.Atoi(job_id)
	if err != nil {
		fmt.Println("转换失败:", err)
	}

	agent.AgentCICD.HandleRecvJson2(ctx, jobId)

	// var script_obj ScriptObj = agent.AgentCICD.GetScriptByTask(jobId)
	// script_body := script_obj.Body
	// script_envs := script_obj.Envs
	// script_args := script_obj.Args

	// g.Log().Debugf(ctx, "HandleJob GetScriptByTask script_body: %s", script_body)
	// g.Log().Debugf(ctx, "HandleJob GetScriptByTask script_args: %s", script_args)
	// g.Log().Debugf(ctx, "HandleJob GetScriptByTask script_envs: %s", script_envs)

	g.Log().Debugf(ctx, "HandleJob GetScriptByTask MaxRunningJobs: %d", agent.MaxRunningJobs)

	return &TaskResult{
		Status: "status",
		JobId:  job_id,
	}, nil
}

func (s *taskService) Abort(job_id string) (*TaskResult, error) {
	return &TaskResult{
		Status: "status",
	}, nil
}
