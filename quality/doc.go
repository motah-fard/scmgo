// Package quality provides practical quality management calculations.
//
// It includes:
//
//   - DPMO: defects per million opportunities
//   - Cp: process capability (spread vs. tolerance)
//   - Cpk: process capability index (spread and centering vs. tolerance)
//   - CostOfQuality: prevention, appraisal, and failure cost breakdown
//   - SigmaLevel: converts DPMO to a Six-Sigma process sigma level
//
// Important assumptions:
//
//   - Cp and Cpk assume a normal process distribution, the standard
//     assumption behind both indices.
//   - Cpk's Mean may fall outside [LSL, USL]; the result (which can be
//     negative in that case) reflects a genuinely off-center or
//     out-of-tolerance process rather than being rejected as invalid
//     input — that information is meaningful.
//   - CostOfQuality uses the standard four-category breakdown (prevention,
//     appraisal, internal failure, external failure) rather than a
//     caller-supplied component list like procurement's LandedCost/
//     TotalCostOfOwnership, because this categorization is close to
//     universal in quality management, unlike landed cost or TCO
//     categories which vary by industry.
//   - SigmaLevel requires DPMO strictly between 0 and 1,000,000: at either
//     boundary the resulting yield is exactly 1 or 0, and the inverse
//     normal CDF behind the conversion is undefined (an infinity) there.
//     It uses the conventional 1.5-sigma shift.
package quality
