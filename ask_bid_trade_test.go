/*
 * Copyright (c) 2018. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestFeeder(t *testing.T) {

	Convey("record", t, func() {
		Convey("newrecord()", func() {

			r := Record{
				Low:   1,
				Close: 2,
				High:  3,
				Open:  0.5,
			}

			r2 := NewRecord(r)
			So(r2.Open, ShouldEqual, r.Close)
			So(r2.Close, ShouldEqual, r.Close)
			So(r2.High, ShouldEqual, r.Close)
			So(r2.Low, ShouldEqual, r.Close)
		})
	})
	Convey("AskBidTrade", t, func() {

		f := &AskBidTrade{
			InputDate: time.Now(),
			Volume:    123,
			Ask: Record{
				Low:   1,
				High:  2,
				Close: 3,
				Open:  4,
			},
			Bid: Record{
				Low:   2,
				High:  3,
				Open:  4,
				Close: 5,
			},
			Trade: Record{
				Low:   4,
				High:  5,
				Close: 6,
				Open:  7,
			},
		}

		Convey("hasZero", func() {

			So(f.HasZero(), ShouldBeFalse)
			f.Ask.Low = 0
			So(f.HasZero(), ShouldBeTrue)
			f.Ask.Low = 1

			f.Bid.Low = 0
			So(f.HasZero(), ShouldBeTrue)
			f.Bid.Low = 2

			f.Trade.Low = 0
			So(f.HasZero(), ShouldBeTrue)
			f.Trade.Low = 4

		})

		Convey("NewAskBidTradeFromTrade", func() {
			t1 := time.Now()
			f1 := NewAskBidTradeFromTrade(10.5, 123, t1)
			So(f1.InputDate, ShouldEqual, t1)
			So(f1.Volume, ShouldEqual, 123)
			So(f1.Trade.High, ShouldEqual, 10.5)
			So(f1.Trade.Low, ShouldEqual, 10.5)
			So(f1.Trade.Open, ShouldEqual, 10.5)
			So(f1.Trade.Close, ShouldEqual, 10.5)
		})

		Convey("NewAskBidTradeFromTrade Wth Data", func() {
			t1 := time.Now()

			f1 := NewAskBidTradeFromTrade(10.5, 123, t1)

			So(f1.InputDate, ShouldEqual, t1)

			So(f1.Volume, ShouldEqual, 123)

			So(f1.Trade.High, ShouldEqual, 10.5)
			So(f1.Trade.Low, ShouldEqual, 10.5)
			So(f1.Trade.Open, ShouldEqual, 10.5)
			So(f1.Trade.Close, ShouldEqual, 10.5)

			So(f1.Bid.Close, ShouldEqual, 0)
			So(f1.Bid.Open, ShouldEqual, 0)
			So(f1.Bid.High, ShouldEqual, 0)
			So(f1.Bid.Low, ShouldEqual, 0)

			So(f1.Ask.Close, ShouldEqual, 0)
			So(f1.Ask.Open, ShouldEqual, 0)
			So(f1.Ask.High, ShouldEqual, 0)
			So(f1.Ask.Low, ShouldEqual, 0)
		})

		Convey("NewAskBidTradeFromAskBid", func() {

			t1 := time.Now()

			f1 := NewAskBidTradeFromAskBid(10, 11, t1)
			So(f1.InputDate, ShouldEqual, t1)
			So(f1.Bid.Open, ShouldEqual, 11)
			So(f1.Bid.Close, ShouldEqual, 11)
			So(f1.Bid.Low, ShouldEqual, 11)
			So(f1.Bid.High, ShouldEqual, 11)

			So(f1.Ask.Open, ShouldEqual, 10)
			So(f1.Ask.Close, ShouldEqual, 10)
			So(f1.Ask.Low, ShouldEqual, 10)
			So(f1.Ask.High, ShouldEqual, 10)

			So(f1.Volume, ShouldEqual, 0)
			So(f1.Trade.High, ShouldEqual, 0)
			So(f1.Trade.Low, ShouldEqual, 0)
			So(f1.Trade.Open, ShouldEqual, 0)
			So(f1.Trade.Close, ShouldEqual, 0)
		})

		Convey("PublishLiveTrade", func() {

			t1 := time.Now().Add(20 * time.Minute)

			f1 := AskBidTrade{}

			f1.PingTrade(10, 100, t1)

			So(f1.Volume, ShouldEqual, 100)
			So(f1.InputDate, ShouldEqual, t1)
			So(f1.Trade.Open, ShouldEqual, 10)
			So(f1.Trade.Close, ShouldEqual, 10)
			So(f1.Trade.High, ShouldEqual, 10)
			So(f1.Trade.Low, ShouldEqual, 10)

			f1.PingTrade(11, 100, t1.Add(1*time.Minute))

			So(f1.Volume, ShouldEqual, 200)
			So(f1.InputDate, ShouldEqual, t1.Add(1*time.Minute))
			So(f1.Trade.Open, ShouldEqual, 10)
			So(f1.Trade.Close, ShouldEqual, 11)
			So(f1.Trade.High, ShouldEqual, 11)
			So(f1.Trade.Low, ShouldEqual, 10)

			f1.PingTrade(9, 100, t1.Add(2*time.Minute))

			So(f1.Volume, ShouldEqual, 300)
			So(f1.InputDate, ShouldEqual, t1.Add(2*time.Minute))
			So(f1.Trade.Open, ShouldEqual, 10)
			So(f1.Trade.Close, ShouldEqual, 9)
			So(f1.Trade.High, ShouldEqual, 11)
			So(f1.Trade.Low, ShouldEqual, 9)

		})

		Convey("MemPingAsk", func() {
			t1 := time.Now().Add(10 * time.Minute)
			t2 := time.Now().Add(20 * time.Minute)
			t3 := time.Now().Add(25 * time.Minute)

			f1 := AskBidTrade{}

			f1.PingAsk(10, t1)

			So(f1.InputDate, ShouldEqual, t1)
			So(f1.Ask.Close, ShouldEqual, 10)
			So(f1.Ask.High, ShouldEqual, 10)
			So(f1.Ask.Low, ShouldEqual, 10)
			So(f1.Ask.Open, ShouldEqual, 10)

			f1.PingAsk(11, t2)

			So(f1.InputDate, ShouldEqual, t2)
			So(f1.Ask.Close, ShouldEqual, 11)
			So(f1.Ask.High, ShouldEqual, 11)
			So(f1.Ask.Low, ShouldEqual, 10)
			So(f1.Ask.Open, ShouldEqual, 10)

			f1.PingAsk(9, t3)

			So(f1.InputDate, ShouldEqual, t3)
			So(f1.Ask.Close, ShouldEqual, 9)
			So(f1.Ask.High, ShouldEqual, 11)
			So(f1.Ask.Low, ShouldEqual, 9)
			So(f1.Ask.Open, ShouldEqual, 10)

		})

		Convey("MemPingBid", func() {

			t1 := time.Now().Add(10 * time.Minute)
			t2 := time.Now().Add(20 * time.Minute)
			t3 := time.Now().Add(25 * time.Minute)

			f1 := AskBidTrade{}

			f1.PingBid(10, t1)

			So(f1.InputDate, ShouldEqual, t1)
			So(f1.Bid.Close, ShouldEqual, 10)
			So(f1.Bid.High, ShouldEqual, 10)
			So(f1.Bid.Low, ShouldEqual, 10)
			So(f1.Bid.Open, ShouldEqual, 10)

			f1.PingBid(11, t2)

			So(f1.InputDate, ShouldEqual, t2)
			So(f1.Bid.Close, ShouldEqual, 11)
			So(f1.Bid.High, ShouldEqual, 11)
			So(f1.Bid.Low, ShouldEqual, 10)
			So(f1.Bid.Open, ShouldEqual, 10)

			f1.PingBid(9, t3)

			So(f1.InputDate, ShouldEqual, t3)
			So(f1.Bid.Close, ShouldEqual, 9)
			So(f1.Bid.High, ShouldEqual, 11)
			So(f1.Bid.Low, ShouldEqual, 9)
			So(f1.Bid.Open, ShouldEqual, 10)
		})
	})
}
