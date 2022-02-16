package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/SKF/go-eventsource/v2/eventsource"
	"github.com/SKF/go-pas-client/internal/events"
	"github.com/SKF/go-utility/v2/uuid"
	pas "github.com/SKF/proto/v2/pas"
)

func Test_ThresholdEvent_FromInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		set      func(*testing.T, *events.SetPointAlarmThresholdEvent)
		expected *ThresholdEvent
	}{
		{
			set: func(t *testing.T, event *events.SetPointAlarmThresholdEvent) {
				var err error

				event.Overall, err = proto.Marshal(&pas.Overall{
					Unit: "C",
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
				})

				require.NoError(t, err)
			},
			expected: &ThresholdEvent{
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
				Threshold: Threshold{
					Overall: &Overall{
						Unit:      "C",
						OuterHigh: f64p(70),
						InnerHigh: f64p(50),
						InnerLow:  f64p(20),
						OuterLow:  f64p(10),
					},
					BandAlarms: []BandAlarm{},
					HALAlarms:  []HALAlarm{},
				},
			},
		},
		{
			set: func(t *testing.T, event *events.SetPointAlarmThresholdEvent) {
				var err error

				event.RateOfChange, err = proto.Marshal(&pas.RateOfChange{
					Unit: "gE",
					OuterHigh: &pas.DoubleObject{
						Value: 20,
					},
					InnerHigh: &pas.DoubleObject{
						Value: 10,
					},
					InnerLow: &pas.DoubleObject{
						Value: -10,
					},
					OuterLow: &pas.DoubleObject{
						Value: -20,
					},
				})

				require.NoError(t, err)
			},
			expected: &ThresholdEvent{
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
				Threshold: Threshold{
					RateOfChange: &RateOfChange{
						Unit:      "gE",
						OuterHigh: f64p(20),
						InnerHigh: f64p(10),
						InnerLow:  f64p(-10),
						OuterLow:  f64p(-20),
					},
					BandAlarms: []BandAlarm{},
					HALAlarms:  []HALAlarm{},
				},
			},
		},
		{
			set: func(t *testing.T, event *events.SetPointAlarmThresholdEvent) {
				var err error

				event.Inspection, err = proto.Marshal(&pas.Inspection{
					Choices: []*pas.InspectionChoice{
						{
							Answer:      "good",
							Instruction: "good?",
							Status:      pas.AlarmStatus_GOOD,
						},
					},
				})

				require.NoError(t, err)
			},
			expected: &ThresholdEvent{
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
				Threshold: Threshold{
					Inspection: &Inspection{
						Choices: []InspectionChoice{
							{
								Answer:      "good",
								Instruction: "good?",
								Status:      AlarmStatusGood,
							},
						},
					},
					BandAlarms: []BandAlarm{},
					HALAlarms:  []HALAlarm{},
				},
			},
		},
		{
			set: func(t *testing.T, event *events.SetPointAlarmThresholdEvent) {
				var err error

				event.BandAlarms = make([][]byte, 1)

				event.BandAlarms[0], err = proto.Marshal(&pas.BandAlarm{})

				require.NoError(t, err)
			},
			expected: &ThresholdEvent{
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
				Threshold: Threshold{
					BandAlarms: []BandAlarm{
						{},
					},
					HALAlarms: []HALAlarm{},
				},
			},
		},
		{
			set: func(t *testing.T, event *events.SetPointAlarmThresholdEvent) {
				var err error

				event.HalAlarms = make([][]byte, 1)

				event.HalAlarms[0], err = proto.Marshal(&pas.HalAlarm{})

				require.NoError(t, err)
			},
			expected: &ThresholdEvent{
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
				Threshold: Threshold{
					BandAlarms: []BandAlarm{},
					HALAlarms: []HALAlarm{
						{},
					},
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			given := &events.SetPointAlarmThresholdEvent{
				BaseEvent: &eventsource.BaseEvent{
					AggregateID: uuid.EmptyUUID.String(),
					UserID:      uuid.EmptyUUID.String(),
				},
			}

			test.set(t, given)

			buf, err := json.Marshal(given)
			require.NoError(t, err)

			actual := new(ThresholdEvent)

			err = actual.FromInternal(buf)
			require.NoError(t, err)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_ThresholdEvent_FromInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var event *ThresholdEvent

		err := event.FromInternal([]byte{})
		require.NoError(t, err)
	})
}

func Test_ThresholdEvent_FromInternal_InvalidBody(t *testing.T) {
	t.Parallel()

	event := new(ThresholdEvent)

	err := event.FromInternal([]byte(`not-valid`))

	assert.Error(t, err)
}

func Test_ThresholdEvent_FromInternal_InvalidEncodedThreshold(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given *events.SetPointAlarmThresholdEvent
	}{
		{
			given: &events.SetPointAlarmThresholdEvent{
				BaseEvent: &eventsource.BaseEvent{
					AggregateID: uuid.EmptyUUID.String(),
					UserID:      uuid.EmptyUUID.String(),
				},
				Overall: []byte("not-valid"),
			},
		},
		{
			given: &events.SetPointAlarmThresholdEvent{
				BaseEvent: &eventsource.BaseEvent{
					AggregateID: uuid.EmptyUUID.String(),
					UserID:      uuid.EmptyUUID.String(),
				},
				RateOfChange: []byte("not-valid"),
			},
		},
		{
			given: &events.SetPointAlarmThresholdEvent{
				BaseEvent: &eventsource.BaseEvent{
					AggregateID: uuid.EmptyUUID.String(),
					UserID:      uuid.EmptyUUID.String(),
				},
				Inspection: []byte("not-valid"),
			},
		},
		{
			given: &events.SetPointAlarmThresholdEvent{
				BaseEvent: &eventsource.BaseEvent{
					AggregateID: uuid.EmptyUUID.String(),
					UserID:      uuid.EmptyUUID.String(),
				},
				BandAlarms: [][]byte{
					[]byte("not-valid"),
				},
			},
		},
		{
			given: &events.SetPointAlarmThresholdEvent{
				BaseEvent: &eventsource.BaseEvent{
					AggregateID: uuid.EmptyUUID.String(),
					UserID:      uuid.EmptyUUID.String(),
				},
				HalAlarms: [][]byte{
					[]byte("not-valid"),
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			buf, err := json.Marshal(test.given)
			require.NoError(t, err)

			event := new(ThresholdEvent)

			err = event.FromInternal(buf)

			assert.Error(t, err)
		})
	}
}

func Test_AlarmStatusEvent_FromInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    []byte
		expected *AlarmStatusEvent
	}{
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"alarmStatus": 0
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					Status: AlarmStatusNotConfigured,
					Band:   []BandAlarmStatus{},
					HAL:    []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"alarmStatus": 1
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					Status: AlarmStatusNoData,
					Band:   []BandAlarmStatus{},
					HAL:    []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"alarmStatus": 2
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					Status: AlarmStatusGood,
					Band:   []BandAlarmStatus{},
					HAL:    []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"alarmStatus": 3
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					Status: AlarmStatusAlert,
					Band:   []BandAlarmStatus{},
					HAL:    []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"alarmStatus": 4
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					Status: AlarmStatusDanger,
					Band:   []BandAlarmStatus{},
					HAL:    []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"overallAlarm": {
		"status": 2,
		"triggeringMeasurement": "00000000-0000-0000-0000-000000000000"
	}
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					Overall: &GenericAlarmStatus{
						Status:                AlarmStatusGood,
						TriggeringMeasurement: uuid.EmptyUUID,
					},
					Band: []BandAlarmStatus{},
					HAL:  []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"rateOfChangeAlarm": {
		"status": 2,
		"triggeringMeasurement": "00000000-0000-0000-0000-000000000000"
	}
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					RateOfChange: &GenericAlarmStatus{
						Status:                AlarmStatusGood,
						TriggeringMeasurement: uuid.EmptyUUID,
					},
					Band: []BandAlarmStatus{},
					HAL:  []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"inspectionAlarm": {
		"status": 2,
		"triggeringMeasurement": "00000000-0000-0000-0000-000000000000"
	}
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					Inspection: &GenericAlarmStatus{
						Status:                AlarmStatusGood,
						TriggeringMeasurement: uuid.EmptyUUID,
					},
					Band: []BandAlarmStatus{},
					HAL:  []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"externalAlarm": {
		"status": 2
	}
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					External: &ExternalAlarmStatus{
						Status: AlarmStatusGood,
					},
					Band: []BandAlarmStatus{},
					HAL:  []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"bandAlarms": [
		{
			"label": "10x RPM",
			"status": 2,
			"triggeringMeasurement": "00000000-0000-0000-0000-000000000000",
			"minFrequency": {
				"valueType": 1,
				"value": 100
			},
			"maxFrequency": {
				"valueType": 1,
				"value": 500
			},
			"calculatedOverall": {
				"unit": "gE",
				"value": 0.5
			}
		}
	]
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					Band: []BandAlarmStatus{
						{
							Label: "10x RPM",
							GenericAlarmStatus: GenericAlarmStatus{
								Status:                AlarmStatusGood,
								TriggeringMeasurement: uuid.EmptyUUID,
							},
							MinFrequency: BandAlarmFrequency{
								ValueType: BandAlarmFrequencyFixed,
								Value:     100,
							},
							MaxFrequency: BandAlarmFrequency{
								ValueType: BandAlarmFrequencyFixed,
								Value:     500,
							},
							CalculatedOverall: &BandAlarmStatusCalculatedOverall{
								Unit:  "gE",
								Value: 0.5,
							},
						},
					},
					HAL: []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"halAlarms": [
		{
			"label": "10x RPM",
			"status": 2,
			"triggeringMeasurement": "00000000-0000-0000-0000-000000000000",
			"halIndex": 1.22,
			"faultFrequency": 12.22,
			"rpmFactor": 10,
			"numberOfHarmonicsUsed": 5
		}
	]
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					Band: []BandAlarmStatus{},
					HAL: []HALAlarmStatus{
						{
							Label: "10x RPM",
							GenericAlarmStatus: GenericAlarmStatus{
								Status:                AlarmStatusGood,
								TriggeringMeasurement: uuid.EmptyUUID,
							},
							HALIndex:              f64p(1.22),
							FaultFrequency:        f64p(12.22),
							RPMFactor:             f64p(10),
							NumberOfHarmonicsUsed: i64p(5),
						},
					},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := new(AlarmStatusEvent)

			err := actual.FromInternal(test.given)
			require.NoError(t, err)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_AlarmStatusEvent_FromInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var event *AlarmStatusEvent

		event.FromInternal([]byte{})
	})
}

func Test_AlarmStatusEvent_FromInternal_InvalidBody(t *testing.T) {
	t.Parallel()

	event := new(AlarmStatusEvent)

	err := event.FromInternal([]byte(`not-valid`))

	assert.Error(t, err)
}
