package inventory

// fillRateBisectionBound is the search range for the service factor z. It
// covers unit normal loss values from about 1e-23 (z=10, an essentially
// unachievable-in-practice fill rate) down to -z (z=-10), which is far
// beyond any practically useful fill-rate target.
const fillRateBisectionBound = 10.0

// fillRateBisectionIterations is chosen so the search interval (initially
// 2*fillRateBisectionBound wide) shrinks well past float64 precision;
// bisection converges long before this many iterations run.
const fillRateBisectionIterations = 100

// FillRateSafetyStock calculates the safety stock required to achieve a
// target item fill rate (P2), by numerically inverting ExpectedFillRate's
// formula for z (the standard normal unit loss function, see
// UnitNormalLoss, is strictly decreasing and has no closed-form inverse):
//
//	target loss = (1 - TargetFillRate) × OrderQuantity / σL
//	solve L(z) = target loss for z via bisection
//	safety stock = z × σL
//
// TargetFillRate must be strictly between 0 and 1. StdDevDemandDuringLeadTime
// and OrderQuantity must be greater than zero. Returns ErrFillRateNotAchievable
// if the target cannot be reached within the supported search range (an
// extremely high target fill rate relative to a very small OrderQuantity, or
// an extremely low target relative to a very large one). The result can be
// negative for a low enough target: OrderQuantity alone may already exceed
// it without any additional buffer (see ExpectedFillRate).
func FillRateSafetyStock(in FillRateSafetyStockInput) (float64, error) {
	if err := validateFinite(in.TargetFillRate, in.StdDevDemandDuringLeadTime, in.OrderQuantity); err != nil {
		return 0, err
	}
	if in.TargetFillRate <= 0 || in.TargetFillRate >= 1 {
		return 0, ErrInvalidFillRate
	}
	if in.StdDevDemandDuringLeadTime <= 0 {
		return 0, ErrInvalidStandardDeviation
	}
	if in.OrderQuantity <= 0 {
		return 0, ErrInvalidOrderQuantity
	}

	targetLoss := (1 - in.TargetFillRate) * in.OrderQuantity / in.StdDevDemandDuringLeadTime

	lo, hi := -fillRateBisectionBound, fillRateBisectionBound
	lossAtLo, _ := UnitNormalLoss(lo)
	lossAtHi, _ := UnitNormalLoss(hi)
	if targetLoss > lossAtLo || targetLoss < lossAtHi {
		return 0, ErrFillRateNotAchievable
	}

	for i := 0; i < fillRateBisectionIterations; i++ {
		mid := (lo + hi) / 2
		lossAtMid, _ := UnitNormalLoss(mid)
		// UnitNormalLoss is strictly decreasing, so a loss still above
		// target means z needs to increase.
		if lossAtMid > targetLoss {
			lo = mid
		} else {
			hi = mid
		}
	}
	z := (lo + hi) / 2

	return z * in.StdDevDemandDuringLeadTime, nil
}
