// Package production provides practical manufacturing and production
// planning calculations.
//
// It includes:
//
//   - OEE: overall equipment effectiveness, from raw availability,
//     performance, and quality inputs
//   - WIPFromLittlesLaw, CycleTimeFromLittlesLaw, ThroughputFromLittlesLaw:
//     the three directions of Little's Law (WIP = Throughput × CycleTime)
//   - TaktTime: available production time per unit of customer demand
//
// Important assumptions:
//
//   - OEE computes Availability, Performance, and Quality from raw
//     inputs (planned production time, run time, ideal cycle time, total
//     count, good count) rather than taking them as precomputed ratios,
//     and returns all four so intermediate values are visible, not just
//     the final OEE score.
//   - OEE does not clamp its result to [0, 1]; a value above 1 (e.g. from
//     an overstated total count or understated ideal cycle time) is
//     returned as-is and signals a data quality issue in the caller's
//     inputs, not a library bug.
//   - Little's Law is exposed as three separate, explicitly named
//     functions (one per "solve for X") rather than one function with an
//     implicit "which field is unknown" parameter.
//   - This package does not include bill-of-materials (BOM) explosion or
//     lot-sizing algorithms (Wagner-Whitin, Silver-Meal, Part-Period
//     Balancing). Those are genuinely different problems — a BOM is a
//     graph-traversal problem (with cycle detection), and lot-sizing
//     algorithms are dynamic-programming/heuristic optimizations — not
//     closed-form formulas, and need more design and testing rigor than a
//     single formula before they can be trusted in this library.
package production
