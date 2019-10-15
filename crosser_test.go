/*
 * Copyright (c) 2019. Antoine LUCAS
 * All Rights reserved.
 */

package specification_pattern

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestCrosser(t *testing.T) {

	t.Run("basic crossing Under", func(t *testing.T) {

		g := NewGomegaWithT(t)

		crosser := NewCrosser(0.1, CrossPositionUnder)

		detected, err := crosser.Calculate(2, 1)
		g.Expect(err).To(BeNil())

		g.Expect(detected).To(BeFalse())

		//g.Expect(crosser.lastDiff).To(BeEquivalentTo(&DiffData{
		//	Ratio:      0.5,
		//	Percentage: 50,
		//}))

		g.Expect(crosser.currentPosition).To(Equal(CrossPositionOver))
		g.Expect(crosser.lastDetectedIntersection).To(Equal(CrossPositionOver))

		detected, err = crosser.Calculate(1, 2)
		g.Expect(err).To(BeNil())

		g.Expect(detected).To(BeTrue())

		//g.Expect(crosser.lastDiff).To(BeEquivalentTo(&DiffData{
		//	Ratio:      -0.5,
		//	Percentage: -50,
		//}))

		g.Expect(crosser.currentPosition).To(Equal(CrossPositionUnder))
		g.Expect(crosser.lastDetectedIntersection).To(Equal(CrossPositionUnder))
	})

	t.Run("basic crossing Over", func(t *testing.T) {

		g := NewGomegaWithT(t)

		crosser := NewCrosser(0.1, CrossPositionOver)

		detected, err := crosser.Calculate(1, 2)
		g.Expect(err).To(BeNil())
		g.Expect(detected).To(BeFalse())

		//g.Expect(crosser.lastDiff).To(BeEquivalentTo(&DiffData{
		//	Ratio:      -0.5,
		//	Percentage: -50,
		//}))

		g.Expect(crosser.currentPosition).To(Equal(CrossPositionUnder))
		g.Expect(crosser.lastDetectedIntersection).To(Equal(CrossPositionUnder))

		detected, err = crosser.Calculate(2, 1)
		g.Expect(err).To(BeNil())
		g.Expect(detected).To(BeTrue())

		//g.Expect(crosser.lastDiff).To(BeEquivalentTo(&DiffData{
		//	Ratio:      0.5,
		//	Percentage: 50,
		//}))

		g.Expect(crosser.currentPosition).To(Equal(CrossPositionOver))
		g.Expect(crosser.lastDetectedIntersection).To(Equal(CrossPositionOver))
	})

	t.Run("With a linear curve crossing a constant", func(t *testing.T) {
		g := NewGomegaWithT(t)

		crosser := NewCrosser(0.1, CrossPositionUnder)

		constant := 5.0
		linear := func(x float64) float64 {
			return x
		}

		detected, err := crosser.Calculate(constant, linear(1))
		g.Expect(err).To(BeNil())
		g.Expect(detected).To(BeFalse())
		detected, err = crosser.Calculate(constant, linear(4.9))
		g.Expect(err).To(BeNil())
		g.Expect(detected).To(BeFalse())
		detected, err = crosser.Calculate(constant, linear(5.1))
		g.Expect(err).To(BeNil())
		g.Expect(detected).To(BeFalse())
		detected, err = crosser.Calculate(constant, linear(6))
		g.Expect(err).To(BeNil())
		g.Expect(detected).To(BeTrue())
		detected, err = crosser.Calculate(constant, linear(7))
		g.Expect(err).To(BeNil())
		g.Expect(detected).To(BeFalse())
		detected, err = crosser.Calculate(constant, linear(8))
		g.Expect(err).To(BeNil())
		g.Expect(detected).To(BeFalse())
	})

}
