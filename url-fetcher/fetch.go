package url_fetcher

import (
	"net/http"
	. "worker-pool/job"
	. "worker-pool/result"
)

func Fetch(job Job) Result {
	response, err := http.Get(job.Url())
	return NewResult(job, response, err)
}
