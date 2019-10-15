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

func TestIndicatorConstant(t *testing.T) {

	newRecord := func(val float64, d time.Time) *AskBidTrade {
		return &AskBidTrade{
			InputDate: d,
			Ask: Record{
				Close: val,
			},
		}
	}

	t.Run("Just basic - All Append", func(t *testing.T) {

		history := NewPerformanceAskBidTradeHistory()

		cache := NewIndicatorCache(history)

		g := NewGomegaWithT(t)
		ind := NewIndicatorConstant(1.5, cache)
		ind.SetHistory(history)

		start := time.Now()
		history.Append(newRecord(1, start))
		ind.Calculate(start.Unix(), true)

		g.Expect(ind.Val()).To(BeEquivalentTo(1.5))
		g.Expect(ind.Values()).To(BeEquivalentTo([]float64{1.5}))

		start = start.Add(time.Second)
		history.Append(newRecord(2, start))
		ind.Calculate(start.Unix(), true)

		g.Expect(ind.Val()).To(BeEquivalentTo(1.5))
		g.Expect(ind.Values()).To(BeEquivalentTo([]float64{1.5, 1.5}))

		start = start.Add(time.Second)
		history.Append(newRecord(2, start))
		ind.Calculate(start.Unix(), true)

		g.Expect(ind.Val()).To(BeEquivalentTo(1.5))
		g.Expect(ind.Values()).To(BeEquivalentTo([]float64{1.5, 1.5, 1.5}))

	})

	t.Run("Just basic sma - All Append - then update", func(t *testing.T) {
		g := NewGomegaWithT(t)

		history := NewPerformanceAskBidTradeHistory()
		cache := NewIndicatorCache(history)

		ind := NewIndicatorConstant(1.5, cache)

		start := time.Now()
		add := func(value float64, add bool) {
			start = start.Add(time.Minute)
			record := newRecord(value, start)
			if add {
				history.Append(record)
			} else {
				history.SetCurrent(record)
			}
			g.Expect(ind.Calculate(start.Unix(), add)).To(BeNil())

		}

		add(5, false)
		g.Expect(ind.Values()).To(BeEquivalentTo([]float64{}))
		g.Expect(ind.Val()).To(BeEquivalentTo(1.5))

		add(15, false)
		g.Expect(ind.Values()).To(BeEquivalentTo([]float64{}))
		g.Expect(ind.Val()).To(BeEquivalentTo(1.5))

		add(15, true)
		g.Expect(ind.Values()).To(BeEquivalentTo([]float64{1.5}))
		g.Expect(ind.Val()).To(BeEquivalentTo(1.5))
	})
}
