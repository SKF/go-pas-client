package models

import models "github.com/SKF/go-pas-client/internal/models"

type RateOfChange struct {
	OuterHigh *float64
	InnerHigh *float64
	InnerLow  *float64
	OuterLow  *float64
	Unit      string
}

func (r *RateOfChange) FromInternal(internal *models.ModelsRateOfChange) {
	if r == nil || internal == nil {
		return
	}

	r.Unit = internal.Unit
	r.OuterHigh = internal.OuterHigh
	r.InnerHigh = internal.InnerHigh
	r.InnerLow = internal.InnerLow
	r.OuterLow = internal.OuterLow
}

func (r *RateOfChange) ToInternal() *models.ModelsRateOfChange {
	if r == nil {
		return nil
	}

	return &models.ModelsRateOfChange{
		Unit:      r.Unit,
		OuterHigh: r.OuterHigh,
		InnerHigh: r.InnerHigh,
		InnerLow:  r.InnerLow,
		OuterLow:  r.OuterLow,
	}
}
