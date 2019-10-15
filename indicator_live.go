/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

type IndicatorLive struct {
	source RecordDataSource
	*IndicatorCommon
}

func NewIndicatorLive(source RecordDataSource, cache *IndicatorCache, history *PerformanceAskBidTradeHistory) IIndicator {
	ind := &IndicatorLive{
		source:          source,
		IndicatorCommon: NewIndicatorCommon(LIVE, "", source.String()),
	}
	if cache.Has(ind.Name()) {
		return cache.Get(ind.Name())
	}
	ind.SetCache(cache)
	ind.SetHistory(history)
	ind.indicatorCache.Set(ind)
	return ind
}

func (sma *IndicatorLive) GetRecordSource() RecordSource {
	switch sma.source {
	case AskLow:
		return Ask
	case AskHigh:
		return Ask
	case AskClose:
		return Ask
	case BidLow:
		return Bid
	case BidHigh:
		return Bid
	case BidClose:
		return Bid
	default:
		panic("not handled")
	}
}

func (sma *IndicatorLive) Calculate(date int64, add bool) error {
	if sma.IsCalculated(date) {
		return nil
	}
	sma.val = sma.history.Current().GetFromSource(sma.source)
	return sma.done(date, add)
}
