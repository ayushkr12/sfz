package urlparser

import (
	"testing"
)

func TestInjectFUZZIdentifiers(t *testing.T) {
	path := "/api/v2/resource/resource"
	fuzzIdentifier := "FUZZ"

	fuzzablePaths := InjectFUZZIdentifiers(path, fuzzIdentifier)

	for _, path := range fuzzablePaths {
		t.Log(path)
	}

}
