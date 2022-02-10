package models

import models "github.com/SKF/go-pas-client/internal/models"

type (
	BandAlarm struct {
		Label            string
		MinFrequency     BandAlarmFrequency
		MaxFrequency     BandAlarmFrequency
		OverallThreshold *BandAlarmOverallThreshold
	}

	BandAlarmFrequency struct {
		ValueType BandAlarmFrequencyValueType
		Value     float64
	}

	BandAlarmThreshold struct {
		ValueType BandAlarmThresholdType
		Value     float64
	}

	BandAlarmOverallThreshold struct {
		Unit        string
		UpperAlert  *BandAlarmThreshold
		UpperDanger *BandAlarmThreshold
	}
)

func (b *BandAlarm) FromInternal(internal *models.ModelsBandAlarm) {
	if b == nil || internal == nil {
		return
	}

	b.Label = internal.Label

	if internal.OverallThreshold != nil {
		b.OverallThreshold = new(BandAlarmOverallThreshold)
		b.OverallThreshold.FromInternal(internal.OverallThreshold)
	}

	if internal.MaxFrequency != nil {
		if internal.MaxFrequency.ValueType != nil {
			b.MaxFrequency.ValueType = BandAlarmFrequencyValueType(*internal.MaxFrequency.ValueType)
		}

		if internal.MaxFrequency.Value != nil {
			b.MaxFrequency.Value = *internal.MaxFrequency.Value
		}
	}

	if internal.MinFrequency != nil {
		if internal.MinFrequency.ValueType != nil {
			b.MinFrequency.ValueType = BandAlarmFrequencyValueType(*internal.MinFrequency.ValueType)
		}

		if internal.MinFrequency.Value != nil {
			b.MinFrequency.Value = *internal.MinFrequency.Value
		}
	}
}

func (b *BandAlarmOverallThreshold) FromInternal(internal *models.ModelsBandAlarmOverallThreshold) {
	if b == nil || internal == nil {
		return
	}

	b.Unit = internal.Unit

	if internal.UpperAlert != nil {
		b.UpperAlert = &BandAlarmThreshold{
			ValueType: BandAlarmThresholdTypeUnknown,
			Value:     0,
		}

		if internal.UpperAlert.ValueType != nil {
			b.UpperAlert.ValueType = BandAlarmThresholdType(*internal.UpperAlert.ValueType)
		}

		if internal.UpperAlert.Value != nil {
			b.UpperAlert.Value = *internal.UpperAlert.Value
		}
	}

	if internal.UpperDanger != nil {
		b.UpperDanger = &BandAlarmThreshold{
			ValueType: BandAlarmThresholdTypeUnknown,
			Value:     0,
		}

		if internal.UpperDanger.ValueType != nil {
			b.UpperDanger.ValueType = BandAlarmThresholdType(*internal.UpperDanger.ValueType)
		}

		if internal.UpperDanger.Value != nil {
			b.UpperDanger.Value = *internal.UpperDanger.Value
		}
	}
}
