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

func Test_Overall_FromInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var overall *Overall

		overall.FromInternal(&models.ModelsOverall{})
	})

	assert.NotPanics(t, func() {
		overall := &Overall{}

		overall.FromInternal(nil)
	})
}

func Test_Overall_ToInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *Overall
		expected *models.ModelsOverall
	}{
		{
			given:    nil,
			expected: nil,
		},
		{
			given:    &Overall{},
			expected: &models.ModelsOverall{},
		},
		{
			given: &Overall{
				Unit: "gE",
			},
			expected: &models.ModelsOverall{
				Unit: "gE",
			},
		},
		{
			given: &Overall{
				Unit:      "gE",
				OuterHigh: f64p(70),
				InnerHigh: f64p(50),
				InnerLow:  f64p(20),
				OuterLow:  f64p(10),
			},
			expected: &models.ModelsOverall{
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

func Test_Overall_ToInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var overall *Overall

		actual := overall.ToInternal()

		assert.Nil(t, actual)
	})
}

func Test_Overall_Convert(t *testing.T) {
	given := &Overall{
		Unit:      "gE",
		OuterHigh: f64p(70),
		InnerHigh: f64p(50),
		InnerLow:  f64p(20),
		OuterLow:  f64p(10),
	}

	actual := &Overall{}

	actual.FromInternal(given.ToInternal())

	assert.Equal(t, given, actual)
}
