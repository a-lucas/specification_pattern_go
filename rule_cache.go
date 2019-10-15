/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

import (
	"github.com/sirupsen/logrus"
	"time"
)

type SatisfactionStatus int

// Satisfaction cache
type RuleCache struct {
	rules               []IRule
	ruleNames           map[string]int
	history             *PerformanceAskBidTradeHistory
	satisfiedRules      []string
	satisfiedRulesIndex []int
	buyProcessors       [][]int
	sellProcessors      [][]int
}

//
//// Satisfaction cache
//type RuleCache struct {
//	rules           map[string]IRule
//	satisfied       map[string]bool
//	calculated      map[string]bool
//	emptyCalculated map[string]bool
//	ruleNames       []string
//	history         *models.PerformanceAskBidTradeHistory
//}

func NewRuleCache(history *PerformanceAskBidTradeHistory) *RuleCache {
	return &RuleCache{
		rules:               make([]IRule, 0),
		history:             history,
		ruleNames:           make(map[string]int),
		satisfiedRules:      make([]string, 0),
		satisfiedRulesIndex: make([]int, 0),
		buyProcessors:       make([][]int, 0),
		sellProcessors:      make([][]int, 0),
	}
}

func (rc *RuleCache) Len() int {
	return len(rc.ruleNames)
}

func (rc *RuleCache) has(name string) bool {
	_, ok := rc.ruleNames[name]
	return ok
}

func (rc *RuleCache) GetByName(name string) IRule {
	return rc.rules[rc.ruleNames[name]]
}

func (rc *RuleCache) GetById(index int) IRule {
	return rc.rules[index]
}

func (rc *RuleCache) IsSatisfied(date int64, index int) (bool, error) {
	return rc.rules[index].IsSatisfied(date)
}

func (rc *RuleCache) Set(rule IRule) IRule {
	name := rule.Name()
	if rc.has(name) {
		return rc.rules[rc.ruleNames[name]]
	}
	rc.rules = append(rc.rules, rule)
	index := len(rc.rules) - 1
	rc.ruleNames[name] = index
	rc.buyProcessors = append(rc.buyProcessors, make([]int, 0))
	rc.sellProcessors = append(rc.sellProcessors, make([]int, 0))
	rule.SetIndex(index)
	return rule
}

func (rc *RuleCache) AddBuyProcessor(ruleIndex int, processorIndex int) {
	rc.buyProcessors[ruleIndex] = append(rc.buyProcessors[ruleIndex], processorIndex)
}

func (rc *RuleCache) AddSellProcessor(ruleIndex int, processorIndex int) {
	rc.sellProcessors[ruleIndex] = append(rc.sellProcessors[ruleIndex], processorIndex)
}

func (rc *RuleCache) calculate(date int64, index int) (bool, error) {
	return rc.rules[index].IsSatisfied(date)
}

func (rc *RuleCache) Calculate(date time.Time) error {
	unix := date.Unix()
	rc.satisfiedRules = make([]string, 0)
	rc.satisfiedRulesIndex = make([]int, 0)
	for index := 0; index < len(rc.rules); index++ {
		if satisfied, err := rc.calculate(unix, index); err != nil {
			logrus.WithField("History Length", rc.history.Len()).WithField("InputDate", rc.history.InputDate[rc.history.Len()-1].String()).WithError(err).WithField("name", rc.rules[index].Name()).Error("error calculating rule")
			//panic(err)
		} else {
			if satisfied {
				rc.satisfiedRules = append(rc.satisfiedRules, rc.rules[index].Name())
				rc.satisfiedRulesIndex = append(rc.satisfiedRulesIndex, index)
			}
		}
	}
	return nil
}

func (rc *RuleCache) GetSatisfiedRules() []string {
	return rc.satisfiedRules
}

func (rc *RuleCache) GetSatisfiedProcessors() ([]int, []int) {
	uniqueBuys := make(map[int]struct{})
	uniqueSells := make(map[int]struct{})

	for _, ruleIndex := range rc.satisfiedRulesIndex {
		for _, processorIndex := range rc.buyProcessors[ruleIndex] {
			uniqueBuys[processorIndex] = struct{}{}
		}
		for _, processorIndex := range rc.sellProcessors[ruleIndex] {
			uniqueSells[processorIndex] = struct{}{}
		}
	}
	buys := make([]int, 0)
	sells := make([]int, 0)
	for b := range uniqueBuys {
		buys = append(buys, b)
	}
	for s := range uniqueSells {
		sells = append(sells, s)
	}
	return buys, sells
}
