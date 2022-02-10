package models

import (
	"testing"

	"github.com/stretchr/testify/assert"

	models "github.com/SKF/go-pas-client/internal/models"
)

func Test_RateOfChange_FromInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *models.ModelsRateOfChange
		expected RateOfChange
	}{
		{
			given:    nil,
			expected: RateOfChange{},
		},
		{
			given:    &models.ModelsRateOfChange{},
			expected: RateOfChange{},
		},
		{
			given: &models.ModelsRateOfChange{
				Unit: "gE",
			},
			expected: RateOfChange{
				Unit: "gE",
			},
		},
		{
			given: &models.ModelsRateOfChange{
				Unit:      "gE",
				OuterHigh: f64p(70),
				InnerHigh: f64p(50),
				InnerLow:  f64p(20),
				OuterLow:  f64p(10),
			},
			expected: RateOfChange{
				Unit:      "gE",
				OuterHigh: f64p(70),
				InnerHigh: f64p(50),
				InnerLow:  f64p(20),
				OuterLow:  f64p(10),
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := RateOfChange{}

			actual.FromInternal(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}
