// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// CicdLog is the golang structure for table cicd_log.
type CicdLog struct {
	Id         uint64 `json:"id"         orm:"id"          description:""` //
	PipelineId uint   `json:"pipelineId" orm:"pipeline_id" description:""` //
	AgentId    uint   `json:"agentId"    orm:"agent_id"    description:""` //
	JobType    string `json:"jobType"    orm:"job_type"    description:""` //
	JobId      uint   `json:"jobId"      orm:"job_id"      description:""` //
	TaskStatus string `json:"taskStatus" orm:"task_status" description:""` //
	Ipaddr     string `json:"ipaddr"     orm:"ipaddr"      description:""` //
	UpdatedAt  uint64 `json:"updatedAt"  orm:"updated_at"  description:""` //
	Output     string `json:"output"     orm:"output"      description:""` //
}
