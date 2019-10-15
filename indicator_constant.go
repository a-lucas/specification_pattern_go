/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

type ConstantIndicator struct {
	*IndicatorCommon
	constantValue float64
}

func NewIndicatorConstant(constant float64, cache *IndicatorCache) IIndicator {
	ind := &ConstantIndicator{
		constantValue:   constant,
		IndicatorCommon: NewIndicatorCommon(CONSTANT, FloatToStringPrecision(constant, 2), ""),
	}
	if cache.Has(ind.Name()) {
		return cache.Get(ind.Name())
	}
	ind.SetCache(cache)
	ind.indicatorCache.Set(ind)
	return ind
}

func (ind *ConstantIndicator) Calculate(date int64, add bool) error {
	if ind.IsCalculated(date) {
		return nil
	}
	ind.val = ind.constantValue
	return ind.done(date, add)
}
