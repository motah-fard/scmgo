package abc

// Combine joins ABC and XYZ classification results by ID into the classic
// ABC-XYZ matrix (e.g. "AX" for a high-value, predictable-demand item that
// deserves the tightest control, versus "CZ" for a low-value,
// unpredictable one).
//
// This is an inner join: an ID present in only one of abcResults or
// xyzResults is omitted from the result. The result order follows
// abcResults.
func Combine(abcResults []ClassifiedItem, xyzResults []VariabilityClassifiedItem) []CombinedClass {
	xyzByID := make(map[string]string, len(xyzResults))
	for _, item := range xyzResults {
		xyzByID[item.ID] = item.Class
	}

	combined := make([]CombinedClass, 0, len(abcResults))
	for _, item := range abcResults {
		xyzClass, ok := xyzByID[item.ID]
		if !ok {
			continue
		}
		combined = append(combined, CombinedClass{
			ID:       item.ID,
			ABCClass: item.Class,
			XYZClass: xyzClass,
		})
	}

	return combined
}
