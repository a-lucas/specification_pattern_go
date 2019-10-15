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

func TestIndicatorSMA(t *testing.T) {

	newRecord := func(val float64, d time.Time) *AskBidTrade {
		return &AskBidTrade{
			InputDate: d,
			Ask: Record{
				Close: val,
			},
		}
	}

	t.Run("Just basic sma - All Append", func(t *testing.T) {

		start := time.Now()
		history := NewPerformanceAskBidTradeHistory()
		cache := NewIndicatorCache(history)
		g := NewGomegaWithT(t)
		live := NewIndicatorLive(AskClose, cache, history)
		ind := NewIndicatorSMA(2, live, cache)

		add := func(value float64) {
			start = start.Add(time.Minute)
			history.Append(newRecord(value, start))
			err := ind.Calculate(start.Unix(), true)
			g.Expect(err).To(BeNil())
		}

		add(1)
		g.Expect(ind.Val()).To(BeEquivalentTo(1.0))

		add(2)
		g.Expect(ind.Val()).To(BeEquivalentTo(1.5))

		add(3)
		g.Expect(ind.Val()).To(BeEquivalentTo(2.5))

		add(4)
		g.Expect(ind.Val()).To(BeEquivalentTo(3.5))

	})

	t.Run("Just basic sma - All Append - then update", func(t *testing.T) {

		start := time.Now()
		history := NewPerformanceAskBidTradeHistory()
		cache := NewIndicatorCache(history)
		g := NewGomegaWithT(t)
		live := NewIndicatorLive(AskClose, cache, history)
		ind := NewIndicatorSMA(2, live, cache)

		add := func(value float64) {
			start = start.Add(time.Minute)
			history.Append(newRecord(value, start))
			err := ind.Calculate(start.Unix(), true)
			g.Expect(err).To(BeNil())
		}

		setCurrent := func(value float64) {
			start = start.Add(time.Minute)
			history.SetCurrent(newRecord(value, start))
			err := ind.Calculate(start.Unix(), false)
			g.Expect(err).To(BeNil())
		}

		add(1)
		g.Expect(ind.Val()).To(BeEquivalentTo(1))

		add(2)
		g.Expect(ind.Val()).To(BeEquivalentTo(1.5))

		add(3)
		g.Expect(ind.Val()).To(BeEquivalentTo(2.5))

		setCurrent(4)
		g.Expect(ind.Val()).To(BeEquivalentTo(3.5))

		setCurrent(3)
		g.Expect(ind.Val()).To(BeEquivalentTo(3))

		add(4)
		g.Expect(ind.Val()).To(BeEquivalentTo(3.5))

		setCurrent(4)
		g.Expect(ind.Val()).To(BeEquivalentTo(4))

		g.Expect(ind.Values()).To(BeEquivalentTo([]float64{1, 1.5, 2.5, 3.5}))

	})

	t.Run("2 SMA with 10 minutes per day", func(t *testing.T) {

		start := time.Now()
		history := NewPerformanceAskBidTradeHistory()
		cache := NewIndicatorCache(history)
		g := NewGomegaWithT(t)
		live := NewIndicatorLive(AskClose, cache, history)
		ind := NewIndicatorSMA(10, live, cache)

		add := func(value float64) {
			start = start.Add(time.Minute)
			history.Append(newRecord(value, start))
			err := ind.Calculate(start.Unix(), true)
			g.Expect(err).To(BeNil())
		}

		add(1)
		g.Expect(ind.Values()).To(Equal([]float64{1}))
		g.Expect(ind.Val()).To(BeEquivalentTo(1))

		add(2)
		g.Expect(ind.Values()).To(Equal([]float64{1, 1.5}))
		g.Expect(ind.Val()).To(Equal(1.5))

		add(3)
		g.Expect(ind.Values()).To(Equal([]float64{1, 1.5, 2}))
		g.Expect(ind.Val()).To(BeEquivalentTo(2))

		add(4)
		g.Expect(ind.Values()).To(Equal([]float64{1, 1.5, 2, 2.5}))
		g.Expect(ind.Val()).To(BeEquivalentTo(2.5))

		add(5)
		g.Expect(ind.Values()).To(Equal([]float64{1, 1.5, 2, 2.5, 3}))
		g.Expect(ind.Val()).To(BeEquivalentTo(3))

		add(6)
		g.Expect(ind.Values()).To(Equal([]float64{1, 1.5, 2, 2.5, 3, 3.5}))
		g.Expect(ind.Val()).To(BeEquivalentTo(3.5))

		add(7)
		g.Expect(ind.Values()).To(Equal([]float64{1, 1.5, 2, 2.5, 3, 3.5, 4}))
		g.Expect(ind.Val()).To(BeEquivalentTo(4))

		add(8)
		g.Expect(ind.Values()).To(Equal([]float64{1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5}))
		g.Expect(ind.Val()).To(BeEquivalentTo(4.5))

		add(9)
		g.Expect(ind.Values()).To(Equal([]float64{1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}))
		g.Expect(ind.Val()).To(BeEquivalentTo(5))

		add(10)
		g.Expect(ind.Values()).To(Equal([]float64{1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5}))
		g.Expect(ind.Val()).To(BeEquivalentTo(5.5))

		add(11)
		g.Expect(ind.Values()).To(Equal([]float64{1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5, 6.5}))
		g.Expect(ind.Val()).To(BeEquivalentTo(6.5))
	})
}
