package main

import (
	"context"

	dd_http "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"

	pas "github.com/SKF/go-pas-client"
	"github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-rest-utility/client/auth"
	"github.com/SKF/go-utility/v2/stages"
)

const serviceName = "example-service"

type tokenProvider struct{}

func (t *tokenProvider) GetRawToken(ctx context.Context) (auth.RawToken, error) {
	return auth.RawToken(""), nil
}

func main() {
	client := pas.New(
		pas.WithStage(stages.StageSandbox),
		client.WithDatadogTracing(dd_http.RTWithServiceName(serviceName)),
		client.WithTokenProvider(&tokenProvider{}),
	)

	_ = client
}
