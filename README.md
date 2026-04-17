# scmgo

[![Go Reference](https://pkg.go.dev/badge/github.com/motah-fard/scmgo/inventory.svg)](https://pkg.go.dev/github.com/motah-fard/scmgo/inventory)
[![License](https://img.shields.io/github/license/motah-fard/scmgo)](https://github.com/motah-fard/scmgo/blob/main/LICENSE)
[![GitHub release](https://img.shields.io/github/v/release/motah-fard/scmgo)](https://github.com/motah-fard/scmgo/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/motah-fard/scmgo)](https://goreportcard.com/report/github.com/motah-fard/scmgo)

`scmgo` is a Go library for practical inventory and supply-chain calculations.

The first package, `inventory`, provides clear and reusable functions for common inventory policy calculations such as reorder point, safety stock, EOQ, min/max levels, and service-level-based threshold planning.

The goal is to keep the API:

- simple
- transparent
- practical
- easy to embed in Go applications

## Current Scope

As of `v0.3.0`, the `inventory` package includes:

- `ReorderPoint`
- `SafetyStockBasic`
- `EOQ`
- `MinMaxLevels`
- `ZScoreForServiceLevel`
- `SafetyStockWithServiceLevel`
- `ReorderPointWithServiceLevel`
- `DemandDuringLeadTime`
- `StdDevDemandDuringLeadTime`

## Why scmgo

Many inventory and supply-chain calculations still live in spreadsheets, internal notes, or one-off scripts. `scmgo` provides a lightweight Go-native alternative for developers building:

- inventory tools
- supply-chain applications
- planning dashboards
- internal operations services
- educational and analytical tools

The package is intentionally small, explicit, and easy to embed.

## Installation

```bash
go get github.com/motah-fard/scmgo