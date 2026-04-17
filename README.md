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

As of `v0.2.0`, the `inventory` package includes:

- `ReorderPoint`
- `SafetyStockBasic`
- `EOQ`
- `MinMaxLevels`
- `ZScoreForServiceLevel`
- `SafetyStockWithServiceLevel`
- `ReorderPointWithServiceLevel`

## Why scmgo

Many supply-chain and inventory calculations still live in spreadsheets, internal notes, or one-off scripts. `scmgo` aims to provide a lightweight Go-native alternative for developers building:

- inventory tools
- supply-chain applications
- operations dashboards
- internal planning services
- educational or analytical tools

## Installation

```bash
go get github.com/motah-fard/scmgo