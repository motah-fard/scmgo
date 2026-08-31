// Package procurement provides practical purchasing and sourcing cost
// calculations.
//
// It includes:
//
//   - LandedCost: total cost of getting a purchased item to its
//     destination, from labeled cost components
//   - PurchasePriceVariance: the cost impact of paying a different price
//     than standard/budgeted
//   - TotalCostOfOwnership: total cost of owning an asset or item over its
//     lifecycle, from labeled cost components
//
// Important assumptions:
//
//   - LandedCost and TotalCostOfOwnership both sum a caller-supplied list
//     of labeled cost components rather than assuming a fixed set of
//     categories (freight, duty, insurance, ... for landed cost;
//     acquisition, operating, maintenance, ... for TCO) — what belongs in
//     each varies too much by industry and organization to hardcode.
//     A component's Amount may be negative (e.g. a rebate, trade-in
//     credit, or resale value reducing total cost).
//   - PurchasePriceVariance follows the standard-costing convention: a
//     positive result means a favorable variance (paid less than
//     standard), a negative result means unfavorable (paid more).
//   - This package does not select or negotiate suppliers, does not model
//     multi-supplier sourcing allocation, and does not include the
//     quantity-discount EOQ or backorder-EOQ variants (see
//     github.com/motah-fard/scmgo/inventory for those — they're inventory
//     lot-sizing decisions, not procurement cost calculations).
package procurement
