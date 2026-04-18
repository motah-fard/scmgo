// Package inventory provides practical inventory policy calculations for
// basic supply chain and stock management use cases.
//
// Version v0.5.0 includes:
//
//   - Reorder point
//   - Basic safety stock
//   - Economic order quantity (EOQ)
//   - Min/max inventory levels
//   - Z-score lookup for service levels
//   - Safety stock based on service level
//   - Reorder point based on service level
//   - Demand during lead time
//   - Standard deviation of demand during lead time
//   - Target inventory level
//   - Target inventory level based on service level
//   - Min/max levels based on service level
//   - Deterministic policy summary helpers
//   - Service-level-based policy summary helpers
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
//   - EOQ uses the classic Wilson EOQ formula.
//   - SafetyStockBasic uses a simple max/average demand and lead-time formula.
//   - SafetyStockWithServiceLevel assumes a normal approximation and uses
//     demand variability together with a target cycle service level.
//   - ReorderPointWithServiceLevel combines expected lead-time demand with
//     service-level-based safety stock.
//   - StdDevDemandDuringLeadTime assumes independent demand variability
//     across lead-time periods.
//   - TargetInventoryLevel combines expected demand coverage with safety stock.
//   - Policy summary helpers combine lead-time demand, review-period demand,
//     safety stock, reorder point, target inventory level, and min/max outputs
//     into one higher-level result.
//   - This package does not include forecasting, stochastic optimization,
//     multi-echelon inventory models, or fill-rate-based service models.
package inventory
