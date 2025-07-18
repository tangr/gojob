package api

import (
	"context"
	"gojob/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type JobScriptReq struct {
	g.Meta `path:"/job/script/{id}" tags:"job script" method:"get" summary:"JobScriptReq"`
}
type JobAbortReq struct {
	g.Meta `path:"/job/log/{id}" tags:"job log" method:"post" summary:"JobLogReq"`
}

type GetOneReq struct {
	g.Meta `path:"/job/{id}" method:"get" tags:"api" summary:"Get one job"`
	Id     int64 `v:"required" dc:"job id"`
}

type GetOneRes struct {
	*entity.CicdJob `dc:"job"`
}

type IApiV1 interface {
	JobScript(ctx context.Context, req *JobScriptReq) (res *ghttp.Response, err error)
}

type ControllerV1 struct{}

func NewV1() IApiV1 {
	return &ControllerV1{}
}
