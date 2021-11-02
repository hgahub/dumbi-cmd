package command

import (
	"fmt"
	"github.com/hgahub/dumbi-cmd/internal/dumbi/model"
)

const ValidateCommandName = "validate"
const ValidateCommandDescription = "validates the configuration files"

type ValidateCommand struct {
	meta *Meta
}

func (c *ValidateCommand) run() int {
	if _, err := model.LoadDumbi(c.meta.workspacePath, DumbiIndexFile); err != nil {
		fmt.Println(err)
		return 1
	}

	fmt.Println("The model is valid.")
	return 0
}
