package agent

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) Agent(ctx context.Context, req *AgentReq) (response *ghttp.Response, err error) {
	r := g.RequestFromCtx(ctx)

	err = r.Response.WriteTpl("cicd/list.html", g.Map{
		"Title":          "111欢迎页面",
		"Content":        "aaaa:bbb:ccc",
		"url":            "/",
		"pipelines":      "pipelines",
		"newPipelineUrl": "/pipelines/new",
	})
	return nil, err
}
