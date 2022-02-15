package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ProblemDecoder(t *testing.T) {
	tests := []struct {
		given *http.Response
	}{
		{
			given: &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(`{}`))),
			},
		},
		{
			given: &http.Response{
				StatusCode: http.StatusConflict,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(`{}`))),
			},
		},
		{
			given: &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(`{}`))),
			},
		},
		{
			given: &http.Response{
				StatusCode: http.StatusUnauthorized,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(`{}`))),
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			p := &ProblemDecoder{}

			_, err := p.DecodeProblem(context.TODO(), test.given)

			assert.NoError(t, err)
		})
	}
}
