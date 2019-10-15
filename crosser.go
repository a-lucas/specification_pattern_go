/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

import (
	"errors"
	"github.com/sirupsen/logrus"
	"math"
)

type DiffData struct {
	Ratio float64
}

func NewDiffData(val1, val2 float64) DiffData {
	return DiffData{
		Ratio: (val1 - val2) / math.Max(val1, val2),
	}
}

func (d *DiffData) IsZero(zeroThreshold float64) bool {
	return math.Abs(d.Ratio) < zeroThreshold
}

func (d *DiffData) Sign(zeroThreshold float64) CrossPosition {
	if d.IsZero(zeroThreshold) {
		return CrossPositionZero
	}

	if d.Ratio > 0 {
		return CrossPositionOver // indicator is above Live
	} else {
		return CrossPositionUnder // indicator is under live
	}
}

type Crosser struct {
	threshold                float64 // threshold is the ratio - not the percentage
	crossPositionToDetect    CrossPosition
	currentPosition          CrossPosition
	lastDetectedIntersection CrossPosition
	pristine                 bool
}

/*
CrossPositionUnder => indicator1 cross indicator2 from above to under
CrossPositionOver => indicator1 cross indicator2 from under to over
*/

func NewCrosser(threshold float64, position CrossPosition) *Crosser {
	return &Crosser{
		threshold:             threshold,
		crossPositionToDetect: position,
		pristine:              true,
	}
}

func (c *Crosser) Calculate(val1, val2 float64) (bool, error) {
	if val1 == 0 && val2 == 0 {
		logrus.WithField("Param1", c.threshold).WithField("Pristine", c.pristine).Error("We got some 0 indicator data")
		return false, errors.New("val1 and val2 are zero")
	}
	diff := NewDiffData(val1, val2)

	currentSign := diff.Sign(c.threshold)

	if c.pristine {
		switch currentSign {
		case CrossPositionOver:
			c.pristine = false
			c.currentPosition = CrossPositionOver
			c.lastDetectedIntersection = CrossPositionOver
		case CrossPositionUnder:
			c.pristine = false
			c.currentPosition = CrossPositionUnder
			c.lastDetectedIntersection = CrossPositionUnder
		}
		return false, nil
	}

	switch currentSign {
	case CrossPositionZero:
		c.currentPosition = CrossPositionZero
		return false, nil
	case CrossPositionOver:
		switch c.currentPosition {
		case CrossPositionOver:
			c.currentPosition = CrossPositionOver
			return false, nil
		case CrossPositionUnder:
			c.lastDetectedIntersection = CrossPositionOver
			c.currentPosition = CrossPositionOver
			return c.lastDetectedIntersection == c.crossPositionToDetect, nil
		case CrossPositionZero:
			if c.lastDetectedIntersection == CrossPositionOver {
				// actually same side
				c.currentPosition = CrossPositionOver
				return false, nil
			} else {
				c.lastDetectedIntersection = CrossPositionOver
				c.currentPosition = CrossPositionOver
				return c.lastDetectedIntersection == c.crossPositionToDetect, nil
			}
		}
	case CrossPositionUnder:
		switch c.currentPosition {
		case CrossPositionUnder:
			c.currentPosition = CrossPositionUnder
			return false, nil
		case CrossPositionOver:
			c.lastDetectedIntersection = CrossPositionUnder
			c.currentPosition = CrossPositionUnder
			return c.lastDetectedIntersection == c.crossPositionToDetect, nil
		case CrossPositionZero:
			if c.lastDetectedIntersection == CrossPositionUnder {
				// actually same side
				c.currentPosition = CrossPositionUnder
				return false, nil
			} else {
				c.lastDetectedIntersection = CrossPositionUnder
				c.currentPosition = CrossPositionUnder
				return c.lastDetectedIntersection == c.crossPositionToDetect, nil
			}
		}
	}
	return false, errors.New("logixc error")

}
