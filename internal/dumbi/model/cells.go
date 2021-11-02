package model

import (
	"fmt"
)

type Cell struct {
	// Id, cell identifier
	Id *string `json:"id"`

	Type *CellType `json:"cell_type"`
}

func (c *Cell) IdString() string {
	if id := c.Id; id != nil {
		return *c.Id
	}

	return ""
}

func (c *Cell) Validate() error {
	var err error

	if c.Id == nil {
		return fmt.Errorf("CellIdError: id property is mandatory")
	}

	if c.Type == nil {
		return fmt.Errorf("CellTypeError: cell_type property is mandatory")
	} else {
		err = c.Type.Validate()
	}

	return err
}
