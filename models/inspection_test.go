package models

import (
	"testing"

	"github.com/stretchr/testify/assert"

	models "github.com/SKF/go-pas-client/internal/models"
)

func Test_Inspection_FromInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *models.ModelsInspection
		expected Inspection
	}{
		{
			given:    nil,
			expected: Inspection{},
		},
		{
			given: &models.ModelsInspection{},
			expected: Inspection{
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
			expected: Inspection{
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
			expected: Inspection{
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
			expected: Inspection{
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
			actual := Inspection{}

			actual.FromInternal(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_InspectionChoice_Nil(t *testing.T) {
	inspection := InspectionChoice{}

	inspection.FromInternal(nil)

	assert.Equal(t, InspectionChoice{}, inspection)
}
