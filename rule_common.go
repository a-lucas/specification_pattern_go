/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

type CommonRule struct {
	index                int
	rule                 IRule
	name                 string
	cache                *RuleCache
	indicatorsCalculated bool
	triggerId            int
	isCalculated         int64
	satisfied            bool
}

func (cr *CommonRule) IsCalculated(date int64) bool {
	return cr.isCalculated == date
}

func (cr *CommonRule) done(date int64, satisfied bool) bool {
	cr.isCalculated = date
	cr.satisfied = satisfied
	return satisfied
}

func (cr *CommonRule) TriggerId() int {
	return cr.triggerId
}

func (cr *CommonRule) Name() string {
	return cr.name
}

func (cr *CommonRule) Index() int {
	return cr.index
}

func (cr *CommonRule) SetIndex(index int) {
	cr.index = index
}

func (cr *CommonRule) And(rule IRule) IRule {
	return NewAndRule(cr.rule, rule, cr.cache)
}

func (cr *CommonRule) Or(rule IRule) IRule {
	return NewOrRule(cr.rule, rule, cr.cache)
}

func (cr *CommonRule) Xor(rule IRule) IRule {
	return NewXorRule(cr.rule, rule, cr.cache)
}

func (cr *CommonRule) Negation() IRule {
	return NewNotRule(cr.rule, cr.cache)
}
