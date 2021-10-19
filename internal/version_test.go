package internal

import (
	"regexp"
	"strings"
	"testing"
)

var isDigit = regexp.MustCompile(`^[0-9]+$`).MatchString

func TestStringWithEmptyParameter(t *testing.T) {
	version := NewVersion("", "", "")

	if !isDigit(version.String()) {
		t.Fatalf("Invalid versionString: %s", version.String())
	}

	if len(version.String()) != 8 {
		t.Fatalf("Invalid versionString length: %d", len(version.String()))
	}
}

func TestStringWithCommit(t *testing.T) {
	commit := "4a396b587c7ff"
	version := NewVersion(commit, "", "")
	testBuildParts(t, version.String(), commit)
}

func TestStringWithReleaseParameter(t *testing.T) {
	commit := "4a396b587c7ff"
	releaseVersion := "1.0.0"
	releaseDate := "2001-12-01T10:11:12.123Z"
	version := NewVersion(commit, releaseVersion, releaseDate)

	parts := strings.Split(version.String(), "(")

	if len(parts) != 2 {
		t.Fatalf("Invalid versionString: %s", version.String())
	}

	if strings.TrimSpace(parts[0]) != releaseVersion {
		t.Fatalf("Invalid release part: %s", strings.TrimSpace(parts[0]))
	}

	testBuildParts(t, strings.Trim(parts[1], ")"), commit)
}

func testBuildParts(t *testing.T, buildParts string, commit string) {
	parts := strings.Split(buildParts, "-")

	if len(parts) != 2 {
		t.Fatalf("Invalid build parts: %s", buildParts)
	}

	if !isDigit(parts[0]) {
		t.Fatalf("Invalid date: %s", parts[0])
	}

	if len(parts[0]) != 8 {
		t.Fatalf("Invalid date length: %d", len(parts[0]))
	}

	if parts[1] != commit[0:7] {
		t.Fatalf("Invalid commit: %s", parts[1])
	}

	if len(parts[1]) != 7 {
		t.Fatalf("Invalid commit length: %d", len(parts[1]))
	}
}
