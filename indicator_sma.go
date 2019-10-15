/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

type IndicatorSMA struct {
	*IndicatorCommon
	indicatorSource IIndicator
	nbPoints        int
	_lastTotal      float64
	_lastLength     int
	_lastAdd        bool
	_lastValue      float64
}

func NewIndicatorSMA(nbPoints float64, source IIndicator, cache *IndicatorCache) IIndicator {

	ind := &IndicatorSMA{
		indicatorSource: source,

		nbPoints:        int(nbPoints),
		IndicatorCommon: NewIndicatorCommon(SMA, PointToPeriod(int(nbPoints)), source.Name()),
		_lastTotal:      0,
		_lastAdd:        false,
		_lastValue:      0,
		_lastLength:     0,
	}
	if cache.Has(ind.Name()) {
		return cache.Get(ind.Name())
	}
	ind.SetCache(cache)
	ind.indicatorCache.Set(ind)
	return ind
}

func (sma *IndicatorSMA) calculateFromLast(newVal float64, add bool) (float64, error) {

	if !sma._lastAdd {
		// previously we didn't add, example of SMA=3
		// before 1, 2, [3, 4] , LastVal
		// after 1, 2, [3, 4] , newVal
		sma._lastTotal = sma._lastTotal - sma._lastValue + newVal
		sma._lastValue = newVal
		sma._lastAdd = add

		if add {
			if sma._lastLength < sma.nbPoints {
				sma._lastLength = sma._lastLength + 1
			}
		}
		return sma._lastTotal / float64(sma._lastLength), nil
	} else {
		//last time we added
		// before 1, [2, 3, 4] , LastVal
		// after 1, 2, [3, 4, 5] , newVal add=true
		// after 1, 2, [3, 4] , newVal add=false
		sma._lastAdd = add

		var total float64
		var length int

		if sma._lastLength < sma.nbPoints {

			sma._lastValue = newVal
			if add {
				sma._lastTotal = sma._lastTotal + newVal
				sma._lastLength = sma._lastLength + 1
				return sma._lastTotal / float64(sma._lastLength), nil
			} else {
				total = sma._lastTotal + newVal
				length = sma._lastLength + 1
				return total / float64(length), nil
			}
		} else {

			length = sma._lastLength
			littleHack := 0
			if add {
				littleHack = 1
			}
			index := sma.indicatorSource.Len() - littleHack - sma.nbPoints
			pointToRemove := sma.indicatorSource.GetValue(index)
			sma._lastTotal = sma._lastTotal - pointToRemove + newVal
		}
		sma._lastValue = newVal
		return sma._lastTotal / float64(sma._lastLength), nil

	}

}

func (sma *IndicatorSMA) Calculate(date int64, add bool) error {
	if sma.IsCalculated(date) {
		return nil
	}

	if err := sma.indicatorSource.Calculate(date, add); err != nil {
		return err
	}

	if smaVal, err := sma.calculateFromLast(sma.indicatorSource.Val(), add); err != nil {
		return err
	} else {
		sma.val = smaVal
		return sma.done(date, add)
	}
}
