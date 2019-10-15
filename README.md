# specification_pattern_go

The **Specification Pattern** is a pattern aimed for composing abstract conditions. 

Check this amazing read by Eric Evans and Martin Fowler  https://www.martinfowler.com/apsupp/spec.pdf which explains in details the advantages of the specification pattern over DDD and other architectural methods. 

You can also learn more on the wikipedia page (don't forget to donate) https://en.wikipedia.org/wiki/Specification_pattern

This small demo shows how to implement the specification pattern in Go with a technical analysis example.
As long as a struct implement the following interface, it can be reused and composed with other structs.

```go

type IRule interface {
	IsCalculated(date int64) bool
	IsSatisfied(date int64) (bool, error)
	And(rule IRule) IRule
	Or(rule IRule) IRule
	Xor(rule IRule) IRule
	Negation() IRule
	// simplified
}

```

This is not a package aimed at production - merely an example of the specification pattern implementation.

This is hugely inspired by the `ta4j` java library https://github.com/ta4j/ta4j  - which I took some of the resource testing data from.

The original code is proprietary (owned by me), and I stripped down the sensitive parts. - which is the reason why you still get remnants of caching implementation for high performance processing.

Technically the original code can be used this way: 

```go

// this is a dummy example

history := NewPerformanceAskBidTradeHistory()
cache := NewIndicatorCache(history)
ruleCache := NewRuleCache(history)

constant5 := NewIndicatorConstant(5, cache)
constant6 := NewIndicatorConstant(4, cache)

sma := NewIndicatorSMA(60, live, cache)

smaOver5 := NewOverIndicatorRule(sma, constant5, 0, rulecache)

smaOver6 := NewOverIndicatorRule(sma, constant6, 0, rulecache)
smaUnder6 := NewNotRule(smaOver64, ruleCache)

smaOver5AndSmaUnder6 := NewAndRule(smaOver5, smaUnder6, ruleCache)
smaOver5AndSmaUnder6.IsSatisfied(time.Now().Unix())

```

This code is certainly faster in some parts than the Java original: 

For example, the SMA implementation of ta4j makes the SUM divided by the length every time to calculate it. 
https://github.com/ta4j/ta4j/blob/master/ta4j-core/src/main/java/org/ta4j/core/indicators/SMAIndicator.java

```java
protected Num calculate(int index) {
    Num sum = getTimeSeries().numOf(0);
    for (int i = Math.max(0, index - barCount + 1); i <= index; i++) {
        sum = sum.plus(indicator.getValue(i));
    }

    final int realBarCount = Math.min(barCount, index + 1);
    return sum.dividedBy(getTimeSeries().numOf(realBarCount));
}
```

Instead my implementation uses the previous calculated value and do one addition and one subtraction instead. It is uglier - but faster. 


I don't know how much faster it is, because I don't bother benchmarking Java code - but it certainly is.