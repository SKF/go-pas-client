package models

import (
	models "github.com/SKF/go-pas-client/internal/models"
	"github.com/SKF/go-utility/v2/uuid"
)

type Threshold struct {
	NodeID        uuid.UUID
	ThresholdType ThresholdType
	Overall       *Overall
	RateOfChange  *RateOfChange
	Inspection    *Inspection
	FullScale     *float64
	BandAlarms    []BandAlarm
	HALAlarms     []HALAlarm
}

func (t *Threshold) FromInternal(internal models.ModelsGetPointAlarmThresholdResponse) (err error) {
	if t == nil {
		return
	}

	t.ThresholdType = ThresholdTypeNone
	t.Overall.FromInternal(internal.Overall)
	t.RateOfChange.FromInternal(internal.RateOfChange)
	t.Inspection.FromInternal(internal.Inspection)
	t.FullScale = internal.FullScale
	t.BandAlarms = make([]BandAlarm, len(internal.BandAlarms))
	t.HALAlarms = make([]HALAlarm, len(internal.HalAlarms))

	if internal.NodeID != nil {
		t.NodeID = uuid.UUID(internal.NodeID.String())

		if err = t.NodeID.Validate(); err != nil {
			return
		}
	}

	if internal.ThresholdType != nil {
		t.ThresholdType = ThresholdType(*internal.ThresholdType)
	}

	for i, bandAlarm := range internal.BandAlarms {
		t.BandAlarms[i].FromInternal(bandAlarm)
	}

	for i, halAlarm := range internal.HalAlarms {
		t.HALAlarms[i].FromInternal(halAlarm)
	}

	return nil
}
