// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// CicdJob is the golang structure for table cicd_job.
type CicdJob struct {
	Id         uint64 `json:"id"         orm:"id"          description:""` //
	PipelineId uint   `json:"pipelineId" orm:"pipeline_id" description:""` //
	AgentId    uint   `json:"agentId"    orm:"agent_id"    description:""` //
	JobType    string `json:"jobType"    orm:"job_type"    description:""` //
	JobStatus  string `json:"jobStatus"  orm:"job_status"  description:""` //
	Script     string `json:"script"     orm:"script"      description:""` //
	Comment    string `json:"comment"    orm:"comment"     description:""` //
	Author     string `json:"author"     orm:"author"      description:""` //
	CreatedAt  uint64 `json:"createdAt"  orm:"created_at"  description:""` //
}
