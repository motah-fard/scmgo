// Package numeric holds tiny numeric helpers shared by scmgo's public
// packages. It is internal because it is an implementation detail, not part
// of the library's API surface.
package numeric

import "math"

// AllFinite reports whether every value is finite: not NaN and not an
// infinity.
//
// It exists because comparisons against NaN are always false in IEEE 754,
// so a plain "v < 0" or range check silently passes NaN through instead of
// rejecting it. Every exported function in inventory and forecast checks
// its raw float64 inputs with this before any other validation.
func AllFinite(values ...float64) bool {
	for _, v := range values {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}
