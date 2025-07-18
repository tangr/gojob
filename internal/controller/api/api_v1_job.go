package api

import (
	"context"
	"gojob/internal/dao"
	"gojob/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
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
