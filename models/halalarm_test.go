package models

import (
	"math"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/SKF/go-pas-client/internal/events"
	models "github.com/SKF/go-pas-client/internal/models"
	"github.com/SKF/go-utility/v2/uuid"
	pas "github.com/SKF/proto/v2/pas"
)

func Test_HALAlarm_FromInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *models.ModelsHALAlarm
		expected *HALAlarm
	}{
		{
			given:    nil,
			expected: &HALAlarm{},
		},
		{
			given:    &models.ModelsHALAlarm{},
			expected: &HALAlarm{},
		},
		{
			given: &models.ModelsHALAlarm{
				HalAlarmType: models.ModelsHALAlarmHalAlarmTypeGLOBAL,
			},
			expected: &HALAlarm{
				HALAlarmType: HALAlarmTypeGlobal,
			},
		},
		{
			given: &models.ModelsHALAlarm{
				HalAlarmType: models.ModelsHALAlarmHalAlarmTypeFREQUENCY,
			},
			expected: &HALAlarm{
				HALAlarmType: HALAlarmTypeFaultFrequency,
			},
		},
		{
			given: &models.ModelsHALAlarm{
				Bearing: &models.ModelsBearing{
					Manufacturer: stringp("SKF"),
					ModelNumber:  stringp("2222"),
				},
			},
			expected: &HALAlarm{
				Bearing: &Bearing{
					Manufacturer: "SKF",
					ModelNumber:  "2222",
				},
			},
		},
		{
			given: &models.ModelsHALAlarm{
				UpperAlert: f64p(10),
			},
			expected: &HALAlarm{
				UpperAlert: f64p(10),
			},
		},
		{
			given: &models.ModelsHALAlarm{
				UpperDanger: f64p(20),
			},
			expected: &HALAlarm{
				UpperDanger: f64p(20),
			},
		},
		{
			given: &models.ModelsHALAlarm{
				UpperAlert:  f64p(10),
				UpperDanger: f64p(20),
			},
			expected: &HALAlarm{
				UpperAlert:  f64p(10),
				UpperDanger: f64p(20),
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := new(HALAlarm)

			actual.FromInternal(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_HALAlarm_FromInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var halAlarm *HALAlarm

		halAlarm.FromInternal(&models.ModelsHALAlarm{})
	})

	assert.NotPanics(t, func() {
		halAlarm := new(HALAlarm)

		halAlarm.FromInternal(nil)
	})
}

func Test_HALAlarm_ToInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *HALAlarm
		expected *models.ModelsHALAlarm
	}{
		{
			given:    nil,
			expected: nil,
		},
		{
			given:    &HALAlarm{},
			expected: &models.ModelsHALAlarm{},
		},
		{
			given: &HALAlarm{
				HALAlarmType: HALAlarmTypeGlobal,
			},
			expected: &models.ModelsHALAlarm{
				HalAlarmType: models.ModelsHALAlarmHalAlarmTypeGLOBAL,
			},
		},
		{
			given: &HALAlarm{
				HALAlarmType: HALAlarmTypeFaultFrequency,
			},
			expected: &models.ModelsHALAlarm{
				HalAlarmType: models.ModelsHALAlarmHalAlarmTypeFREQUENCY,
			},
		},
		{
			given: &HALAlarm{
				Bearing: &Bearing{
					Manufacturer: "SKF",
					ModelNumber:  "2222",
				},
			},
			expected: &models.ModelsHALAlarm{
				Bearing: &models.ModelsBearing{
					Manufacturer: stringp("SKF"),
					ModelNumber:  stringp("2222"),
				},
			},
		},
		{
			given: &HALAlarm{
				UpperAlert: f64p(10),
			},
			expected: &models.ModelsHALAlarm{
				UpperAlert: f64p(10),
			},
		},
		{
			given: &HALAlarm{
				UpperDanger: f64p(20),
			},
			expected: &models.ModelsHALAlarm{
				UpperDanger: f64p(20),
			},
		},
		{
			given: &HALAlarm{
				UpperAlert:  f64p(10),
				UpperDanger: f64p(20),
			},
			expected: &models.ModelsHALAlarm{
				UpperAlert:  f64p(10),
				UpperDanger: f64p(20),
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

func Test_HALAlarm_ToInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var halAlarm *HALAlarm

		actual := halAlarm.ToInternal()

		assert.Nil(t, actual)
	})
}

