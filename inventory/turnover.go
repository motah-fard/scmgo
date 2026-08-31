package inventory

// Turnover calculates the inventory turnover ratio using:
//
//	inventory turnover = COGS / average inventory value
//
// COGS (cost of goods sold, over the same period as the average inventory
// value) must be non-negative. AverageInventoryValue must be greater than
// zero.
func Turnover(in TurnoverInput) (float64, error) {
	if err := validateFinite(in.COGS, in.AverageInventoryValue); err != nil {
		return 0, err
	}
	if in.COGS < 0 {
		return 0, ErrNegativeCOGS
	}
	if in.AverageInventoryValue <= 0 {
		return 0, ErrInvalidAverageInventoryValue
	}

	return in.COGS / in.AverageInventoryValue, nil
}
