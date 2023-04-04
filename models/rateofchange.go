//nolint:dupl
package models

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	models "github.com/SKF/go-pas-client/internal/models"
	pas "github.com/SKF/proto/v2/pas"
)

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

func (r RateOfChange) ToInternal() *models.ModelsRateOfChange {
	return &models.ModelsRateOfChange{
		Unit:      r.Unit,
		OuterHigh: r.OuterHigh,
		InnerHigh: r.InnerHigh,
		InnerLow:  r.InnerLow,
		OuterLow:  r.OuterLow,
	}
}

func (r *RateOfChange) FromProto(buf []byte) error {
	if r == nil || len(buf) == 0 {
		return nil
	}

	var internal pas.RateOfChange

	if err := proto.Unmarshal(buf, &internal); err != nil {
		return fmt.Errorf("decoding rate of change alarm failed: %w", err)
	}

	r.Unit = internal.Unit

	if internal.OuterHigh != nil {
		r.OuterHigh = &internal.OuterHigh.Value
	}

	if internal.InnerHigh != nil {
		r.InnerHigh = &internal.InnerHigh.Value
	}

	if internal.InnerLow != nil {
		r.InnerLow = &internal.InnerLow.Value
	}

	if internal.OuterLow != nil {
		r.OuterLow = &internal.OuterLow.Value
	}

	return nil
}
