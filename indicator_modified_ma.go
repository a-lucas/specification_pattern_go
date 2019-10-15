/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

/*
public class MMAIndicator extends AbstractEMAIndicator {
    private static final long serialVersionUID = -7287520945130507544L;
	public MMAIndicator(Indicator<Num> indicator, int barCount) {
		super(indicator, barCount, 1.0/barCount);
	}
}
*/

func NewIndicatorMMA(source IIndicator, barCount float64, cache *IndicatorCache) IIndicator {
	multiplier := 1.0 / barCount
	return NewIndicatorAbstractEma(MMA, source, barCount, multiplier, cache)
}
