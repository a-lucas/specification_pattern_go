/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

func NewIndicatorEMA(nbPoints float64, source IIndicator, cache *IndicatorCache) IIndicator {
	multiplier := 2.0 / (nbPoints + 1)
	return NewIndicatorAbstractEma(EMA, source, nbPoints, multiplier, cache)
}
