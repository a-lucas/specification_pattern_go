/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

import (
	"math"
	"time"
)

type Record struct {
	Low   float64
	High  float64
	Open  float64
	Close float64
}

func (r Record) IsZero() bool {
	return r.Low == 0 || r.Close == 0 || r.High == 0 || r.Open == 0
}

func (r Record) Similar(record Record, percentage float64) bool {
	valid := func(a, b float64) bool {
		if a == b {
			return true
		}
		return 100*math.Abs(a-b)/math.Min(a, b) <= percentage
	}

	return valid(r.Close, record.Close) && valid(r.Low, record.Low) && valid(r.High, record.High)
	//return r.Close == record.Close && r.Low == record.Low && r.High == record.High
}

func NewEmptyRecord() Record {
	return Record{
		Open:  0,
		High:  0,
		Close: 0,
		Low:   0,
	}
}

func NewRecord(previousRecord Record) Record {
	return Record{
		Open:  previousRecord.Close,
		High:  previousRecord.Close,
		Close: previousRecord.Close,
		Low:   previousRecord.Close,
	}
}

type AskBidTrade struct {
	InputDate time.Time
	Volume    float64
	Trade     Record
	Ask       Record
	Bid       Record
}

func (f *AskBidTrade) GetBuyValue() float64 {
	return math.Sqrt(f.Ask.Close * f.Ask.Low)
}

func (f *AskBidTrade) GetSellValue() float64 {
	return math.Sqrt(f.Bid.Close * f.Bid.High)
}

func (f *AskBidTrade) HasZero() bool {
	return f.Ask.IsZero() || f.Bid.IsZero() || f.Trade.IsZero()
}
func (f *AskBidTrade) Similar(abd *AskBidTrade, percentage float64) bool {
	return f.Ask.Similar(abd.Ask, percentage) && f.Bid.Similar(abd.Bid, percentage)
}

func (f *AskBidTrade) PingTrade(transactionPrice float64, volume float64, inputDate time.Time) {
	f.InputDate = inputDate
	if f.Trade.IsZero() {
		// fmt.Println("IS ZERO", inputDate.Format(time.Stamp), "->", transactionPrice)
		f.Trade = Record{
			Open:  transactionPrice,
			High:  transactionPrice,
			Close: transactionPrice,
			Low:   transactionPrice,
		}
		f.Volume = volume
	} else {
		f.Trade.Close = transactionPrice
		f.Volume += volume
		if f.Trade.Low > transactionPrice {
			f.Trade.Low = transactionPrice
		}
		if f.Trade.High < transactionPrice {
			f.Trade.High = transactionPrice
		}
	}
}

func (f *AskBidTrade) PingAsk(askPrice float64, inputDate time.Time) {
	f.InputDate = inputDate
	if f.Ask.IsZero() {
		f.Ask = Record{
			Low:   askPrice,
			Close: askPrice,
			High:  askPrice,
			Open:  askPrice,
		}
	} else {
		f.Ask.Close = askPrice

		if f.Ask.Low > askPrice {
			f.Ask.Low = askPrice
		}
		if f.Ask.High < askPrice {
			f.Ask.High = askPrice
		}
	}
}

func (f *AskBidTrade) PingBid(bidPrice float64, inputDate time.Time) {
	f.InputDate = inputDate
	if f.Bid.IsZero() {
		f.Bid = Record{
			Low:   bidPrice,
			Close: bidPrice,
			High:  bidPrice,
			Open:  bidPrice,
		}
	} else {
		f.Bid.Close = bidPrice
		if f.Bid.Low > bidPrice {
			f.Bid.Low = bidPrice
		}
		if f.Bid.High < bidPrice {
			f.Bid.High = bidPrice
		}
	}
}

func NewAskBidTradeFromTrade(transactionsPrice float64, volume float64, inputDate time.Time) *AskBidTrade {
	return &AskBidTrade{
		Ask:       NewEmptyRecord(),
		Bid:       NewEmptyRecord(),
		InputDate: inputDate,
		Volume:    volume,
		Trade: Record{
			Close: transactionsPrice,
			High:  transactionsPrice,
			Low:   transactionsPrice,
			Open:  transactionsPrice,
		},
	}
}

func NewAskBidTradeFromAskBid(askPrice float64, bidPrice float64, inputDate time.Time) *AskBidTrade {
	return &AskBidTrade{
		Ask: Record{
			Open:  askPrice,
			Low:   askPrice,
			High:  askPrice,
			Close: askPrice,
		},
		Bid: Record{
			Close: bidPrice,
			High:  bidPrice,
			Low:   bidPrice,
			Open:  bidPrice,
		},
		InputDate: inputDate,
		Volume:    0,
		Trade:     NewEmptyRecord(),
	}
}

func (f *AskBidTrade) GetFromSource(source RecordDataSource) float64 {
	switch source {
	case AskLow:
		return f.Ask.Low
	case AskHigh:
		return f.Ask.High
	case AskClose:
		return f.Ask.Close
	case BidLow:
		return f.Bid.Low
	case BidHigh:
		return f.Bid.High
	case BidClose:
		return f.Bid.Close
	}
	panic("unknown property")
}