func Test_HALAlarm_FromProto(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *pas.HalAlarm
		expected *HALAlarm
	}{
		{
			given: &pas.HalAlarm{
				HalAlarmType: models.ModelsHALAlarmHalAlarmTypeGLOBAL,
			},
			expected: &HALAlarm{
				HALAlarmType: HALAlarmTypeGlobal,
			},
		},
		{
			given: &pas.HalAlarm{
				HalAlarmType: models.ModelsHALAlarmHalAlarmTypeFREQUENCY,
			},
			expected: &HALAlarm{
				HALAlarmType: HALAlarmTypeFaultFrequency,
			},
		},
		{
			given: &pas.HalAlarm{
				Bearing: &pas.Bearing{
					Manufacturer: "SKF",
					ModelNumber:  "2222",
				},
			},
			expected: &HALAlarm{
				Bearing: &Bearing{
					Manufacturer: "SKF",
					ModelNumber:  "2222",
				},
			},
		},
		{
			given: &pas.HalAlarm{
				UpperAlert: &pas.DoubleObject{
					Value: 10,
				},
			},
			expected: &HALAlarm{
				UpperAlert: f64p(10),
			},
		},
		{
			given: &pas.HalAlarm{
				UpperDanger: &pas.DoubleObject{
					Value: 20,
				},
			},
			expected: &HALAlarm{
				UpperDanger: f64p(20),
			},
		},
		{
			given: &pas.HalAlarm{
				UpperAlert: &pas.DoubleObject{
					Value: 10,
				},
				UpperDanger: &pas.DoubleObject{
					Value: 20,
				},
			},
			expected: &HALAlarm{
				UpperAlert:  f64p(10),
				UpperDanger: f64p(20),
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			buf, err := proto.Marshal(test.given)
			require.NoError(t, err)

			actual := new(HALAlarm)

			err = actual.FromProto(buf)
			require.NoError(t, err)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_HALAlarm_FromProto_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var halAlarm *HALAlarm

		err := halAlarm.FromProto([]byte{})

		assert.NoError(t, err)
	})

	assert.NotPanics(t, func() {
		halAlarm := new(HALAlarm)

		err := halAlarm.FromProto(nil)

		assert.NoError(t, err)
	})
}

func Test_HALAlarm_FromProto_InvalidBody(t *testing.T) {
	t.Parallel()

	actual := new(HALAlarm)

	err := actual.FromProto([]byte("not-valid"))

	assert.Error(t, err)
}

func Test_HALAlarmStatus_FromInternal(t *testing.T) {
	t.Parallel()

	triggeringMeasurement := strfmt.UUID(uuid.EmptyUUID.String())

	tests := []struct {
		given    *models.ModelsGetAlarmStatusResponseHALAlarm
		expected *HALAlarmStatus
	}{
		{
			given:    &models.ModelsGetAlarmStatusResponseHALAlarm{},
			expected: &HALAlarmStatus{},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				Label: stringp("10x RPM"),
			},
			expected: &HALAlarmStatus{
				Label: "10x RPM",
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				Status: i32p(0), // not configured
			},
			expected: &HALAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusNotConfigured,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				Status: i32p(1), // no data
			},
			expected: &HALAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusNoData,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				Status: i32p(2), // good
			},
			expected: &HALAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusGood,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				Status: i32p(3), // alert
			},
			expected: &HALAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusAlert,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				Status: i32p(4), // danger
			},
			expected: &HALAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusDanger,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				TriggeringMeasurement: &triggeringMeasurement,
			},
			expected: &HALAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					TriggeringMeasurement: uuid.EmptyUUID,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				Bearing: &models.ModelsBearing{
					Manufacturer: stringp("SKF"),
					ModelNumber:  stringp("2222"),
				},
			},
			expected: &HALAlarmStatus{
				Bearing: &Bearing{
					Manufacturer: "SKF",
					ModelNumber:  "2222",
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				HalIndex: f64p(math.Pi),
			},
			expected: &HALAlarmStatus{
				HALIndex: f64p(math.Pi),
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				FaultFrequency: f64p(math.Pi),
			},
			expected: &HALAlarmStatus{
				FaultFrequency: f64p(math.Pi),
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				NumberOfHarmonicsUsed: i64p(25),
			},
			expected: &HALAlarmStatus{
				NumberOfHarmonicsUsed: i64p(25),
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				RpmFactor: f64p(10),
			},
			expected: &HALAlarmStatus{
				RPMFactor: f64p(10),
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				ErrorDescription: stringp("only peaks"),
			},
			expected: &HALAlarmStatus{
				ErrorDescription: stringp("only peaks"),
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := new(HALAlarmStatus)

			actual.FromInternal(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_HALAlarmStatus_FromInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var status *HALAlarmStatus

		status.FromInternal(&models.ModelsGetAlarmStatusResponseHALAlarm{})
	})

	assert.NotPanics(t, func() {
		status := new(HALAlarmStatus)

		status.FromInternal(nil)
	})
}

