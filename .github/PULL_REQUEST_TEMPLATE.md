## Summary

<!-- What does this PR add or change, and why? -->

## Checklist

- [ ] Follows the package pattern described in CONTRIBUTING.md (Input struct, validation with sentinel errors, doc.go, tests, example)
- [ ] `go build ./...`, `go vet ./...`, and `gofmt -l .` are clean
- [ ] `go test ./... -race -cover` passes and covers new branches (including invalid-input cases)
- [ ] Public functions and types have doc comments
- [ ] README updated if this adds/changes a public function or package
- [ ] CHANGELOG.md updated under `[Unreleased]`
- [ ] New formulas cite a source (textbook, paper, standard) in the doc comment
