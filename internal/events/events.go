package events

import (
	"github.com/SKF/go-eventsource/v2/eventsource"
	"github.com/SKF/go-utility/v2/uuid"
)

// These types are taken directly from the PAS API.

type SetPointAlarmThresholdEvent struct {
	*eventsource.BaseEvent
	Type       int32  `json:"thresholdType"`
	Inspection []byte `json:"inspection"`

	FullScale    *float64 `json:"thresholdFullScale,omitempty"`
	BandAlarms   [][]byte `json:"thresholdBandAlarms,omitempty"`
	HalAlarms    [][]byte `json:"thresholdHalAlarms,omitempty"`
	Overall      []byte   `json:"thresholdOverall"`
	RateOfChange []byte   `json:"thresholdRateOfChange,omitempty"`
}

type (
	PointAlarmStatusEvent struct {
		*eventsource.BaseEvent
		AlarmStatus   int32 `json:"alarmStatus"`
		AlarmsChanged bool  `json:"alarmsChanged"`
		UpdatedAt     int64 `json:"updatedAt"`

		BandAlarms        []BandAlarmStatus `json:"bandAlarms"`
		HalAlarms         []HalAlarmStatus  `json:"halAlarms"`
		OverallAlarm      *GenericAlarm     `json:"overallAlarm,omitempty"`
		ExternalAlarm     *ExternalAlarm    `json:"externalAlarm,omitempty"`
		InspectionAlarm   *GenericAlarm     `json:"inspectionAlarm,omitempty"`
		RateOfChangeAlarm *GenericAlarm     `json:"rateOfChangeAlarm,omitempty"`
	}

	ExternalAlarm struct {
		Status int32      `json:"status"`
		SetBy  *uuid.UUID `json:"setBy,omitempty"`
	}

	GenericAlarm struct {
		TriggeringMeasurement uuid.UUID `json:"triggeringMeasurement"`
		Status                int32     `json:"status"`
	}

	CalculatedOverall struct {
		Unit  string  `json:"unit"`
		Value float64 `json:"value"`
	}

	Frequency struct {
		ValueType int32   `json:"valueType"`
		Value     float64 `json:"value"`
	}

	BandAlarmStatus struct {
		Label                 string             `json:"label"`
		MinFrequency          Frequency          `json:"minFrequency"`
		MaxFrequency          Frequency          `json:"maxFrequency"`
		CalculatedOverall     *CalculatedOverall `json:"calculatedOverall,omitempty"`
		Status                int32              `json:"status"`
		TriggeringMeasurement uuid.UUID          `json:"triggeringMeasurement"`
	}

	Bearing struct {
		Manufacturer string `json:"manufacturer"`
		ModelNumber  string `json:"modelNumber"`
	}

	HalAlarmStatus struct {
		Status                int32     `json:"status"`
		TriggeringMeasurement uuid.UUID `json:"triggeringMeasurement"`
		Label                 string    `json:"label"`
		Bearing               *Bearing  `json:"bearing,omitempty"`
		FaultFrequency        *float64  `json:"faultFrequency,omitempty"`
		HALIndex              *float64  `json:"halIndex,omitempty"`
		NumberOfHarmonicsUsed *int64    `json:"numberOfHarmonicsUsed,omitempty"`
		RpmFactor             *float64  `json:"rpmFactor,omitempty"`
		ErrorDescription      *string   `json:"errorDescription,omitempty"`
	}
)
