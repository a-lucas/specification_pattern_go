/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

type OrRule struct {
	*CommonRule
	rule1 IRule
	rule2 IRule
}

func NewOrRule(rule1, rule2 IRule, cache *RuleCache) IRule {
	r := &OrRule{
		rule1: rule1,
		rule2: rule2,
		CommonRule: &CommonRule{
			name:  rule1.Name() + " OR " + rule2.Name(),
			cache: cache,
		},
	}
	r.rule = r
	return r.cache.Set(r)
}

func (r *OrRule) IsSatisfied(date int64) (bool, error) {
	if r.IsCalculated(date) {
		return r.satisfied, nil
	}
	if satisfied1, err := r.rule1.IsSatisfied(date); err != nil {
		return false, err
	} else {
		if satisfied1 {
			return r.done(date, true), nil
		}
		if satisfied2, err := r.rule2.IsSatisfied(date); err != nil {
			return false, err
		} else {
			return r.done(date, satisfied1 || satisfied2), nil
		}
	}
}
