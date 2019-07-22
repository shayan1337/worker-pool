package worker_pool

type Job struct {
	requestId int
	url string
}

func (this Job) Url() string {
	return this.url
}

func (this Job) RequestId() int {
	return this.requestId
}


