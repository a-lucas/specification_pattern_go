/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

import (
	"github.com/davecgh/go-spew/spew"
)

type IndicatorCommon struct {
	vals           []float64
	val            float64
	name           string
	history        *PerformanceAskBidTradeHistory
	indicatorCache *IndicatorCache
	indicatorType  IndicatorType
	isCalculated   int64
}

func NewIndicatorCommon(indicatorType IndicatorType, name string, sourceName string) *IndicatorCommon {
	return &IndicatorCommon{
		vals:          make([]float64, 0),
		name:          string(indicatorType) + " " + name + "(" + sourceName + ")",
		indicatorType: indicatorType,
		isCalculated:  0,
	}
}

func (ic *IndicatorCommon) done(date int64, add bool) error {
	if add {
		ic.vals = append(ic.vals, ic.val)
	}
	ic.isCalculated = date
	return nil
}

func (ic *IndicatorCommon) IsCalculated(date int64) bool {
	return ic.isCalculated == date
}

func (ic *IndicatorCommon) SetHistory(history *PerformanceAskBidTradeHistory) {
	ic.history = history
}

func (ic *IndicatorCommon) Type() IndicatorType {
	return ic.indicatorType
}

func (ic *IndicatorCommon) SetCache(cache *IndicatorCache) {
	ic.indicatorCache = cache
}

func (ic *IndicatorCommon) Val() float64 {
	return ic.val
}

func (ic *IndicatorCommon) Len() int {
	return len(ic.vals)
}

func (ic *IndicatorCommon) GetLastValues(nb int) []float64 {
	if nb >= len(ic.vals) {
		ret := make([]float64, len(ic.vals))
		copy(ret, ic.vals)
		return ret
	}
	return ic.vals[len(ic.vals)-nb:]
}

func (ic *IndicatorCommon) GetValue(index int) float64 {
	if ic.Length() == 0 {
		spew.Dump(ic)
		panic("invalid indicator length requested = there are no records yet")
	}
	if ic.Length() < index {
		spew.Dump(ic)
		panic("invalid indicator length requested")
	}
	return ic.vals[index]
}

func (ic *IndicatorCommon) GetValues(fromIndex, toIndex int) []float64 {
	if ic.Length() < toIndex {
		spew.Dump(ic)
		panic("invalid indicator length requested")
	}
	if fromIndex < 0 {
		fromIndex = 0
	}
	return ic.vals[fromIndex:toIndex]
}

func (ic *IndicatorCommon) Length() int {
	return len(ic.vals)
}

func (ic *IndicatorCommon) Name() string {
	return ic.name
}

func (ic *IndicatorCommon) Values() []float64 {
	return ic.vals
}
