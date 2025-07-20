package api

import (
	"context"
	"gojob/internal/dao"
	"gojob/internal/model/do"
	"gojob/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) JobScript(ctx context.Context, req *JobScriptReq) (response *ghttp.Response, err error) {
	r := g.RequestFromCtx(ctx)

	job_id := r.Get("id").Int()
	g.Log().Debug(ctx, "job_id: ", job_id)

	pipeline_body, err := service.Script.GetOne(job_id)
	if err != nil {
		return nil, err
	}
	g.Log().Debug(ctx, "pipeline_body: ", pipeline_body)

	r.Response.WriteExit(pipeline_body)
	return nil, nil
}

func (c *ControllerV1) GetOne(ctx context.Context, req *GetOneReq) (res *GetOneRes, err error) {
	r := g.RequestFromCtx(ctx)

	job_id := r.Get("id").Int()
	g.Log().Debug(ctx, "job_id: ", job_id)

	res = &GetOneRes{}
	err = dao.CicdJob.Ctx(ctx).WherePri(req.Id).Scan(&res.CicdJob)
	return
}

func (c *ControllerV1) JobLog(ctx context.Context, req *JobLogReq) (response *ghttp.Response, err error) {
	r := g.RequestFromCtx(ctx)
	g.Log().Debugf(ctx, "sendMap: %s", gconv.String(req))

	job_id := r.Get("id").String()

	g.Log().Debug(ctx, "JobLog req: ", req)

	// _, err = dao.CicdLog.Ctx(ctx).Data(do.CicdLog{
	// 	PipelineId: req.PipelineId,
	// 	AgentId:    req.AgentId,
	// 	JobType:    req.JobType,
	// 	JobId:      req.JobId,
	// 	TaskStatus: req.TaskStatus,
	// 	Ipaddr:     req.Ipaddr,
	// 	UpdatedAt:  req.UpdatedAt,
	// 	Output:     req.Output,
	// }).WherePri(job_id).Update()
	// if err != nil {
	// 	g.Log().Debug(ctx, "JobLog CicdLog err: ", err)
	// 	// return nil, err
	// }

	_, err = dao.CicdJob.Ctx(ctx).Data(do.CicdJob{
		JobStatus: req.TaskStatus,
		Output:    req.Output,
	}).WherePri(job_id).Update()
	if err != nil {
		g.Log().Debug(ctx, "JobLog CicdJob err: ", err)
		// return nil, err
	}

	return
}

func (c *ControllerV1) JobGetLog(ctx context.Context, req *JobGetLogReq) (response *ghttp.Response, err error) {
	r := g.RequestFromCtx(ctx)

	pipeline_id := r.Get("pipeline_id").Int()
	job_id := r.Get("job_id").Int()
	g.Log().Debug(ctx, "JobGetLog pipeline_id job_id: ", pipeline_id, job_id)

	pipeline_body, err := service.Cicd.JobGetLog(ctx, pipeline_id, job_id)
	if err != nil {
		return nil, err
	}
	g.Log().Debug(ctx, "pipeline_body: ", pipeline_body)

	r.Response.WriteExit(pipeline_body)
	return nil, nil
}
