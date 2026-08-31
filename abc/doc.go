// Package abc classifies a list of inventory items by value concentration
// (ABC / Pareto analysis) and by demand variability (XYZ analysis), to help
// prioritize which items deserve tighter inventory control.
//
// It includes:
//
//   - Classify: ABC classification by cumulative value contribution
//     (a generalized Pareto analysis — the cumulative-percent field it
//     returns is the Pareto curve itself, so there is no separate
//     "Pareto analysis" function)
//   - ClassifyVariability: XYZ classification by demand coefficient of
//     variation
//   - Combine: joins ABC and XYZ results by item ID into the classic
//     ABC-XYZ matrix
//
// Important assumptions:
//
//   - Classify computes cumulative percentages by internally sorting items
//     by value descending (ties broken by original input order, for
//     deterministic output), but returns results in the same order as the
//     input Items slice, so ABC and XYZ results for the same item set stay
//     index-aligned and easy to combine.
//   - Classify requires all item values to be non-negative and their sum to
//     be greater than zero (cumulative percentage is undefined for an
//     all-zero value set).
//   - ClassifyVariability requires MeanDemand to be strictly positive for
//     every item (the coefficient of variation, StdDevDemand/MeanDemand,
//     is undefined at MeanDemand = 0).
//   - Combine performs an inner join on ID: an item present in only one of
//     the two input slices is omitted from the result, not defaulted.
//   - This package does not perform time-series-based variability analysis
//     (e.g. deriving MeanDemand/StdDevDemand from a raw demand history —
//     see the sibling github.com/motah-fard/scmgo/forecast package for
//     that), and does not choose thresholds for you: AThreshold/BThreshold
//     and XThreshold/YThreshold are caller-supplied, not defaulted.
package abc
