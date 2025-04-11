package urlparser

import (
	"strings"
)

func InjectFUZZIdentifiers(path string, fuzzIdentifier string) []string {
	unique := make(map[string]struct{})
	segments := strings.Split(path, "/")

	// Replace each segment with FUZZ one by one
	for i := range segments {
		if segments[i] == "" {
			continue
		}
		modified := make([]string, len(segments))
		copy(modified, segments)
		modified[i] = fuzzIdentifier
		fuzzed := strings.Join(modified, "/")
		unique[fuzzed] = struct{}{}
	}

	// Trim segments and add FUZZ at the end (shortened paths)
	for i := len(segments) - 1; i > 0; i-- {
		if segments[i] == "" {
			continue
		}
		trimmed := append(segments[:i], fuzzIdentifier)
		fuzzed := strings.Join(trimmed, "/")
		unique[fuzzed] = struct{}{}
	}

	// Append /FUZZ to full path
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	unique[path+fuzzIdentifier] = struct{}{}

	// Replace full path with just /FUZZ
	unique["/"+fuzzIdentifier] = struct{}{}

	// Convert map to slice
	var fuzzedPaths []string
	for p := range unique {
		fuzzedPaths = append(fuzzedPaths, p)
	}

	return fuzzedPaths
}
