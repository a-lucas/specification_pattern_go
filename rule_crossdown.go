/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

type CrossDownIndicatorRule struct {
	*CommonRule
	indicator1 IIndicator
	indicator2 IIndicator
	crosser    *Crosser
}

func NewCrossDownIndicatorRule(indicator1, indicator2 IIndicator, threshold float64, cache *RuleCache) IRule {

	r := &CrossDownIndicatorRule{
		indicator1: indicator1,
		indicator2: indicator2,
		crosser:    NewCrosser(threshold, CrossPositionUnder),
		CommonRule: &CommonRule{
			name:  indicator1.Name() + " CROSS DOWN " + indicator2.Name() + " Param1 " + FloatToStringPrecision(100*threshold, 2),
			cache: cache,
		},
	}
	r.rule = r
	return r.cache.Set(r)
}

func (r *CrossDownIndicatorRule) IsSatisfied(date int64) (bool, error) {
	if r.IsCalculated(date) {
		return r.satisfied, nil
	}
	satisfied, err := r.crosser.Calculate(r.indicator1.Val(), r.indicator2.Val())
	return r.done(date, satisfied), err
}
