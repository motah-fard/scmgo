# Changelog

All notable changes to this project will be documented in this file.

## [v0.1.2] - 2026-04-14
### Fixed
- Updated license copyright name

## [v0.1.1] - 2026-04-14
### Added
- Included the `inventory` package correctly in the tagged release
- Added edge-case test coverage

## [v0.1.0] - 2026-04-14
### Added
- Initial release of `scmgo`
- Added `inventory` package
- Added `ReorderPoint`
- Added `SafetyStockBasic`
- Added `EOQ`
- Added `MinMaxLevels`
- Added unit tests and examples
- Added README and MIT license

## [v0.2.0] - 2026-04-16
### Added
- ZScoreForServiceLevel for converting target service levels to z-scores
- SafetyStockWithServiceLevel for probabilistic safety stock calculations
- ReorderPointWithServiceLevel for reorder point calculations using service level targets

### Improved
- validation coverage for service-level-based inventory functions
- clearer error handling for invalid service level and negative standard deviation inputs