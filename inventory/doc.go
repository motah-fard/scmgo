// Package inventory provides practical inventory policy calculations for
// inventory control and supply chain planning.
//
// It includes:
//
//   - Reorder point
//   - Basic safety stock
//   - Economic order quantity (EOQ)
//   - Min/max inventory levels
//   - Z-score lookup for service levels
//   - Service-level-based safety stock
//   - Service-level-based reorder point
//   - Demand during lead time
//   - Standard deviation of demand during lead time
//   - Target inventory level
//   - Service-level-based target inventory level
//   - Service-level-based min/max inventory levels
//   - Deterministic policy summary helpers
//   - Service-level-based policy summary helpers
//   - Batch variants of the policy summary helpers for SKU lists
//   - Safety stock and reorder point accounting for variable lead time
//   - Fill-rate (item fill rate / P2) safety stock, both directions:
//     expected fill rate from a given safety stock, and the safety stock
//     required for a target fill rate
//   - Economic production quantity (EPQ)
//   - Economic order quantity across a quantity discount schedule
//   - Newsvendor (single-period) optimal order quantity
//   - Inventory turnover, days of inventory on hand, and GMROI
//   - Economic order interval (EOI) and safety time
//
// The formulas in this package are intentionally simple and transparent.
// They are designed for practical use in applications, internal tools,
// learning, and lightweight decision support.
//
// Important assumptions:
//
//   - Input units must be consistent.
//     For example, if demand is measured per day, lead time should also be in days.
//   - Safety stock is expressed in inventory units.
//   - All float64 inputs must be finite: NaN and +/-Inf are rejected with
//     ErrNonFiniteInput rather than silently propagating into the result.
//   - EOQ uses the classic Wilson EOQ formula.
//   - SafetyStockBasic uses a simple max/average demand and lead-time formula.
//   - SafetyStockWithServiceLevel assumes a normal approximation and combines
//     demand variability with a target cycle service level.
//   - ReorderPointWithServiceLevel combines expected lead-time demand with
//     service-level-based safety stock.
//   - StdDevDemandDuringLeadTime assumes independent demand variability
//     across lead-time periods.
//   - TargetInventoryLevel combines expected demand coverage with safety stock.
//   - Policy summary helpers combine lead-time demand, review-period demand,
//     safety stock, reorder point, target inventory level, and min/max outputs
//     into one higher-level result.
//   - SafetyStockWithVariableLeadTime accounts for variability in both
//     demand and lead time; it reduces to SafetyStockWithServiceLevel when
//     lead-time variability is zero.
//   - ExpectedFillRate and FillRateSafetyStock assume a normal
//     approximation of demand during lead time. FillRateSafetyStock's
//     result can be negative for a low enough target fill rate (see its
//     doc comment) -- this is not a bug, and ExpectedFillRate accepts a
//     negative safety stock accordingly.
//   - EOQWithQuantityDiscounts evaluates each price tier's own EOQ,
//     clamped to that tier's valid quantity range, and returns whichever
//     tier/quantity combination minimizes total annual cost.
//   - EOI requires AnnualDemand strictly greater than zero (unlike EOQ,
//     which allows zero), since EOI divides by it.
//   - This package does not include demand forecasting (see the sibling
//     github.com/motah-fard/scmgo/forecast package), stochastic
//     optimization, or multi-echelon inventory models (a guaranteed-service
//     multi-echelon model requires network topology and NP-hard
//     optimization over a DAG of stocking locations -- a fundamentally
//     different kind of problem from the closed-form/numerically-inverted
//     formulas in this package, not a planned extension of it).
package inventory
