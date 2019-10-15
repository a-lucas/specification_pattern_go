/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

type BooleanRule struct {
	*CommonRule
	val bool
}

func NewBooleanRule(val bool, cache *RuleCache) IRule {
	r := &BooleanRule{
		val: val,
		CommonRule: &CommonRule{
			cache: cache,
		},
	}
	if val {
		r.name = "Always True"
	} else {
		r.name = "Always False"
	}
	r.rule = r
	return r.cache.Set(r)
}

func (r *BooleanRule) IsSatisfied(date int64) (bool, error) {
	return r.val, nil
}
