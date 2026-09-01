package forecast

// Syntetos & Boylan (2005) classification cutoffs.
const (
	demandClassificationADIThreshold = 1.32
	demandClassificationCV2Threshold = 0.49
)

// ClassifyDemandPattern classifies a demand series using the Syntetos &
// Boylan (2005) framework, which determines which forecasting method is
// appropriate:
//
//	ADI = periods / non-zero periods
//	CV² = (StdDev / Mean)², computed over non-zero demand sizes only
//
//	ADI < 1.32, CV² < 0.49  -> smooth       (simple methods work well)
//	ADI >= 1.32, CV² < 0.49 -> intermittent (Croston or SBA)
//	ADI < 1.32, CV² >= 0.49 -> erratic      (high variability, low sparsity)
//	ADI >= 1.32, CV² >= 0.49 -> lumpy       (Croston or SBA; hardest to forecast)
//
// StdDev and Mean are computed as population statistics (dividing by
// count, not count-1), so a single non-zero period is well-defined (CV² =
// 0) rather than undefined.
//
// History must contain at least one non-zero period.
func ClassifyDemandPattern(in DemandClassificationInput) (DemandClassificationResult, error) {
	if len(in.History) == 0 {
		return DemandClassificationResult{}, ErrEmptyHistory
	}
	if err := validateNonNegative(in.History); err != nil {
		return DemandClassificationResult{}, err
	}

	var nonZero []float64
	for _, v := range in.History {
		if v != 0 {
			nonZero = append(nonZero, v)
		}
	}
	if len(nonZero) == 0 {
		return DemandClassificationResult{}, ErrNoNonZeroDemand
	}

	adi := float64(len(in.History)) / float64(len(nonZero))

	var sum float64
	for _, v := range nonZero {
		sum += v
	}
	mean := sum / float64(len(nonZero))

	var sumSquaredDeviation float64
	for _, v := range nonZero {
		d := v - mean
		sumSquaredDeviation += d * d
	}
	variance := sumSquaredDeviation / float64(len(nonZero))
	cv2 := variance / (mean * mean)

	class := "lumpy"
	switch {
	case adi < demandClassificationADIThreshold && cv2 < demandClassificationCV2Threshold:
		class = "smooth"
	case adi >= demandClassificationADIThreshold && cv2 < demandClassificationCV2Threshold:
		class = "intermittent"
	case adi < demandClassificationADIThreshold && cv2 >= demandClassificationCV2Threshold:
		class = "erratic"
	}

	return DemandClassificationResult{
		ADI:   adi,
		CV2:   cv2,
		Class: class,
	}, nil
}
