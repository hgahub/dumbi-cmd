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
	commit := "4a396b587c7ff"
	versionString := String(commit)
	parts := strings.Split(versionString, "-")

	if len(parts) != 2 {
		t.Fatalf("Invalid versionString: %s", versionString)
	}

	if !isDigit(parts[0]) {
		t.Fatalf("Invalid date: %s", versionString)
	}

	if parts[1] != commit[0:7] {
		t.Fatalf("Invalid commit: %s", versionString)
	}

	if len(parts[1]) > 7 {
		t.Fatalf("Invalid commit length: %s", versionString)
	}
}
