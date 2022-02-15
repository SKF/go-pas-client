package models

import (
	"github.com/SKF/go-pas-client/internal/events"
	models "github.com/SKF/go-pas-client/internal/models"
	"github.com/SKF/go-utility/v2/uuid"
)

type (
	HALAlarm struct {
		Label        string
		Bearing      *Bearing
		HALAlarmType HALAlarmType
		UpperDanger  *float64
		UpperAlert   *float64
	}

	Bearing struct {
		Manufacturer string
		ModelNumber  string
	}
)

func (h *HALAlarm) FromInternal(internal *models.ModelsHALAlarm) {
	if h == nil || internal == nil {
		return
	}

	h.Label = internal.Label
	h.HALAlarmType = HALAlarmType(internal.HalAlarmType)
	h.UpperAlert = internal.UpperAlert
	h.UpperDanger = internal.UpperDanger

	if internal.Bearing != nil {
		h.Bearing = &Bearing{
			Manufacturer: *internal.Bearing.Manufacturer,
			ModelNumber:  *internal.Bearing.ModelNumber,
		}
	}
}

func (h *HALAlarm) ToInternal() *models.ModelsHALAlarm {
	if h == nil {
		return nil
	}

	halAlarm := &models.ModelsHALAlarm{
		Label:        h.Label,
		HalAlarmType: string(h.HALAlarmType),
		UpperAlert:   h.UpperAlert,
		UpperDanger:  h.UpperDanger,
		Bearing:      nil,
	}

	if h.Bearing != nil {
		halAlarm.Bearing = &models.ModelsBearing{
			Manufacturer: &h.Bearing.Manufacturer,
			ModelNumber:  &h.Bearing.ModelNumber,
		}
	}

	return halAlarm
}

type HALAlarmStatus struct {
	GenericAlarmStatus
	Label                 string
	Bearing               *Bearing
	HALIndex              *float64
	FaultFrequency        *float64
	RPMFactor             *float64
	NumberOfHarmonicsUsed *int64
	ErrorDescription      *string
}

func (h *HALAlarmStatus) FromInternal(internal *models.ModelsGetAlarmStatusResponseHALAlarm) {
	if h == nil || internal == nil {
		return
	}

	if internal.Label != nil {
		h.Label = *internal.Label
	}

	if internal.Status != nil {
		h.Status = AlarmStatusType(*internal.Status)
	}

	if internal.TriggeringMeasurement != nil {
		h.TriggeringMeasurement = uuid.UUID(internal.TriggeringMeasurement.String())
	}

	if internal.Bearing != nil {
		h.Bearing = &Bearing{
			Manufacturer: *internal.Bearing.Manufacturer,
			ModelNumber:  *internal.Bearing.ModelNumber,
		}
	}

	h.FaultFrequency = internal.FaultFrequency
	h.RPMFactor = internal.RpmFactor
	h.HALIndex = internal.HalIndex
	h.NumberOfHarmonicsUsed = internal.NumberOfHarmonicsUsed
	h.ErrorDescription = internal.ErrorDescription
}

func (h *HALAlarmStatus) FromEvent(internal events.HalAlarmStatus) {
	if h == nil {
		return
	}

	h.Label = internal.Label
	h.Status = AlarmStatusType(internal.Status)
	h.TriggeringMeasurement = internal.TriggeringMeasurement

	if internal.Bearing != nil {
		h.Bearing = &Bearing{
			Manufacturer: internal.Bearing.Manufacturer,
			ModelNumber:  internal.Bearing.ModelNumber,
		}
	}

	h.FaultFrequency = internal.FaultFrequency
	h.RPMFactor = internal.RpmFactor
	h.HALIndex = internal.HALIndex
	h.NumberOfHarmonicsUsed = internal.NumberOfHarmonicsUsed
	h.ErrorDescription = internal.ErrorDescription
}
