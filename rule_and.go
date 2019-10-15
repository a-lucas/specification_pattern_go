/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

type AbstractRule struct{}

type AndRule struct {
	*CommonRule
	rule1 IRule
	rule2 IRule
}

func NewAndRule(rule1, rule2 IRule, cache *RuleCache) IRule {
	r := &AndRule{
		rule1: rule1,
		rule2: rule2,
		CommonRule: &CommonRule{
			name:                 "(" + rule1.Name() + ") And (" + rule2.Name() + ")",
			indicatorsCalculated: false,
			cache:                cache,
		}}
	r.rule = r
	return r.cache.Set(r)
}

func (r *AndRule) IsSatisfied(date int64) (bool, error) {
	if r.IsCalculated(date) {
		return r.satisfied, nil
	}
	if satisfied1, err := r.rule1.IsSatisfied(date); err != nil {
		return satisfied1, err
	} else {
		if !satisfied1 {
			return r.done(date, false), nil
		}
		if satisfied2, err := r.rule2.IsSatisfied(date); err != nil {
			return satisfied2, err
		} else {
			return r.done(date, satisfied1 && satisfied2), nil
		}
	}
}

/*

public class AndRule extends AbstractRule {

private Rule rule1;

private Rule rule2;


public AndRule(Rule rule1, Rule rule2) {
this.rule1 = rule1;
this.rule2 = rule2;
}

@Override
public boolean isSatisfied(int index, TradingRecord tradingRecord) {
final boolean satisfied = rule1.isSatisfied(index, tradingRecord) && rule2.isSatisfied(index, tradingRecord);
traceIsSatisfied(index, satisfied);
return satisfied;
}
}

*/
//
//func( r *AbstractRule) And(rule *AbstractRule) *AbstractRule {
//	return NewAndRule(r, rule)
//}
//
//
//type Rule interface {
//	And(rule Rule) Rule
//default Rule and(Rule rule) {
//return new AndRule(this, rule);
//}
//
///**
// * @param rule another trading rule
// * @return a rule which is the OR combination of this rule with the provided one
// */
//default Rule or(Rule rule) {
//return new OrRule(this, rule);
//}
//
///**
// * @param rule another trading rule
// * @return a rule which is the XOR combination of this rule with the provided one
// */
//default Rule xor(Rule rule) {
//return new XorRule(this, rule);
//}
//
///**
// * @return a rule which is the logical negation of this rule
// */
//default Rule negation() {
//return new NotRule(this);
//}
//
///**
// * @param index the bar index
// * @return true if this rule is satisfied for the provided index, false otherwise
// */
//default boolean isSatisfied(int index) {
//return isSatisfied(index, null);
//}
//
///**
// * @param index the bar index
// * @param tradingRecord the potentially needed trading history
// * @return true if this rule is satisfied for the provided index, false otherwise
// */
//boolean isSatisfied(int index, TradingRecord tradingRecord);
//}
//
