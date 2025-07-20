package api

import (
	"context"
	"gojob/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type JobScriptReq struct {
	g.Meta `path:"/job/script/{id}" method:"get" tags:"job script" summary:"JobScriptReq"`
}
type JobAbortReq struct {
	g.Meta `path:"/job/log/{id}" method:"post" tags:"job log" summary:"JobAbortReq"`
}

type GetOneReq struct {
	g.Meta `path:"/job/{id}" method:"get" tags:"api" summary:"Get one job"`
	Id     int64 `v:"required" dc:"job id"`
}

type GetOneRes struct {
	*entity.CicdJob `dc:"job"`
}

type JobLogReq struct {
	g.Meta `path:"/log/{id}" tags:"job log" method:"put" summary:"JobLogReq"`

	Id         uint64 `v:"required" json:"id" dc:"Log ID"`
	PipelineId int    `v:"required-without:agentId" json:"pipelineId" dc:"Pipeline ID"`
	AgentId    int    `v:"required-without:pipelineId" json:"agentId" dc:"Agent ID"`
	JobType    string `v:"required|in:BUILD,DEPLOY" json:"jobType" dc:"Job type"`
	JobId      int    `v:"required" json:"jobId" dc:"Job ID"`
	TaskStatus string `v:"required|in:pending,running,failed,success" json:"taskStatus" dc:"Task status"`
	Ipaddr     string `v:"required" json:"ipaddr" dc:"IP address"`
	UpdatedAt  int64  `v:"required" json:"updatedAt" dc:"Update timestamp"`
	Output     string `json:"output" dc:"Output content"`
}

type JobGetLogReq struct {
	g.Meta `path:"/job/{pipeline_id}/{job_id}/log" tags:"job log" method:"get" summary:"JobGetLogReq"`
}

type IApiV1 interface {
	JobScript(ctx context.Context, req *JobScriptReq) (res *ghttp.Response, err error)
	JobLog(ctx context.Context, req *JobLogReq) (res *ghttp.Response, err error)
}

type ControllerV1 struct{}

func NewV1() IApiV1 {
	return &ControllerV1{}
}
