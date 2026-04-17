// Package inventory provides practical inventory policy calculations for
// basic supply chain and stock management use cases.
//
// Version v0.2.0 includes:
//
//   - Reorder point
//   - Basic safety stock
//   - Economic order quantity (EOQ)
//   - Min/max inventory levels
//   - Z-score lookup for service levels
//   - Safety stock based on service level
//   - Reorder point based on service level
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
//   - This package does not include forecasting, stochastic optimization,
//     multi-echelon inventory models, or fill-rate-based service models.
package inventory
