package version

import (
	"fmt"
	"time"
)

// The main version number that is being run at the moment.
var version string = time.Now().Format("20060102")

// Header is the header name used to send the current dumbi version in http requests.
const Header = "Dumbi-Version"

// String returns the complete version string, including short Git commit hash.
func String(gitVersion string) string {
	if gitVersion != "" {
		return fmt.Sprintf("%s-%s", version, gitVersion)
	}

	return version
}
