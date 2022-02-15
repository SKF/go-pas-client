package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	models "github.com/SKF/go-pas-client/internal/models"
	pas "github.com/SKF/proto/v2/pas"
)

func Test_Inspection_FromInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *models.ModelsInspection
		expected *Inspection
	}{
		{
			given:    nil,
			expected: &Inspection{},
		},
		{
			given: &models.ModelsInspection{},
			expected: &Inspection{
				Choices: []InspectionChoice{},
			},
		},
		{
			given: &models.ModelsInspection{
				Choices: []*models.ModelsInspectionChoice{
					{
						Answer:      "missing status",
						Instruction: "is missing status?",
					},
				},
			},
			expected: &Inspection{
				Choices: []InspectionChoice{
					{
						Answer:      "missing status",
						Instruction: "is missing status?",
						Status:      AlarmStatusNotConfigured,
					},
				},
			},
		},
		{
			given: &models.ModelsInspection{
				Choices: []*models.ModelsInspectionChoice{
					{
						Answer:      "good",
						Instruction: "is good?",
						Status:      i32p(2), // good
					},
				},
			},
			expected: &Inspection{
				Choices: []InspectionChoice{
					{
						Answer:      "good",
						Instruction: "is good?",
						Status:      AlarmStatusGood,
					},
				},
			},
		},
		{
			given: &models.ModelsInspection{
				Choices: []*models.ModelsInspectionChoice{
					{
						Answer:      "not configured",
						Instruction: "is not configured?",
						Status:      i32p(0), // not_configured
					},
					{
						Answer:      "no data",
						Instruction: "is no data?",
						Status:      i32p(1), // no_data
					},
					{
						Answer:      "good",
						Instruction: "is good?",
						Status:      i32p(2), // good
					},
					{
						Answer:      "alert",
						Instruction: "is alert?",
						Status:      i32p(3), // alert
					},
					{
						Answer:      "danger",
						Instruction: "is danger?",
						Status:      i32p(4), // danger
					},
				},
			},
			expected: &Inspection{
				Choices: []InspectionChoice{
					{
						Answer:      "not configured",
						Instruction: "is not configured?",
						Status:      AlarmStatusNotConfigured,
					},
					{
						Answer:      "no data",
						Instruction: "is no data?",
						Status:      AlarmStatusNoData,
					},
					{
						Answer:      "good",
						Instruction: "is good?",
						Status:      AlarmStatusGood,
					},
					{
						Answer:      "alert",
						Instruction: "is alert?",
						Status:      AlarmStatusAlert,
					},
					{
						Answer:      "danger",
						Instruction: "is danger?",
						Status:      AlarmStatusDanger,
					},
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := new(Inspection)

			actual.FromInternal(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_InspectionChoice_Nil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		inspection := new(InspectionChoice)

		inspection.FromInternal(nil)

		assert.Equal(t, &InspectionChoice{}, inspection)
	})
}

func Test_Inspection_ToInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *Inspection
		expected *models.ModelsInspection
	}{
		{
			given:    nil,
			expected: nil,
		},
		{
			given: &Inspection{
				Choices: []InspectionChoice{},
			},
			expected: &models.ModelsInspection{
				Choices: []*models.ModelsInspectionChoice{},
			},
		},
		{
			given: &Inspection{
				Choices: []InspectionChoice{
					{
						Answer:      "good",
						Instruction: "is good?",
						Status:      AlarmStatusGood,
					},
				},
			},
			expected: &models.ModelsInspection{
				Choices: []*models.ModelsInspectionChoice{
					{
						Answer:      "good",
						Instruction: "is good?",
						Status:      i32p(2), // good
					},
				},
			},
		},
		{
			given: &Inspection{
				Choices: []InspectionChoice{
					{
						Answer:      "not configured",
						Instruction: "is not configured?",
						Status:      AlarmStatusNotConfigured,
					},
					{
						Answer:      "no data",
						Instruction: "is no data?",
						Status:      AlarmStatusNoData,
					},
					{
						Answer:      "good",
						Instruction: "is good?",
						Status:      AlarmStatusGood,
					},
					{
						Answer:      "alert",
						Instruction: "is alert?",
						Status:      AlarmStatusAlert,
					},
					{
						Answer:      "danger",
						Instruction: "is danger?",
						Status:      AlarmStatusDanger,
					},
				},
			},
			expected: &models.ModelsInspection{
				Choices: []*models.ModelsInspectionChoice{
					{
						Answer:      "not configured",
						Instruction: "is not configured?",
						Status:      i32p(0), // not_configured
					},
					{
						Answer:      "no data",
						Instruction: "is no data?",
						Status:      i32p(1), // no_data
					},
					{
						Answer:      "good",
						Instruction: "is good?",
						Status:      i32p(2), // good
					},
					{
						Answer:      "alert",
						Instruction: "is alert?",
						Status:      i32p(3), // alert
					},
					{
						Answer:      "danger",
						Instruction: "is danger?",
						Status:      i32p(4), // danger
					},
				},
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

func Test_InspectionChoice_ToInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var choice *InspectionChoice

		actual := choice.ToInternal()

		assert.Nil(t, actual)
	})
}

func Test_InspectionChoice_FromProto(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *pas.Inspection
		expected *Inspection
	}{
		{
			given: &pas.Inspection{
				Choices: []*pas.InspectionChoice{
					{
						Answer:      "missing status",
						Instruction: "is missing status?",
					},
				},
			},
			expected: &Inspection{
				Choices: []InspectionChoice{
					{
						Answer:      "missing status",
						Instruction: "is missing status?",
						Status:      AlarmStatusNotConfigured,
					},
				},
			},
		},
		{
			given: &pas.Inspection{
				Choices: []*pas.InspectionChoice{
					{
						Answer:      "good",
						Instruction: "is good?",
						Status:      pas.AlarmStatus_GOOD,
					},
				},
			},
			expected: &Inspection{
				Choices: []InspectionChoice{
					{
						Answer:      "good",
						Instruction: "is good?",
						Status:      AlarmStatusGood,
					},
				},
			},
		},
		{
			given: &pas.Inspection{
				Choices: []*pas.InspectionChoice{
					{
						Answer:      "not configured",
						Instruction: "is not configured?",
						Status:      pas.AlarmStatus_NOT_CONFIGURED,
					},
					{
						Answer:      "no data",
						Instruction: "is no data?",
						Status:      pas.AlarmStatus_NO_DATA,
					},
					{
						Answer:      "good",
						Instruction: "is good?",
						Status:      pas.AlarmStatus_GOOD,
					},
					{
						Answer:      "alert",
						Instruction: "is alert?",
						Status:      pas.AlarmStatus_ALERT,
					},
					{
						Answer:      "danger",
						Instruction: "is danger?",
						Status:      pas.AlarmStatus_DANGER,
					},
				},
			},
			expected: &Inspection{
				Choices: []InspectionChoice{
					{
						Answer:      "not configured",
						Instruction: "is not configured?",
						Status:      AlarmStatusNotConfigured,
					},
					{
						Answer:      "no data",
						Instruction: "is no data?",
						Status:      AlarmStatusNoData,
					},
					{
						Answer:      "good",
						Instruction: "is good?",
						Status:      AlarmStatusGood,
					},
					{
						Answer:      "alert",
						Instruction: "is alert?",
						Status:      AlarmStatusAlert,
					},
					{
						Answer:      "danger",
						Instruction: "is danger?",
						Status:      AlarmStatusDanger,
					},
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			buf, err := proto.Marshal(test.given)
			require.NoError(t, err)

			actual := new(Inspection)

			err = actual.FromProto(buf)
			require.NoError(t, err)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_InspectionChoice_FromProto_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var choice *InspectionChoice

		choice.FromProto(&pas.InspectionChoice{})
	})

	assert.NotPanics(t, func() {
		choice := new(InspectionChoice)

		choice.FromProto(nil)
	})
}

func Test_InspectionChoice_FromProto_InvalidBody(t *testing.T) {
	t.Parallel()

	inspection := new(Inspection)

	err := inspection.FromProto([]byte("not-valid"))

	assert.Error(t, err)
}
