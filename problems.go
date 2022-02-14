package client

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/SKF/go-rest-utility/problems"
)

type ProblemDecoder struct{}

func (p *ProblemDecoder) DecodeProblem(ctx context.Context, r *http.Response) (problems.Problem, error) {
	decoder := json.NewDecoder(r.Body)

	switch r.StatusCode {
	case http.StatusBadRequest, http.StatusConflict:
		var (
			problem = problems.ValidationProblem{}
			err     = decoder.Decode(&problem)
		)

		return problem, err
	case http.StatusInternalServerError:
		var (
			problem = problems.InternalProblem{}
			err     = decoder.Decode(&problem)
		)

		return problem, err
	default:
		var (
			problem = problems.BasicProblem{}
			err     = decoder.Decode(&problem)
		)

		return problem, err
	}
}
