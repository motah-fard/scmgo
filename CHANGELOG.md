# Changelog

All notable changes to this project will be documented in this file.

## [v0.5.0] - 2026-04-18
### Added
- `PolicySummary` for consolidated inventory policy outputs
- `PolicySummaryInput` for deterministic policy summary calculations
- `PolicySummaryServiceLevelInput` for service-level-driven policy summary calculations
- `BuildPolicySummary` for computing expected lead-time demand, safety stock, reorder point, target inventory level, min level, and max level in one call
- `BuildPolicySummaryWithServiceLevel` for computing policy summaries using service-level-based reorder point logic
- examples for deterministic and service-level policy summary workflows
- test coverage for policy summary builders

### Improved
- shared validation helpers for policy summary inputs
- cleaner internal assembly logic for policy summary outputs
- README updated for `v0.5.0` scope and summary-helper usage
- API consistency for higher-level inventory planning workflows

## [v0.4.0] - 2026-04-16
### Added
- `TargetInventoryLevel` for calculating target inventory level from expected demand and safety stock
- `TargetInventoryLevelWithServiceLevel` for calculating target inventory level using service-level-based safety stock
- `MinMaxLevelsWithServiceLevel` for calculating min/max inventory levels using a service-level-based reorder point

### Improved
- reused existing helpers to compose higher-level inventory policy functions
- expanded examples and documentation for target inventory and policy calculations

## [v0.3.0] - 2026-04-16
### Added
- `DemandDuringLeadTime` for calculating expected demand over lead time
- `StdDevDemandDuringLeadTime` for calculating demand variability over lead time

### Improved
- expanded inventory policy support with lead-time demand building blocks
- improved composability for higher-level inventory calculations
- added examples and tests for lead-time demand helpers

## [v0.2.0] - 2026-04-16
### Added
- `ZScoreForServiceLevel` for converting target service levels to z-scores
- `SafetyStockWithServiceLevel` for probabilistic safety stock calculations
- `ReorderPointWithServiceLevel` for reorder point calculations using service level targets

### Improved
- validation coverage for service-level-based inventory functions
- clearer error handling for invalid service level and negative standard deviation inputs

## [v0.1.2] - 2026-04-14
### Fixed
- updated license copyright name

## [v0.1.1] - 2026-04-14
### Added
- included the `inventory` package correctly in the tagged release
- added edge-case test coverage

## [v0.1.0] - 2026-04-14
### Added
- initial release of `scmgo`
- added `inventory` package
- added `ReorderPoint`
- added `SafetyStockBasic`
- added `EOQ`
- added `MinMaxLevels`
- added unit tests and examples
- added README and MIT license