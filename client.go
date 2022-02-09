package client

import (
	"fmt"

	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-utility/v2/stages"
)

type Client interface{}

type client struct {
	*rest.Client
}

func WithStage(stage string) rest.Option {
	if stage == stages.StageProd {
		return rest.WithBaseURL("https://api.point-alarm-status.sandbox.iot.enlight.skf.com/v1")
	}

	return rest.WithBaseURL(fmt.Sprintf("https://api.point-alarm-status.%s.iot.enlight.skf.com/v1", stage))
}

func New(opts ...rest.Option) Client {
	restClient := rest.NewClient(
		append([]rest.Option{
			// Defaults to production stage if no option is supplied
			WithStage(stages.StageProd),
		}, opts...)...,
	)

	return &client{Client: restClient}
}
