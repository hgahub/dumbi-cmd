package version

import (
	"regexp"
	"strings"
	"testing"
)

var isDigit = regexp.MustCompile(`^[0-9]+$`).MatchString

func TestVersionString_Empty(t *testing.T) {
	versionString := String("")

	if !isDigit(versionString) {
		t.Fatalf("Invalid versionString: %s", versionString)
	}
}

func TestVersionString_GitHash(t *testing.T) {
	gitVersion := "112233"
	versionString := String(gitVersion)
	parts := strings.Split(versionString, "-")

	if len(parts) != 2 {
		t.Fatalf("Invalid versionString: %s", versionString)
	}

	if !isDigit(parts[0]) {
		t.Fatalf("Invalid versionString: %s", versionString)
	}

	if parts[1] != gitVersion {
		t.Fatalf("Invalid versionString: %s", versionString)
	}
}
