// Package quality provides practical quality management calculations.
//
// It includes:
//
//   - DPMO: defects per million opportunities
//   - Cp: process capability (spread vs. tolerance)
//   - Cpk: process capability index (spread and centering vs. tolerance)
//   - CostOfQuality: prevention, appraisal, and failure cost breakdown
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
package quality
