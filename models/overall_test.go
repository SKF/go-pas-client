package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	models "github.com/SKF/go-pas-client/internal/models"
	pas "github.com/SKF/proto/v2/pas"
)

func Test_Overall_FromInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *models.ModelsOverall
		expected *Overall
	}{
		{
			given:    nil,
			expected: &Overall{},
		},
		{
			given:    &models.ModelsOverall{},
			expected: &Overall{},
		},
		{
			given: &models.ModelsOverall{
				Unit: "gE",
			},
			expected: &Overall{
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
			expected: &Overall{
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
			actual := new(Overall)

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
		overall := new(Overall)

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

func Test_Overall_Convert(t *testing.T) {
	t.Parallel()

	given := &Overall{
		Unit:      "gE",
		OuterHigh: f64p(70),
		InnerHigh: f64p(50),
		InnerLow:  f64p(20),
		OuterLow:  f64p(10),
	}

	actual := new(Overall)

	actual.FromInternal(given.ToInternal())

	assert.Equal(t, given, actual)
}

func Test_Overall_FromEvent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *pas.Overall
		expected *Overall
	}{
		{
			given:    &pas.Overall{},
			expected: &Overall{},
		},
		{
			given: &pas.Overall{
				Unit: "gE",
			},
			expected: &Overall{
				Unit: "gE",
			},
		},
		{
			given: &pas.Overall{
				Unit: "gE",
				OuterHigh: &pas.DoubleObject{
					Value: 70,
				},
				InnerHigh: &pas.DoubleObject{
					Value: 50,
				},
				InnerLow: &pas.DoubleObject{
					Value: 20,
				},
				OuterLow: &pas.DoubleObject{
					Value: 10,
				},
			},
			expected: &Overall{
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
			buf, err := proto.Marshal(test.given)
			require.NoError(t, err)

			actual := new(Overall)

			err = actual.FromProto(buf)
			require.NoError(t, err)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_Overall_FromEvent_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var overall *Overall

		err := overall.FromProto([]byte{})

		assert.NoError(t, err)
	})

	assert.NotPanics(t, func() {
		overall := new(Overall)

		err := overall.FromProto(nil)

		assert.NoError(t, err)
	})
}

func Test_Overall_FromEvent_InvalidBody(t *testing.T) {
	t.Parallel()

	overall := new(Overall)

	err := overall.FromProto([]byte("not-valid"))

	assert.Error(t, err)
}
