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

func TestRuleCache(t *testing.T) {

	t.Run("Cache build up successfully", func(t *testing.T) {
		g := NewGomegaWithT(t)

		history := NewPerformanceAskBidTradeHistory()
		cache := NewRuleCache(history)

		g.Expect(len(cache.rules)).To(Equal(0))
		g.Expect(cache.Calculate(time.Now())).To(BeNil())
	})

	t.Run("Some Boolean shit", func(t *testing.T) {
		g := NewGomegaWithT(t)
		history := NewPerformanceAskBidTradeHistory()
		cache := NewRuleCache(history)

		rule1 := NewBooleanRule(true, cache)
		rule2 := NewBooleanRule(false, cache)

		g.Expect(cache.has(rule1.Name())).To(BeTrue())
		g.Expect(cache.has(rule2.Name())).To(BeTrue())
		g.Expect(cache.has("Some bullshit")).To(BeFalse())

		r1 := cache.GetByName(rule1.Name())
		g.Expect(r1).To(BeEquivalentTo(rule1))

		r2 := cache.GetByName(rule2.Name())
		g.Expect(r2).To(BeEquivalentTo(rule2))

		i1 := cache.GetById(rule1.Index())
		g.Expect(i1).To(BeEquivalentTo(rule1))

		i2 := cache.GetById(rule2.Index())
		g.Expect(i2).To(BeEquivalentTo(rule2))

		start := time.Now()
		g.Expect(cache.Calculate(start)).To(BeNil())

		g.Expect(cache.GetSatisfiedRules()).To(BeEquivalentTo([]string{rule1.Name()}))

		satisfied1, err := cache.IsSatisfied(start.Unix(), rule1.Index())
		g.Expect(err).To(BeNil())
		g.Expect(satisfied1).To(BeTrue())

		satisfied2, err := cache.IsSatisfied(start.Unix(), rule2.Index())
		g.Expect(err).To(BeNil())
		g.Expect(satisfied2).To(BeFalse())
	})

	t.Run("Caching Over Rule", func(t *testing.T) {
		g := NewGomegaWithT(t)

		history := NewPerformanceAskBidTradeHistory()
		ruleCache := NewRuleCache(history)
		indicatorCache := NewIndicatorCache(history)

		constantA := NewIndicatorConstant(10, indicatorCache)
		liveIndicator := NewIndicatorLive(AskClose, indicatorCache, history)

		ruleA := NewOverIndicatorRule(liveIndicator, constantA, 0.5, ruleCache)

		start := time.Now()
		add := func(value float64) {

			start = start.Add(time.Minute)
			abd := &AskBidTrade{
				Ask:       Record{value, value, value, value},
				Bid:       Record{value, value, value, value},
				Trade:     Record{value, value, value, value},
				InputDate: start,
			}

			history.Append(abd)
			err := indicatorCache.Calculate(abd, true)
			g.Expect(err).To(BeNil())

			err = ruleCache.Calculate(abd.InputDate)
			g.Expect(err).To(BeNil())
		}

		add(9)

		g.Expect(ruleA.IsCalculated(start.Unix())).To(BeTrue())
		g.Expect(ruleCache.ruleNames[ruleA.Name()]).To(Equal(0))
		ok, err := ruleA.IsSatisfied(start.Unix())
		g.Expect(err).To(BeNil())
		g.Expect(ok).To(BeFalse())
		g.Expect(ruleA.IsCalculated(start.Unix())).To(BeTrue())

		add(11)
		g.Expect(ruleA.IsCalculated(start.Unix())).To(BeTrue())
		g.Expect(ruleA.IsSatisfied(start.Unix())).To(BeTrue())

	})
	t.Run("Caching some And Rules", func(t *testing.T) {

		t.SkipNow()
		g := NewGomegaWithT(t)

		history := NewPerformanceAskBidTradeHistory()
		ruleCache := NewRuleCache(history)
		indicatorCache := NewIndicatorCache(history)

		constantA := NewIndicatorConstant(10, indicatorCache)
		constantB := NewIndicatorConstant(20, indicatorCache)

		liveIndicator := NewIndicatorLive(AskClose, indicatorCache, history)

		ruleA := NewOverIndicatorRule(liveIndicator, constantA, 0.5, ruleCache)
		ruleB := NewOverIndicatorRule(liveIndicator, constantB, 0.5, ruleCache)
		ruleC := NewAndRule(ruleA, ruleB, ruleCache)

		start := time.Now()
		add := func(value float64) {

			start = start.Add(time.Minute)
			abd := &AskBidTrade{
				Ask:       Record{value, value, value, value},
				Bid:       Record{value, value, value, value},
				Trade:     Record{value, value, value, value},
				InputDate: start,
			}

			history.Append(abd)
			err := indicatorCache.Calculate(abd, true)
			g.Expect(err).To(BeNil())

			err = ruleCache.Calculate(start)
			g.Expect(err).To(BeNil())
		}

		add(9)

		sa, err := ruleA.IsSatisfied(start.Unix())
		g.Expect(err).To(BeNil())
		g.Expect(sa).To(BeFalse())
		sb, err := ruleB.IsSatisfied(start.Unix())
		g.Expect(err).To(BeNil())
		g.Expect(sb).To(BeFalse())
		sc, err := ruleC.IsSatisfied(start.Unix())
		g.Expect(err).To(BeNil())
		g.Expect(sc).To(BeFalse())

		add(11)

		sa, err = ruleA.IsSatisfied(start.Unix())
		g.Expect(err).To(BeNil())
		g.Expect(sa).To(BeTrue())
		sb, err = ruleB.IsSatisfied(start.Unix())
		g.Expect(err).To(BeNil())
		g.Expect(sb).To(BeFalse())
		sc, err = ruleC.IsSatisfied(start.Unix())
		g.Expect(err).To(BeNil())
		g.Expect(sc).To(BeFalse())

		add(21)
		sa, err = ruleA.IsSatisfied(start.Unix())
		g.Expect(err).To(BeNil())
		g.Expect(sa).To(BeTrue())
		sb, err = ruleB.IsSatisfied(start.Unix())
		g.Expect(err).To(BeNil())
		g.Expect(sb).To(BeTrue())
		sc, err = ruleC.IsSatisfied(start.Unix())
		g.Expect(err).To(BeNil())
		g.Expect(sc).To(BeTrue())

	})

}
