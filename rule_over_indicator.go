/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

type OverIndicatorRule struct {
	*CommonRule
	indicator1 IIndicator
	indicator2 IIndicator
	threshold  float64
}

/**
threshold is a percentage, for example threshold = 0.1, means 0.10%

if zero, it takes strict value ind1>ind2
otherwise - it takes percentage of ind1 > (1 + thresholdInPercent/100) * ind2


ind1 - ind2  -

*/
func NewOverIndicatorRule(indicator1, indicator2 IIndicator, thresholdInPercent float64, cache *RuleCache) IRule {
	r := &OverIndicatorRule{
		indicator1: indicator1,
		indicator2: indicator2,
		threshold:  thresholdInPercent,
		CommonRule: &CommonRule{
			name:  indicator1.Name() + " Over " + indicator2.Name(),
			cache: cache,
		},
	}
	r.rule = r
	return r.cache.Set(r)
}

func (r *OverIndicatorRule) IsSatisfied(date int64) (bool, error) {
	if r.IsCalculated(date) {
		return r.satisfied, nil
	}
	if r.threshold == 0 {
		return r.done(date, r.indicator1.Val() > r.indicator2.Val()), nil
	} else {
		cmp := (1 + r.threshold/100) * r.indicator2.Val()
		return r.done(date, r.indicator1.Val() > cmp), nil
	}
}
