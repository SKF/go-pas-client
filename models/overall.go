package models

import models "github.com/SKF/go-pas-client/internal/models"

type Overall struct {
	OuterHigh *float64
	InnerHigh *float64
	InnerLow  *float64
	OuterLow  *float64
	Unit      string
}

func (o *Overall) FromInternal(internal *models.ModelsOverall) {
	if o == nil || internal == nil {
		return
	}

	o.Unit = internal.Unit
	o.OuterHigh = internal.OuterHigh
	o.InnerHigh = internal.InnerHigh
	o.InnerLow = internal.InnerLow
	o.OuterLow = internal.OuterLow
}