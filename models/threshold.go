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

	if internal.Overall != nil {
		t.Overall = new(Overall)
		t.Overall.FromInternal(internal.Overall)
	}

	if internal.RateOfChange != nil {
		t.RateOfChange = new(RateOfChange)
		t.RateOfChange.FromInternal(internal.RateOfChange)
	}

	if internal.Inspection != nil {
		t.Inspection = new(Inspection)
		t.Inspection.FromInternal(internal.Inspection)
	}

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

func (t *Threshold) ToInternal() models.ModelsSetPointAlarmThresholdRequest {
	if t == nil {
		return models.ModelsSetPointAlarmThresholdRequest{} // nolint:exhaustivestruct
	}

	thresholdType := int32(t.ThresholdType)

	threshold := models.ModelsSetPointAlarmThresholdRequest{
		Origin:        nil,
		ThresholdType: &thresholdType,
		Overall:       t.Overall.ToInternal(),
		RateOfChange:  t.RateOfChange.ToInternal(),
		Inspection:    t.Inspection.ToInternal(),
		FullScale:     0,
		BandAlarms:    make([]*models.ModelsBandAlarm, len(t.BandAlarms)),
		HalAlarms:     make([]*models.ModelsHALAlarm, len(t.HALAlarms)),
	}

	if t.FullScale != nil {
		threshold.FullScale = *t.FullScale
	}

	for i, bandAlarm := range t.BandAlarms {
		threshold.BandAlarms[i] = bandAlarm.ToInternal()
	}

	for i, halAlarm := range t.HALAlarms {
		threshold.HalAlarms[i] = halAlarm.ToInternal()
	}

	return threshold
}
