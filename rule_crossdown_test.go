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

func TestRuleCrossDown(t *testing.T) {

	newRecord := func(val float64, d time.Time) *AskBidTrade {
		return &AskBidTrade{
			InputDate: d,
			Ask: Record{
				Close: val,
			},
		}
	}

	t.Run("Cross Constant With SMA indicator 1 Up ", func(t *testing.T) {

		g := NewGomegaWithT(t)
		history := NewPerformanceAskBidTradeHistory()
		cache := NewIndicatorCache(history)

		ruleCache := NewRuleCache(history)

		constant := NewIndicatorConstant(5.0, cache)
		live := NewIndicatorLive(AskClose, cache, history)
		linear := NewIndicatorSMA(1, live, cache)

		start := time.Now()
		add := func(val float64, nbMinutes int) {
			start = start.Add(time.Duration(nbMinutes) * time.Minute)
			history.Append(newRecord(val, start))
			cache.Calculate(newRecord(val, start), true)
		}

		rule := NewCrossDownIndicatorRule(constant, linear, 0.1, ruleCache)

		add(1, 1)
		ruleCache.Calculate(start)
		g.Expect(ruleCache.IsSatisfied(start.Unix(), rule.Index())).To(BeFalse())

		add(2, 2)
		ruleCache.Calculate(start)
		g.Expect(ruleCache.IsSatisfied(start.Unix(), rule.Index())).To(BeFalse())

		add(4.8, 3)
		ruleCache.Calculate(start)
		g.Expect(ruleCache.IsSatisfied(start.Unix(), rule.Index())).To(BeFalse())

		add(5.1, 4)
		ruleCache.Calculate(start)
		g.Expect(ruleCache.IsSatisfied(start.Unix(), rule.Index())).To(BeFalse())

		add(5.9, 5)
		ruleCache.Calculate(start)
		g.Expect(ruleCache.IsSatisfied(start.Unix(), rule.Index())).To(BeTrue())

		add(6, 6)
		ruleCache.Calculate(start)
		g.Expect(ruleCache.IsSatisfied(start.Unix(), rule.Index())).To(BeFalse())

	})

	t.Run("With Zeroes A", func(t *testing.T) {

		g := NewGomegaWithT(t)

		history := NewPerformanceAskBidTradeHistory()
		cache := NewIndicatorCache(history)

		constant := NewIndicatorConstant(5.0, cache)
		live := NewIndicatorLive(AskClose, cache, history)
		linear := NewIndicatorSMA(1, live, cache)

		start := time.Now()
		add := func(val float64, nbMinutes int) {
			start = start.Add(time.Duration(nbMinutes) * time.Minute)
			history.Append(newRecord(val, start))
			cache.Calculate(newRecord(val, start), true)
		}

		ruleCache := NewRuleCache(history)
		rule := NewCrossDownIndicatorRule(constant, linear, 0.1, ruleCache)

		add(4, 1)
		ruleCache.Calculate(start)
		g.Expect(ruleCache.IsSatisfied(start.Unix(), rule.Index())).To(BeFalse())

		add(4.9, 2)
		ruleCache.Calculate(start)
		g.Expect(ruleCache.IsSatisfied(start.Unix(), rule.Index())).To(BeFalse())

		add(5.1, 3)
		ruleCache.Calculate(start)
		g.Expect(ruleCache.IsSatisfied(start.Unix(), rule.Index())).To(BeFalse())

		add(4.9, 4)
		ruleCache.Calculate(start)
		g.Expect(ruleCache.IsSatisfied(start.Unix(), rule.Index())).To(BeFalse())

		add(4, 5)
		ruleCache.Calculate(start)
		g.Expect(ruleCache.IsSatisfied(start.Unix(), rule.Index())).To(BeFalse())

		add(6, 6)
		ruleCache.Calculate(start)
		g.Expect(ruleCache.IsSatisfied(start.Unix(), rule.Index())).To(BeTrue())

		add(4.9, 7)
		ruleCache.Calculate(start)
		g.Expect(ruleCache.IsSatisfied(start.Unix(), rule.Index())).To(BeFalse())

		add(6, 8)
		ruleCache.Calculate(start)
		g.Expect(ruleCache.IsSatisfied(start.Unix(), rule.Index())).To(BeFalse())

		add(4, 9)
		ruleCache.Calculate(start)
		g.Expect(ruleCache.IsSatisfied(start.Unix(), rule.Index())).To(BeFalse())

		add(6, 10)
		ruleCache.Calculate(start)
		g.Expect(ruleCache.IsSatisfied(start.Unix(), rule.Index())).To(BeTrue())

	})

}
