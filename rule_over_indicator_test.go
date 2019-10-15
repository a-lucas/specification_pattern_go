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

func TestRuleOverIndicator(t *testing.T) {

	newRecord := func(val float64, d time.Time) *AskBidTrade {
		return &AskBidTrade{
			InputDate: d,
			Ask: Record{
				Close: val,
			},
		}
	}

	t.Run("between two constant indicators with zero threshold", func(t *testing.T) {

		g := NewGomegaWithT(t)
		history := NewPerformanceAskBidTradeHistory()
		cache := NewIndicatorCache(history)
		ruleCache := NewRuleCache(history)

		ind1 := NewIndicatorConstant(5, cache)
		ind2 := NewIndicatorConstant(3, cache)

		start := time.Now()
		history.Append(newRecord(1, start))

		rule1 := NewOverIndicatorRule(ind1, ind2, 0, ruleCache)
		cache.Calculate(newRecord(1, time.Now()), true)
		g.Expect(rule1.IsSatisfied(start.Unix())).To(BeTrue())

		rule2 := NewOverIndicatorRule(ind2, ind1, 0, ruleCache)
		g.Expect(rule2.IsSatisfied(start.Unix())).To(BeFalse())

	})

	t.Run("between two constant indicators with 50% threshold", func(t *testing.T) {

		g := NewGomegaWithT(t)
		history := NewPerformanceAskBidTradeHistory()
		cache := NewIndicatorCache(history)
		ruleCache := NewRuleCache(history)

		ind1 := NewIndicatorConstant(5, cache)
		ind2 := NewIndicatorConstant(4, cache)
		ind3 := NewIndicatorConstant(3, cache)

		rule1 := NewOverIndicatorRule(ind1, ind2, 50, ruleCache)
		rule2 := NewOverIndicatorRule(ind1, ind3, 50, ruleCache)
		rule3 := NewOverIndicatorRule(ind3, ind1, 50, ruleCache)
		rule4 := NewOverIndicatorRule(ind2, ind1, 50, ruleCache)

		start := time.Now()
		history.Append(newRecord(1, start))
		cache.Calculate(newRecord(1, start), true)

		s1, err := rule1.IsSatisfied(start.Unix())
		g.Expect(err).To(BeNil())
		g.Expect(s1).To(BeFalse())

		s2, err := rule2.IsSatisfied(start.Unix())
		g.Expect(err).To(BeNil())
		g.Expect(s2).To(BeTrue())

		s3, err := rule3.IsSatisfied(start.Unix())
		g.Expect(err).To(BeNil())
		g.Expect(s3).To(BeFalse())

		s4, err := rule4.IsSatisfied(start.Unix())
		g.Expect(err).To(BeNil())
		g.Expect(s4).To(BeFalse())

	})
}
