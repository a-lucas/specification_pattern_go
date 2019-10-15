/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

type NotRule struct {
	*CommonRule
}

func NewNotRule(rule IRule, cache *RuleCache) IRule {
	r := &NotRule{
		CommonRule: &CommonRule{
			rule:  rule,
			name:  "NOT " + rule.Name(),
			cache: cache,
		},
	}
	return r.cache.Set(r)
}

func (r *NotRule) IsSatisfied(date int64) (bool, error) {
	satisfied, err := r.rule.IsSatisfied(date)
	return !satisfied, err
}
