package models

import (
	"testing"

	"github.com/stretchr/testify/assert"

	models "github.com/SKF/go-pas-client/internal/models"
)

func Test_Overall_FromInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *models.ModelsOverall
		expected Overall
	}{
		{
			given:    nil,
			expected: Overall{},
		},
		{
			given:    &models.ModelsOverall{},
			expected: Overall{},
		},
		{
			given: &models.ModelsOverall{
				Unit: "gE",
			},
			expected: Overall{
				Unit: "gE",
			},
		},
		{
			given: &models.ModelsOverall{
				Unit:      "gE",
				OuterHigh: f64p(70),
				InnerHigh: f64p(50),
				InnerLow:  f64p(20),
				OuterLow:  f64p(10),
			},
			expected: Overall{
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
			actual := Overall{}

			actual.FromInternal(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}
