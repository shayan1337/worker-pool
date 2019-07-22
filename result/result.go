package result

import (
	"net/http"
	"worker-pool/job"
)

type Result struct {
	job job.Job
	response http.Response
}

func (this Result) Response() http.Response {
	return this.response
}

func (this Result) Jon() job.Job {
	return this.job
}
