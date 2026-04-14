# scmgo

[![Go Reference](https://pkg.go.dev/badge/github.com/motah-fard/scmgo/inventory.svg)](https://pkg.go.dev/github.com/motah-fard/scmgo/inventory)
[![License](https://img.shields.io/github/license/motah-fard/scmgo)](https://github.com/motah-fard/scmgo/blob/main/LICENSE)
[![GitHub release](https://img.shields.io/github/v/release/motah-fard/scmgo)](https://github.com/motah-fard/scmgo/releases)

`scmgo` is a Go library for practical inventory and supply-chain threshold calculations.

The first package, `inventory`, provides a small set of clear, reusable formulas for common inventory policy calculations. The goal is to keep the API simple, transparent, and easy to embed in Go applications.

## v0.1.2 Scope

Initial release includes:

- `ReorderPoint`
- `SafetyStockBasic`
- `EOQ`
- `MinMaxLevels`

## Features

Version `v0.1.2` includes support for:

- Reorder point
- Basic safety stock
- Economic order quantity (EOQ)
- Min/max inventory levels

## Installation

```bash
go get github.com/motah-fard/scmgo
