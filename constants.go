/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

import "strconv"

type CrossPosition string

const (
	CrossPositionOver  CrossPosition = "Over"
	CrossPositionUnder CrossPosition = "Under"
	CrossPositionZero  CrossPosition = "Zero"
)

type IndicatorType string

const (
	CONSTANT IndicatorType = "CONSTANT"
	BOOL     IndicatorType = "BOOL"
	SMA      IndicatorType = "SMA"
	EMA      IndicatorType = "EMA"
	ROC      IndicatorType = "ROC"
	LOSS     IndicatorType = "LOSS Average"
	OBV      IndicatorType = "OBV"
	RSI      IndicatorType = "RSI"
	LIVE     IndicatorType = "LIVE"
	EMPTY    IndicatorType = ""
	ADX      IndicatorType = "ADX"
	MMA      IndicatorType = "MMA"
)

type RecordDataSource int

const (
	AskLow   RecordDataSource = 1
	AskHigh  RecordDataSource = 2
	AskClose RecordDataSource = 3
	BidLow   RecordDataSource = 4
	BidHigh  RecordDataSource = 5
	BidClose RecordDataSource = 6
	Volume   RecordDataSource = 7
	Invalid  RecordDataSource = 8
)

type RecordSource int

const (
	Ask RecordSource = 1
	Bid RecordSource = 2
)

func (r RecordSource) String() string {
	if r == Ask {
		return "Ask"
	}
	return "Bid"
}

func (r RecordDataSource) String() string {
	switch r {
	case AskLow:
		return "AskLow"
	case AskHigh:
		return "AskHigh"
	case AskClose:
		return "AskClose"
	case BidLow:
		return "BidLow"
	case BidHigh:
		return "BidHigh"
	case BidClose:
		return "BidClose"
	case Volume:
		return "Volume"
	case Invalid:
		return "Invalid"
	default:
		panic("unknown record data source string" + strconv.Itoa(int(r)))
	}
}

func (r RecordDataSource) RecordSource() RecordSource {
	switch r {
	case AskLow, AskHigh, AskClose:
		return Ask
	case BidLow, BidHigh, BidClose:
		return Bid
	}
	panic("record source not implemented - and probably never will be")
}
