// nolint:dupl
package models

import (
	"fmt"

	p "google.golang.org/protobuf/proto"

	models "github.com/SKF/go-pas-client/internal/models"
	pas "github.com/SKF/proto/v2/pas"
)

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

func (o Overall) ToInternal() *models.ModelsOverall {
	return &models.ModelsOverall{
		Unit:      o.Unit,
		OuterHigh: o.OuterHigh,
		InnerHigh: o.InnerHigh,
		InnerLow:  o.InnerLow,
		OuterLow:  o.OuterLow,
	}
}

func (o *Overall) FromProto(buf []byte) error {
	if o == nil || len(buf) == 0 {
		return nil
	}

	var internal pas.Overall

	if err := p.Unmarshal(buf, &internal); err != nil {
		return fmt.Errorf("decoding overall alarm failed: %w", err)
	}

	o.Unit = internal.Unit

	if internal.OuterHigh != nil {
		o.OuterHigh = &internal.OuterHigh.Value
	}

	if internal.InnerHigh != nil {
		o.InnerHigh = &internal.InnerHigh.Value
	}

	if internal.InnerLow != nil {
		o.InnerLow = &internal.InnerLow.Value
	}

	if internal.OuterLow != nil {
		o.OuterLow = &internal.OuterLow.Value
	}

	return nil
}
