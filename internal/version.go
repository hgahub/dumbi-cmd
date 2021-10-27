package internal

import (
	"fmt"
	"strings"
	"time"
)

type Version struct {
	releaseVersion string
	releaseDate    time.Time
	commit         string
}

const (
	// Header is the header name used to send the current dumbi version in http requests.
	Header = "Dumbi-Version"

	// defaultCommit is a default commit.
	defaultCommit = "0000000"
)

// NewVersion create a Version struct
func NewVersion(commit string, releaseVersion string, releaseDate string) *Version {
	t, err := time.Parse("2006-01-02T15:04:05.000Z", releaseDate)
	if err != nil {
		t = time.Now()
	}

	if !strings.HasSuffix(commit, "-dirty") {
		commit = (commit + defaultCommit)[0:7]
	}

	v := Version{
		releaseVersion: releaseVersion,
		releaseDate:    t,
		commit:         commit,
	}
	return &v
}

// String returns the complete version string, including short Git commit hash.
func (v *Version) String() string {
	if v.releaseVersion == "" || strings.Contains(v.releaseVersion, "-snapshot") {
		if v.commit == defaultCommit {
			return v.releaseDate.Format("20060102")
		}
		return fmt.Sprintf("%s-%s", v.releaseDate.Format("20060102"), v.commit)
	}
	return fmt.Sprintf("%s (%s-%s)", v.releaseVersion, v.releaseDate.Format("20060102"), v.commit)
}
