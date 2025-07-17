// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// CicdPipeline is the golang structure for table cicd_pipeline.
type CicdPipeline struct {
	Id           uint   `json:"id"           orm:"id"            description:""` //
	PipelineName string `json:"pipelineName" orm:"pipeline_name" description:""` //
	GroupId      uint   `json:"groupId"      orm:"group_id"      description:""` //
	AgentId      uint   `json:"agentId"      orm:"agent_id"      description:""` //
	Concurrency  uint   `json:"concurrency"  orm:"concurrency"   description:""` //
	PipelineBody string `json:"pipelineBody" orm:"pipeline_body" description:""` //
	Author       string `json:"author"       orm:"author"        description:""` //
	UpdatedAt    uint64 `json:"updatedAt"    orm:"updated_at"    description:""` //
}
