// Package finance provides practical supply-chain-relevant financial
// calculations.
//
// It includes:
//
//   - DSO: days sales outstanding
//   - DPO: days payable outstanding
//   - CashToCashCycleTime: DIO + DSO - DPO
//   - PerfectOrderRate: the probability an order is on-time, complete,
//     damage-free, and accurately documented
//
// Important assumptions:
//
//   - CashToCashCycleTime takes DIO (days inventory outstanding), DSO, and
//     DPO as direct inputs rather than recomputing them internally, so a
//     caller can compute DIO however they prefer — including with
//     github.com/motah-fard/scmgo/inventory's DaysOfInventoryOnHand — and
//     pass in DSO/DPO from this package.
//   - PerfectOrderRate multiplies its four component rates, which assumes
//     they are independent. Real-world failure modes are often correlated
//     (e.g. a late order is also more likely to arrive damaged), so this
//     is a standard simplification, not a claim of statistical
//     independence in your actual operation — treat the result as an
//     upper-bound estimate if you suspect correlation.
package finance
