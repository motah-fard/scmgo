package inventory

import "math"

// EOQWithQuantityDiscounts finds the cost-minimizing order quantity across
// a quantity discount (price break) schedule.
//
// For each tier, it computes the tier's own EOQ using that tier's holding
// cost (HoldingCostRate × tier price), clamps it to the tier's valid
// quantity range (so a tier's EOQ that falls outside its own range is
// replaced by the nearest boundary of that range — usually the tier's
// minimum quantity), and evaluates the total annual cost
// (purchase + ordering + holding) at that clamped quantity. It returns the
// tier/quantity combination with the lowest total annual cost.
//
// AnnualDemand and OrderingCost must be non-negative. HoldingCostRate must
// be greater than zero. Tiers must contain at least one entry, sorted by
// strictly ascending MinQuantity with the first tier's MinQuantity greater
// than zero, and positive, non-increasing UnitPrice across tiers.
func EOQWithQuantityDiscounts(in EOQWithQuantityDiscountsInput) (EOQWithQuantityDiscountsResult, error) {
	if err := validateFinite(in.AnnualDemand, in.OrderingCost, in.HoldingCostRate); err != nil {
		return EOQWithQuantityDiscountsResult{}, err
	}
	if in.AnnualDemand < 0 {
		return EOQWithQuantityDiscountsResult{}, ErrNegativeDemand
	}
	if in.OrderingCost < 0 {
		return EOQWithQuantityDiscountsResult{}, ErrNegativeOrderingCost
	}
	if in.HoldingCostRate <= 0 {
		return EOQWithQuantityDiscountsResult{}, ErrInvalidHoldingCost
	}
	if err := validateDiscountTiers(in.Tiers); err != nil {
		return EOQWithQuantityDiscountsResult{}, err
	}

	best := EOQWithQuantityDiscountsResult{TotalAnnualCost: math.Inf(1)}

	for i, tier := range in.Tiers {
		upperBound := math.Inf(1)
		if i+1 < len(in.Tiers) {
			upperBound = in.Tiers[i+1].MinQuantity
		}

		holdingCostPerUnit := in.HoldingCostRate * tier.UnitPrice
		tierEOQ := math.Sqrt((2 * in.AnnualDemand * in.OrderingCost) / holdingCostPerUnit)
		quantity := clampFloat(tierEOQ, tier.MinQuantity, upperBound)

		totalCost := in.AnnualDemand*tier.UnitPrice +
			(in.AnnualDemand/quantity)*in.OrderingCost +
			(quantity/2)*holdingCostPerUnit

		if totalCost < best.TotalAnnualCost {
			best = EOQWithQuantityDiscountsResult{
				OrderQuantity:   quantity,
				UnitPrice:       tier.UnitPrice,
				TotalAnnualCost: totalCost,
			}
		}
	}

	return best, nil
}

func validateDiscountTiers(tiers []QuantityDiscountTier) error {
	if len(tiers) == 0 {
		return ErrEmptyDiscountTiers
	}

	prevMinQuantity := 0.0
	prevPrice := math.Inf(1)
	for i, tier := range tiers {
		if err := validateFinite(tier.MinQuantity, tier.UnitPrice); err != nil {
			return err
		}
		if tier.UnitPrice <= 0 {
			return ErrInvalidDiscountTiers
		}
		if i == 0 && tier.MinQuantity <= 0 {
			return ErrInvalidDiscountTiers
		}
		if tier.MinQuantity <= prevMinQuantity && i > 0 {
			return ErrInvalidDiscountTiers
		}
		if tier.UnitPrice > prevPrice {
			return ErrInvalidDiscountTiers
		}
		prevMinQuantity = tier.MinQuantity
		prevPrice = tier.UnitPrice
	}

	return nil
}

// clampFloat restricts v to the closed interval [lo, hi].
func clampFloat(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
