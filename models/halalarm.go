package models

import (
	models "github.com/SKF/go-pas-client/internal/models"
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
