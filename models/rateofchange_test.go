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

func Test_RateOfChange_FromInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var rateOfChange *RateOfChange

		rateOfChange.FromInternal(&models.ModelsRateOfChange{})
	})

	assert.NotPanics(t, func() {
		rateOfChange := &RateOfChange{}

		rateOfChange.FromInternal(nil)
	})
}

func Test_RateOfChange_ToInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *RateOfChange
		expected *models.ModelsRateOfChange
	}{
		{
			given:    nil,
			expected: nil,
		},
		{
			given:    &RateOfChange{},
			expected: &models.ModelsRateOfChange{},
		},
		{
			given: &RateOfChange{
				Unit: "gE",
			},
			expected: &models.ModelsRateOfChange{
				Unit: "gE",
			},
		},
		{
			given: &RateOfChange{
				Unit:      "gE",
				OuterHigh: f64p(70),
				InnerHigh: f64p(50),
				InnerLow:  f64p(20),
				OuterLow:  f64p(10),
			},
			expected: &models.ModelsRateOfChange{
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
			actual := test.given.ToInternal()

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_RateOfChange_ToInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var rateOfChange *RateOfChange

		actual := rateOfChange.ToInternal()

		assert.Nil(t, actual)
	})
}

func Test_RateOfChange_Convert(t *testing.T) {
	given := &RateOfChange{
		Unit:      "gE",
		OuterHigh: f64p(70),
		InnerHigh: f64p(50),
		InnerLow:  f64p(20),
		OuterLow:  f64p(10),
	}

	actual := &RateOfChange{}

	actual.FromInternal(given.ToInternal())

	assert.Equal(t, given, actual)
}
