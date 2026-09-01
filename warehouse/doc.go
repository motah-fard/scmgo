// Package warehouse provides practical warehouse operations calculations.
//
// It includes:
//
//   - StorageUtilization: used vs. total storage locations/floor space
//   - CubeUtilization: used vs. total storage volume
//   - PickRate: picks completed per hour
//
// Important assumptions:
//
//   - StorageUtilization and CubeUtilization are not clamped to [0, 1]: a
//     value above 1 indicates overcommitted storage, which is meaningful
//     information, not an error.
//   - This package does not include velocity-based ABC slotting as a
//     separate function: that's exactly what
//     github.com/motah-fard/scmgo/abc's Classify function already does —
//     classify SKUs by pick velocity (or any other value metric) instead
//     of usage value, and use the resulting A/B/C classes to decide
//     slotting priority. A duplicate function here would just be the same
//     formula under a different name.
package warehouse
