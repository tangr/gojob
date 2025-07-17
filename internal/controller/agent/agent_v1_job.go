package agent

import (
	"context"
	"gojob/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) JobRun(ctx context.Context, req *JobRunReq) (response *ghttp.Response, err error) {
	r := g.RequestFromCtx(ctx)

	job_id := r.Get("id").String()
	g.Log().Debug(ctx, "job_id: ", job_id)

	pipeline_body, err := service.Task.Run(job_id)
	if err != nil {
		return nil, err
	}

	r.Response.WriteExit(pipeline_body)
	return nil, nil
}

func (c *ControllerV1) JobAbort(ctx context.Context, req *JobAbortReq) (response *ghttp.Response, err error) {
	r := g.RequestFromCtx(ctx)

	job_id := r.Get("id").String()
	g.Log().Debug(ctx, "job_id: ", job_id)

	err = r.Response.WriteTpl("cicd/list.html", g.Map{
		"Title":          "111欢迎页面",
		"Content":        "aaaa:bbb:ccc",
		"url":            "/",
		"pipelines":      "pipelines",
		"newPipelineUrl": "/pipelines/new",
	})
	return nil, err
}
