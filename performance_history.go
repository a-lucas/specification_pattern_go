/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

import (
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type PerformanceAskBidTradeHistory struct {
	Volumes      []float64 // 0
	TradeLow     []float64 // 1
	TradeHigh    []float64 // 2
	TradeOpen    []float64 // 3
	TradeClose   []float64 // 4
	AskLow       []float64 // 5
	AskHigh      []float64 // 6
	AskOpen      []float64 // 7
	AskClose     []float64 // 8
	BidLow       []float64 // 9
	BidHigh      []float64 // 10
	BidOpen      []float64 // 11
	BidClose     []float64 // 12
	InputDate    []time.Time
	InputDateStr []string
	len          int
	// locationConfig *LocationConfig
	current *AskBidTrade
}

func NewPerformanceAskBidTradeHistory() *PerformanceAskBidTradeHistory {
	p := &PerformanceAskBidTradeHistory{
		Volumes:      make([]float64, 0),
		AskClose:     make([]float64, 0),
		AskHigh:      make([]float64, 0),
		AskLow:       make([]float64, 0),
		AskOpen:      make([]float64, 0),
		BidClose:     make([]float64, 0),
		BidHigh:      make([]float64, 0),
		BidLow:       make([]float64, 0),
		BidOpen:      make([]float64, 0),
		TradeClose:   make([]float64, 0),
		TradeHigh:    make([]float64, 0),
		TradeLow:     make([]float64, 0),
		TradeOpen:    make([]float64, 0),
		InputDate:    make([]time.Time, 0),
		InputDateStr: make([]string, 0),
		len:          0,
	}
	//ny, _ := GetNewYorkConfig()
	//p.locationConfig = ny
	return p
}
func (p *PerformanceAskBidTradeHistory) Len() int {
	return p.len
}

func (p *PerformanceAskBidTradeHistory) SetCurrent(record *AskBidTrade) {
	p.current = record
}

func (p *PerformanceAskBidTradeHistory) Current() *AskBidTrade {
	return p.current
}

func (p *PerformanceAskBidTradeHistory) Append(record *AskBidTrade) {

	p.current = record
	//inputDate := record.InputDate.In(p.locationConfig.Location)
	p.TradeOpen = append(p.TradeOpen, record.Trade.Open)
	p.TradeLow = append(p.TradeLow, record.Trade.Low)
	p.TradeHigh = append(p.TradeHigh, record.Trade.High)
	p.TradeClose = append(p.TradeClose, record.Trade.Close)

	p.AskOpen = append(p.AskOpen, record.Ask.Open)
	p.AskLow = append(p.AskLow, record.Ask.Low)
	p.AskHigh = append(p.AskHigh, record.Ask.High)
	p.AskClose = append(p.AskClose, record.Ask.Close)

	p.BidOpen = append(p.BidOpen, record.Bid.Open)
	p.BidLow = append(p.BidLow, record.Bid.Low)
	p.BidHigh = append(p.BidHigh, record.Bid.High)
	p.BidClose = append(p.BidClose, record.Bid.Close)

	p.Volumes = append(p.Volumes, record.Volume)

	//p.InputDate = append(p.InputDate, inputDate)
	//p.InputDateStr = append(p.InputDateStr, inputDate.String())
	p.len++
}

func (p *PerformanceAskBidTradeHistory) IncludeCurrent(record *AskBidTrade) *PerformanceAskBidTradeHistory {
	newPerf := &PerformanceAskBidTradeHistory{
		InputDate:    p.InputDate,
		BidClose:     p.BidClose,
		BidOpen:      p.BidOpen,
		BidHigh:      p.BidHigh,
		BidLow:       p.BidLow,
		AskClose:     p.AskClose,
		AskOpen:      p.AskOpen,
		AskHigh:      p.AskHigh,
		AskLow:       p.AskLow,
		TradeClose:   p.TradeClose,
		TradeOpen:    p.TradeOpen,
		TradeHigh:    p.TradeHigh,
		TradeLow:     p.TradeLow,
		Volumes:      p.Volumes,
		len:          p.len,
		InputDateStr: p.InputDateStr,
	}
	newPerf.Append(record)
	return newPerf
}

func (p *PerformanceAskBidTradeHistory) GetPoint(index int) (*AskBidTrade, error) {
	if index < 0 {
		err := errors.New("CRITICAL - Getting negative Point")
		logrus.WithField("Len", p.len).WithError(err).WithField("index asked", index).Error()
		return nil, err
	}
	if index == 0 && p.len == 0 {
		err := errors.New("CRITICAL - Getting index=0 on empty PerformanceAskBidTradeHistory")
		logrus.WithError(err).Error()
		return nil, err
	}
	if p.len-1 < index {
		return nil, errors.New("getting invalid point - inoutDate length asked: " + strconv.Itoa(index) + ", maximal size: " + strconv.Itoa(len(p.InputDate)))
	}
	if p.len-1 == index {
		logrus.Warn("getting the exact last point")
	}

	defer func() {
		if err := recover(); err != nil {
			logrus.WithField("Error", spew.Sdump(err)).WithField("index", index).WithField("Len", p.len).Error("CRITICAL - panic recovered")
		}
	}()

	return &AskBidTrade{
		Bid: Record{
			Open:  p.BidOpen[index],
			Close: p.BidClose[index],
			High:  p.BidHigh[index],
			Low:   p.BidLow[index],
		},
		Ask: Record{
			Open:  p.AskOpen[index],
			Close: p.AskClose[index],
			High:  p.AskHigh[index],
			Low:   p.AskLow[index],
		},
		Trade: Record{
			Open:  p.TradeOpen[index],
			Close: p.TradeClose[index],
			High:  p.TradeHigh[index],
			Low:   p.TradeLow[index],
		},
		Volume:    p.Volumes[index],
		InputDate: p.InputDate[index],
	}, nil
}
