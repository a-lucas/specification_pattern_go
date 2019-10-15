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

func TestIndicatorCache(t *testing.T) {

	newRecord := func(val float64, d time.Time) *AskBidTrade {
		return &AskBidTrade{
			InputDate: d,
			Ask: Record{
				Close: val,
			},
		}
	}

	t.Run("Empty", func(t *testing.T) {
		g := NewGomegaWithT(t)

		history := NewPerformanceAskBidTradeHistory()
		cache := NewIndicatorCache(history)

		g.Expect(cache.Has("nothing")).To(BeFalse())

	})

	t.Run("Add some indicator", func(t *testing.T) {
		g := NewGomegaWithT(t)
		history := NewPerformanceAskBidTradeHistory()
		cache := NewIndicatorCache(history)

		live := NewIndicatorLive(AskClose, cache, history)
		sma := NewIndicatorSMA(10, live, cache)

		cached := cache.Get(sma.Name())
		g.Expect(cached).To(BeEquivalentTo(sma))

		tt := time.Now()
		history.Append(newRecord(2, tt))
		cache.Calculate(newRecord(2, tt), true)
		g.Expect(sma.Val()).To(BeEquivalentTo(2))
		g.Expect(sma.Values()).To(BeEquivalentTo([]float64{2}))

	})

}
