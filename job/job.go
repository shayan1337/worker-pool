package job

type Job struct {
	requestId int
	url string
}

func NewJob(requestId int, url string) Job {
	return Job{
		requestId: requestId,
		url:       url,
	}
}

func (this Job) Url() string {
	return this.url
}

func (this Job) RequestId() int {
	return this.requestId
}


