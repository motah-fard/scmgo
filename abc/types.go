package abc

// Item is one item's value contribution to classify by ABC (Pareto)
// analysis — typically annual usage value (annual quantity × unit cost),
// but any non-negative value metric works.
type Item struct {
	ID    string
	Value float64
}

// ClassifiedItem is one item's ABC classification result.
type ClassifiedItem struct {
	ID    string
	Value float64
	// CumulativePercent is the cumulative percentage of total value
	// contributed by this item and every item with a strictly greater
	// value (ties broken by original input order).
	CumulativePercent float64
	// Class is "A", "B", or "C".
	Class string
}

// ClassifyInput contains the inputs required for ABC classification.
type ClassifyInput struct {
	Items []Item
	// AThreshold is the cumulative-percent cutoff for class A (e.g. 0.80).
	// Must be in (0, 1).
	AThreshold float64
	// BThreshold is the cumulative-percent cutoff for class B (e.g. 0.95).
	// Must be greater than AThreshold and at most 1.
	BThreshold float64
}

// VariabilityItem is one item's demand statistics to classify by XYZ
// analysis.
type VariabilityItem struct {
	ID           string
	MeanDemand   float64
	StdDevDemand float64
}

// VariabilityClassifiedItem is one item's XYZ classification result.
type VariabilityClassifiedItem struct {
	ID string
	// CoefficientOfVariation is StdDevDemand / MeanDemand.
	CoefficientOfVariation float64
	// Class is "X" (low variability), "Y" (medium), or "Z" (high).
	Class string
}

// ClassifyVariabilityInput contains the inputs required for XYZ
// classification.
type ClassifyVariabilityInput struct {
	Items []VariabilityItem
	// XThreshold is the coefficient-of-variation cutoff for class X (low
	// variability). Must be greater than zero.
	XThreshold float64
	// YThreshold is the coefficient-of-variation cutoff for class Y
	// (medium variability); above it is class Z (high variability). Must
	// be greater than XThreshold.
	YThreshold float64
}

// CombinedClass is one item's joint ABC-XYZ classification.
type CombinedClass struct {
	ID       string
	ABCClass string
	XYZClass string
}
