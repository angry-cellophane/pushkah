package pusher

import (
	"github.com/gogo/protobuf/proto"
	"github.com/prometheus/prometheus/prompb"
)

type PushRequest struct {
	metrics prompb.WriteRequest
}

type Pusher interface {
	Push(request PushRequest) error
}

func NewPusher() Pusher {
	return &HttpPusher{}
}

type HttpPusher struct {
	remoteUrl string
}

func (pusher *HttpPusher) Push(request PushRequest) error {
	bytes, err := proto.Marshal(&request.metrics)
	if err != nil {
		return err
	}
	return nil
}
