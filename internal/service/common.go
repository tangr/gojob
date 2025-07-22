package service

import (
	"context"
	"encoding/json"
	"fmt"
	"gojob/internal/dao"
	"math/rand"

	"github.com/gogf/gf/v2/frame/g"
)

var Comm = commService{}

type commService struct{}

func (s *commService) ParseEnvs(envs map[string]interface{}) map[string]string {
	var new_envs map[string]string = make(map[string]string)
	for k, v := range envs {
		switch v := v.(type) {
		case string:
			new_envs[k] = v
		case []interface{}:
			new_v := ""
			for _, u := range v {
				if new_v == "" {
					new_v = u.(string)
					continue
				}
				new_v = new_v + "," + u.(string)
			}
			new_envs[k] = new_v
		}
	}
	return new_envs
}

func (s *commService) RandSeq(randlen int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, randlen)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (s *commService) GetScriptBody(script_name string) string {
	ctx := context.Background()
	script_body, err := dao.CicdScript.Ctx(ctx).
		Fields("script_body").Where("script_name=", script_name).Value()
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return script_body.String()
}

func stringToSlice(str string) []string {
	ctx := context.Background()

	var newslice []string = make([]string, 0)
	err := json.Unmarshal([]byte(str), &newslice)
	if err != nil {
		g.Log().Errorf(ctx, "stringToSlice: ", err)
	}
	return newslice
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (s *commService) CheckAuthor(ctx context.Context, pipeline_id int) bool {
	r := g.RequestFromCtx(ctx)
	var user_id int = r.Session.MustGet("userid").Int()

	group_id_user := User.GetGroupId(user_id)
	group_id_pipeline := Pipeline.GetGroupId(pipeline_id)

	return stringInSlice(fmt.Sprint(group_id_pipeline), group_id_user)
}
