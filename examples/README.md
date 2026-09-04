# Examples

Runnable programs showing complete workflows, not just isolated formula
calls. Each one is `go run`-able directly from the repo root.

| Example | What it shows |
|---|---|
| [`inventory-policy`](inventory-policy/main.go) | Demand + lead time + service target → safety stock, reorder point, target level, and EOQ |
| [`forecasting`](forecasting/main.go) | One historical series → several forecast methods, compared with accuracy metrics |
| [`intermittent-demand`](intermittent-demand/main.go) | A sparse SKU → demand-pattern classification, then Croston/SBA forecasting |
| [`abc-xyz`](abc-xyz/main.go) | A SKU set → ABC (value) and XYZ (variability) classification, combined into the ABC-XYZ matrix |
| [`retail`](retail/main.go) | The full pipeline for several SKUs at once: classify demand pattern → forecast → ABC class → reorder point → EOQ |

```bash
go run ./examples/inventory-policy
go run ./examples/forecasting
go run ./examples/intermittent-demand
go run ./examples/abc-xyz
go run ./examples/retail
```

These are standalone `main` packages for demonstration — they aren't part
of the importable API. For the formula-level API reference, see each
package's [pkg.go.dev](https://pkg.go.dev/github.com/motah-fard/scmgo)
page and its runnable `Example` functions.
