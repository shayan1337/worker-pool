package result

import (
	"net/http"
	. "worker-pool/job"
)

type Result struct {
	job Job
	response *http.Response
	err error
}

func NewResult(job Job, response *http.Response, err error) Result {
	return Result{
		job: job,
		response: response,
		err: err,
	}
}

func (this Result) Err() error {
	return this.err
}

func (this Result) Response() http.Response {
	return this.response
}

func (this Result) Jon() Job {
	return this.job
}
