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

func TestRuleNot(t *testing.T) {
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

		history.Append(newRecord(1, time.Now()))

		rule1True := NewOverIndicatorRule(ind1, ind2, 0, ruleCache)
		rule2False := NewOverIndicatorRule(ind2, ind1, 0, ruleCache)

		start := time.Now()
		cache.Calculate(newRecord(1, start), true)

		g.Expect(rule1True.IsSatisfied(start.Unix())).To(BeTrue())
		g.Expect(rule2False.IsSatisfied(start.Unix())).To(BeFalse())

		g.Expect(rule1True.Negation().IsSatisfied(start.Unix())).To(BeFalse())
		g.Expect(rule2False.Negation().IsSatisfied(start.Unix())).To(BeTrue())

		notRule1 := NewNotRule(rule1True, ruleCache)
		notRule2 := NewNotRule(rule2False, ruleCache)

		g.Expect(notRule1.IsSatisfied(start.Unix())).To(BeFalse())
		g.Expect(notRule2.IsSatisfied(start.Unix())).To(BeTrue())

	})

}
