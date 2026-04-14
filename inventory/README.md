# v0.1.0 Scope

Initial release includes:

- ReorderPoint
- SafetyStockBasic
- EOQ
- MinMaxLevels

# scmgo

`scmgo` is a Go library for practical inventory and supply-chain threshold calculations.

The first package, `inventory`, provides a small set of clear and reusable formulas for common inventory policy calculations.

## Features

Version `v0.1.0` includes:

- Reorder point
- Basic safety stock
- reorder point = average daily demand × lead time + safety stock
- If the result is negative, the library returns 0 for this basic model.
- safety stock = (max daily demand × max lead time) -
               (average daily demand × average lead time)
  
- Economic order quantity (EOQ)
- Min/max inventory levels
- EOQ = sqrt((2 × annual demand × ordering cost) / holding cost per unit)
Design Goals
This library aims to be:
simple
transparent
easy to use
easy to embed in Go applications
practical for backend systems, tools, and learning
Assumptions
Input units must be consistent.
Example: if demand is measured per day, lead time should also be in days.
Inputs must be non-negative unless otherwise stated.
HoldingCostPerUnit must be greater than zero for EOQ.
This version uses basic deterministic formulas and does not yet include:
forecasting
simulation
service-level-based safety stock
stochastic inventory models
Roadmap
Possible future additions:
service-level-based safety stock
z-score helpers
demand forecasting helpers
inventory simulation tools
REST API / app layer built on top of the library
License
MIT

## Important
Replace:

```text
github.com/motah-fard/scmgo
## Installation

```bash
go get github.com/motah-fard/scmgo

