package poller

type PollResult struct {
	httpErrorCode int
}

type PollTask struct {
	url string
}

type Poller interface {
	Poll(task PollTask) PollResult
}
