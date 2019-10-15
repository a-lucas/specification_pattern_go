/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

import (
	"math"
)

type IndicatorAbstractEma struct {
	barCount        float64
	multiplier      float64
	sourceIndicator IIndicator
	_prevValue      float64
	*IndicatorCommon
}

func NewIndicatorAbstractEma(indicatorType IndicatorType, source IIndicator, barCount float64, multiplier float64, cache *IndicatorCache) IIndicator {
	ind := &IndicatorAbstractEma{
		sourceIndicator: source,
		multiplier:      multiplier,
		barCount:        barCount,
		IndicatorCommon: NewIndicatorCommon(indicatorType, PointToPeriod(int(barCount)), source.Name()),
	}
	if cache.Has(ind.Name()) {
		return cache.Get(ind.Name())
	}
	ind.SetCache(cache)
	cache.Set(ind)
	return ind
}

func (ind *IndicatorAbstractEma) Calculate(date int64, add bool) error {
	if ind.IsCalculated(date) {
		return nil
	}
	if !ind.sourceIndicator.IsCalculated(date) {
		if err := ind.sourceIndicator.Calculate(date, add); err != nil {
			return err
		}
	}
	if ind.indicatorCache.history.Len() == 1 {
		ind.val = ind.sourceIndicator.Val()
	} else {
		ind.val = ind._prevValue + (ind.sourceIndicator.Val()-ind._prevValue)*ind.multiplier
		if math.IsNaN(ind.val) {
			panic("It is Nan and it shouldnt")
		}
	}
	ind._prevValue = ind.val
	return ind.done(date, add)
}
