/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

import (
	"bufio"
	"encoding/csv"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"testing"
	"time"
)

func TestMMA(t *testing.T) {
	g := NewGomegaWithT(t)
	csvFile, _ := os.Open("./resources/MMA.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	lineNumber := 0
	records := make([]*AskBidTrade, 0)

	mmas := make([]float64, 0)
	for {
		line, error := reader.Read()
		lineNumber++
		if error == io.EOF {
			break
		} else if error != nil {
			g.Expect(error).To(BeNil())
			logrus.Fatal(error)
		} else {
			start := time.Now()
			if lineNumber > 7 {
				records = append(records, &AskBidTrade{
					Ask: Record{
						Open:  StrToFloat(line[1]),
						High:  StrToFloat(line[2]),
						Low:   StrToFloat(line[3]),
						Close: StrToFloat(line[4]),
					},
					Bid: Record{
						Open:  StrToFloat(line[1]),
						High:  StrToFloat(line[2]),
						Low:   StrToFloat(line[3]),
						Close: StrToFloat(line[4]),
					},
					Trade: Record{
						Open:  StrToFloat(line[1]),
						High:  StrToFloat(line[2]),
						Low:   StrToFloat(line[3]),
						Close: StrToFloat(line[4]),
					},
					InputDate: start.Add(time.Duration(lineNumber) * time.Minute),
				})
				mmas = append(mmas, StrToFloat(line[6]))
			}
		}
	}

	history := NewPerformanceAskBidTradeHistory()
	cache := NewIndicatorCache(history)

	liveIndicator := NewIndicatorLive(AskClose, cache, history)
	mmaIndicator := NewIndicatorMMA(liveIndicator, 13, cache)
	secondMMAIndicator := NewIndicatorMMA(liveIndicator, 13, cache)
	g.Expect(mmaIndicator).To(Equal(secondMMAIndicator))

	for index, abd := range records {

		history.Append(abd)
		err := cache.Calculate(abd, true)
		g.Expect(err).To(BeNil())
		g.Expect(mmaIndicator.Val()).To(BeNumerically("~", mmas[index], 0.00001))
	}

}