func Test_HALAlarmStatus_FromEvent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    events.HalAlarmStatus
		expected *HALAlarmStatus
	}{
		{
			given:    events.HalAlarmStatus{},
			expected: &HALAlarmStatus{},
		},
		{
			given: events.HalAlarmStatus{
				Label: "10x RPM",
			},
			expected: &HALAlarmStatus{
				Label: "10x RPM",
			},
		},
		{
			given: events.HalAlarmStatus{
				Status: 0, // not configured
			},
			expected: &HALAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusNotConfigured,
				},
			},
		},
		{
			given: events.HalAlarmStatus{
				Status: 1, // no data
			},
			expected: &HALAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusNoData,
				},
			},
		},
		{
			given: events.HalAlarmStatus{
				Status: 2, // good
			},
			expected: &HALAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusGood,
				},
			},
		},
		{
			given: events.HalAlarmStatus{
				Status: 3, // alert
			},
			expected: &HALAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusAlert,
				},
			},
		},
		{
			given: events.HalAlarmStatus{
				Status: 4, // danger
			},
			expected: &HALAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusDanger,
				},
			},
		},
		{
			given: events.HalAlarmStatus{
				TriggeringMeasurement: uuid.EmptyUUID,
			},
			expected: &HALAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					TriggeringMeasurement: uuid.EmptyUUID,
				},
			},
		},
		{
			given: events.HalAlarmStatus{
				Bearing: &events.Bearing{
					Manufacturer: "SKF",
					ModelNumber:  "2222",
				},
			},
			expected: &HALAlarmStatus{
				Bearing: &Bearing{
					Manufacturer: "SKF",
					ModelNumber:  "2222",
				},
			},
		},
		{
			given: events.HalAlarmStatus{
				HALIndex: f64p(math.Pi),
			},
			expected: &HALAlarmStatus{
				HALIndex: f64p(math.Pi),
			},
		},
		{
			given: events.HalAlarmStatus{
				FaultFrequency: f64p(math.Pi),
			},
			expected: &HALAlarmStatus{
				FaultFrequency: f64p(math.Pi),
			},
		},
		{
			given: events.HalAlarmStatus{
				NumberOfHarmonicsUsed: i64p(25),
			},
			expected: &HALAlarmStatus{
				NumberOfHarmonicsUsed: i64p(25),
			},
		},
		{
			given: events.HalAlarmStatus{
				RpmFactor: f64p(10),
			},
			expected: &HALAlarmStatus{
				RPMFactor: f64p(10),
			},
		},
		{
			given: events.HalAlarmStatus{
				ErrorDescription: stringp("only peaks"),
			},
			expected: &HALAlarmStatus{
				ErrorDescription: stringp("only peaks"),
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := new(HALAlarmStatus)

			actual.FromEvent(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_HALAlarmStatus_FromEvent_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var h *HALAlarmStatus

		h.FromEvent(events.HalAlarmStatus{})
	})
}
