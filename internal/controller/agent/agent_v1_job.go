package agent

import (
	"context"
	"gojob/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// semaphore
// var sem = make(chan struct{}, 1)

func (c *ControllerV1) JobRun(ctx context.Context, req *JobRunReq) (response *ghttp.Response, err error) {
	r := g.RequestFromCtx(ctx)

	job_id := r.Get("id").String()
	g.Log().Debug(ctx, "job_id: ", job_id)

	// pipeline_body, err := service.Task.Run(job_id)
	// if err != nil {
	// 	return nil, err
	// }
	// g.Log().Debug(ctx, "pipeline_body: ", pipeline_body)

	task_run_result := service.Task.Run(job_id)
	if task_run_result {
		taskRunResult := map[string]interface{}{
			"code":    200,
			"message": "success",
			"data":    "task trigger success",
		}
		r.Response.WriteJsonExit(taskRunResult)
	} else {
		r.Response.WriteStatusExit(429, "too many concurrent commands now, plz retry after a while!!")
	}

	// select {
	// case sem <- struct{}{}:
	// 	// 获取执行权限
	// 	defer func() { <-sem }()
	// 	// 执行命令
	// 	task_run_result := service.Task.Run(job_id)
	// 	if err != nil {
	// 		r.Response.WriteStatusExit(500, err.Error())
	// 		return nil, err
	// 	}
	// 	g.Log().Debug(ctx, "task_run_result: ", task_run_result)

	// 	r.Response.Write(task_run_result)

	// case <-ctx.Done():
	// 	r.Response.WriteStatusExit(429, "too many concurrent commands")
	// }

	// r.Response.WriteExit(pipeline_body)
	return nil, nil
}

func (c *ControllerV1) JobAbort(ctx context.Context, req *JobAbortReq) (response *ghttp.Response, err error) {
	r := g.RequestFromCtx(ctx)

	job_id := r.Get("id").String()
	g.Log().Debug(ctx, "job_id: ", job_id)

	task_run_result := service.Task.Abort(job_id)
	if task_run_result {
		taskRunResult := map[string]interface{}{
			"code":    200,
			"message": "success",
			"data":    "task abort success",
		}
		r.Response.WriteJsonExit(taskRunResult)
	} else {
		r.Response.WriteStatusExit(429, "job abort failed, plz retry!!!")
	}

	return nil, err
}
