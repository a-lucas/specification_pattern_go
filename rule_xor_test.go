/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

import (
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

func TestRuleXOr(t *testing.T) {
	newRecord := func(val float64, d time.Time) *AskBidTrade {
		return &AskBidTrade{
			InputDate: d,
			Ask: Record{
				Close: val,
			},
		}
	}

	t.Run("between two Over indicators", func(t *testing.T) {

		g := NewGomegaWithT(t)
		history := NewPerformanceAskBidTradeHistory()
		cache := NewIndicatorCache(history)
		ruleCache := NewRuleCache(history)

		ind1 := NewIndicatorConstant(5, cache)
		ind2 := NewIndicatorConstant(4, cache)
		ind3 := NewIndicatorConstant(3, cache)

		start := time.Now()
		history.Append(newRecord(1, start))

		rule1True := NewOverIndicatorRule(ind1, ind2, 0, ruleCache)
		rule2True := NewOverIndicatorRule(ind2, ind3, 0, ruleCache)
		rule3False := NewOverIndicatorRule(ind3, ind1, 0, ruleCache)
		rule4False := NewOverIndicatorRule(ind3, ind2, 0, ruleCache)

		rule1 := NewXorRule(rule1True, rule2True, ruleCache)
		rule2 := NewXorRule(rule1True, rule3False, ruleCache)
		rule3 := NewXorRule(rule3False, rule2True, ruleCache)
		rule4 := NewXorRule(rule3False, rule4False, ruleCache)

		cache.Calculate(newRecord(1, start), true)

		g.Expect(rule1True.IsSatisfied(start.Unix())).To(BeTrue())
		g.Expect(rule2True.IsSatisfied(start.Unix())).To(BeTrue())
		g.Expect(rule3False.IsSatisfied(start.Unix())).To(BeFalse())
		g.Expect(rule4False.IsSatisfied(start.Unix())).To(BeFalse())

		g.Expect(rule1.IsSatisfied(start.Unix())).To(BeFalse())
		g.Expect(rule2.IsSatisfied(start.Unix())).To(BeTrue())
		g.Expect(rule3.IsSatisfied(start.Unix())).To(BeTrue())
		g.Expect(rule4.IsSatisfied(start.Unix())).To(BeFalse())

	})

	t.Run("test the Xor()", func(t *testing.T) {

		g := NewGomegaWithT(t)
		history := NewPerformanceAskBidTradeHistory()
		cache := NewIndicatorCache(history)
		ruleCache := NewRuleCache(history)

		ind1 := NewIndicatorConstant(5, cache)
		ind2 := NewIndicatorConstant(4, cache)
		ind3 := NewIndicatorConstant(3, cache)

		start := time.Now()
		history.Append(newRecord(1, start))

		rule1True := NewOverIndicatorRule(ind1, ind2, 0, ruleCache)
		rule2True := NewOverIndicatorRule(ind2, ind3, 0, ruleCache)
		rule3False := NewOverIndicatorRule(ind3, ind1, 0, ruleCache)
		rule4False := NewOverIndicatorRule(ind3, ind2, 0, ruleCache)

		rule1 := rule1True.Xor(rule2True)
		rule2 := rule1True.Xor(rule3False)
		rule3 := rule3False.Xor(rule2True)
		rule4 := rule3False.Xor(rule4False)

		cache.Calculate(newRecord(1, time.Now()), true)

		g.Expect(rule1True.IsSatisfied(start.Unix())).To(BeTrue())
		g.Expect(rule2True.IsSatisfied(start.Unix())).To(BeTrue())
		g.Expect(rule3False.IsSatisfied(start.Unix())).To(BeFalse())
		g.Expect(rule4False.IsSatisfied(start.Unix())).To(BeFalse())

		g.Expect(rule1.IsSatisfied(start.Unix())).To(BeFalse())
		g.Expect(rule2.IsSatisfied(start.Unix())).To(BeTrue())
		g.Expect(rule3.IsSatisfied(start.Unix())).To(BeTrue())
		g.Expect(rule4.IsSatisfied(start.Unix())).To(BeFalse())

	})

}
