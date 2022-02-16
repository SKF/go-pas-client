package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	models "github.com/SKF/go-pas-client/internal/models"
	pas "github.com/SKF/proto/v2/pas"
)

func Test_RateOfChange_FromInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *models.ModelsRateOfChange
		expected *RateOfChange
	}{
		{
			given:    nil,
			expected: &RateOfChange{},
		},
		{
			given:    &models.ModelsRateOfChange{},
			expected: &RateOfChange{},
		},
		{
			given: &models.ModelsRateOfChange{
				Unit: "gE",
			},
			expected: &RateOfChange{
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
			expected: &RateOfChange{
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
			actual := new(RateOfChange)

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
		rateOfChange := new(RateOfChange)

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
	t.Parallel()

	given := &RateOfChange{
		Unit:      "gE",
		OuterHigh: f64p(70),
		InnerHigh: f64p(50),
		InnerLow:  f64p(20),
		OuterLow:  f64p(10),
	}

	actual := new(RateOfChange)

	actual.FromInternal(given.ToInternal())

	assert.Equal(t, given, actual)
}

func Test_RateOfChange_FromProto(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *pas.RateOfChange
		expected *RateOfChange
	}{
		{
			given: &pas.RateOfChange{
				Unit: "gE",
			},
			expected: &RateOfChange{
				Unit: "gE",
			},
		},
		{
			given: &pas.RateOfChange{
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
			expected: &RateOfChange{
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

			actual := new(RateOfChange)

			err = actual.FromProto(buf)
			require.NoError(t, err)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_RateOfChange_FromEvent_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var rateOfChange *RateOfChange

		err := rateOfChange.FromProto([]byte{})

		assert.NoError(t, err)
	})

	assert.NotPanics(t, func() {
		rateOfChange := new(RateOfChange)

		err := rateOfChange.FromProto(nil)

		assert.NoError(t, err)
	})
}

func Test_RateOfChange_FromEvent_InvalidBody(t *testing.T) {
	t.Parallel()

	rateOfChange := new(RateOfChange)

	err := rateOfChange.FromProto([]byte("not-valid"))

	assert.Error(t, err)
}
