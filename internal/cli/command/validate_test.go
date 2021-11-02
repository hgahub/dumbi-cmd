package command

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNotFoundIndexFile(t *testing.T) {
	old := os.Getenv(DumbiWorkspacePathEnvVar)
	defer func(key, value string) {
		_ = os.Setenv(key, value)
	}(DumbiWorkspacePathEnvVar, old)

	c := &Command{
		Name: ValidateCommandName,
	}

	path, _ := os.Getwd()
	testPath := filepath.Join(path, "../../../tests")

	_ = os.Setenv(DumbiWorkspacePathEnvVar, testPath)
	if exitCode := c.Execute(); exitCode != 1 {
		t.Fatalf("Invalid workspace: %d", exitCode)
	}
}

func TestValidIndexFile(t *testing.T) {
	old := os.Getenv(DumbiWorkspacePathEnvVar)
	defer func(key, value string) {
		_ = os.Setenv(key, value)
	}(DumbiWorkspacePathEnvVar, old)

	c := &Command{
		Name: ValidateCommandName,
	}

	path, _ := os.Getwd()
	testPath := filepath.Join(path, "../../../tests/sample/default")

	_ = os.Setenv(DumbiWorkspacePathEnvVar, testPath)
	if exitCode := c.Execute(); exitCode != 0 {
		t.Fatalf("Invalid file: %d", exitCode)
	}
}
