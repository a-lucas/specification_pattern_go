/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"strings"
)

type IndicatorCache struct {
	cache      map[string]IIndicator // Based on live values
	calculated map[string]int64
	history    *PerformanceAskBidTradeHistory
	triggerMap map[int][]string // map the triggerID with the corresponding indicators
}

func NewIndicatorCache(history *PerformanceAskBidTradeHistory) *IndicatorCache {
	return &IndicatorCache{
		cache:      make(map[string]IIndicator),
		calculated: make(map[string]int64),
		history:    history,
		triggerMap: make(map[int][]string),
	}
}

func (ic *IndicatorCache) GetFirstLive() IIndicator {
	for name, ind := range ic.cache {
		if strings.HasPrefix(name, "LIVE") {
			return ind
		}
	}
	panic("LIVE indicator not found")
}

func (ic *IndicatorCache) History() *PerformanceAskBidTradeHistory {
	return ic.history
}

func (ic *IndicatorCache) List() []string {
	list := make([]string, 0)
	for name := range ic.cache {
		list = append(list, name)
	}
	return list
}

func (ic *IndicatorCache) Get(name string) IIndicator {
	if ind, ok := ic.cache[name]; !ok {
		panic("indicator not defined")
	} else {
		return ind
	}
}

func (ic *IndicatorCache) Has(name string) bool {
	_, ok := ic.cache[name]
	return ok
}

func (ic *IndicatorCache) Set(indicator IIndicator) {
	if !ic.Has(indicator.Name()) {
		ic.cache[indicator.Name()] = indicator
		ic.calculated[indicator.Name()] = 0
	}
}

func (ic *IndicatorCache) Calculate(record *AskBidTrade, add bool) error {
	unix := record.InputDate.Unix()
	for _, indicator := range ic.cache {
		if err := indicator.Calculate(unix, add); err != nil {
			logrus.WithError(err).WithField("Record", spew.Sdump(record)).Error("keep calculating other indicators")
			return err
		}
	}
	return nil
}
