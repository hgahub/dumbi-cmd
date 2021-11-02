package command

import "testing"

func TestInvalidCommand(t *testing.T) {
	c := &Command{
		Name: "invalid",
	}

	if exitCode := c.Execute(); exitCode != 126 {
		t.Fatalf("Command execution error: %d", exitCode)
	}
}
