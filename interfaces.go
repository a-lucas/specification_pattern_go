/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

type IIndicator interface {
	Calculate(date int64, add bool) error

	GetLastValues(nb int) []float64
	GetValue(index int) float64
	GetValues(fromIndex, toIndex int) []float64

	//IsCalculated(inputDate time.Time) bool

	IsCalculated(date int64) bool
	Len() int
	Name() string
	SetCache(cache *IndicatorCache)
	SetHistory(history *PerformanceAskBidTradeHistory)
	Type() IndicatorType
	Values() []float64
	Val() float64
}

type IRule interface {
	IsCalculated(date int64) bool
	IsSatisfied(date int64) (bool, error)
	And(rule IRule) IRule
	Or(rule IRule) IRule
	Xor(rule IRule) IRule
	Negation() IRule

	Index() int
	SetIndex(index int)
	Name() string
}
