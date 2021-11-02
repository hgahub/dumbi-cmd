package command

import (
	"os"
	"path/filepath"
)

// DumbiWorkspacePathEnvVar is the name of the environment variable that can be used
// to set the path of the Dumbi workspace
const DumbiWorkspacePathEnvVar = "DUMBI_WORKSPACE_PATH"

// DumbiIndexFile is the name of the index file
const DumbiIndexFile = "index.df"

type Meta struct {
	workspacePath string
}

func NewMeta() *Meta {
	return &Meta{
		workspacePath: getWorkspacePath(),
	}
}

func getWorkspacePath() string {
	if envVar := os.Getenv(DumbiWorkspacePathEnvVar); envVar != "" {
		return envVar
	} else {
		ex, _ := os.Executable()
		return filepath.Dir(ex)
	}
}
