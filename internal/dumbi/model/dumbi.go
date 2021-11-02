package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Dumbi struct {
	// Id, model identifier
	Id *string `json:"id"`

	// Cells, cell list
	Cells []Cell `json:"cells"`
}

func LoadDumbi(workspacePath string, dumbiFile string) (*Dumbi, error) {
	var dumbiModel Dumbi
	dfPath := filepath.Join(workspacePath, dumbiFile)

	err := dumbiModel.load(dfPath)

	return &dumbiModel, err
}

func (d *Dumbi) load(dfPath string) error {
	content, err := ioutil.ReadFile(dfPath)
	if err != nil {
		return fmt.Errorf("FileIOError: %s", err)
	}

	err = json.Unmarshal(content, &d)
	if err != nil {
		return fmt.Errorf("FileFormatError: error during unmarshal(): %s", err)
	}

	if d.Id == nil {
		return fmt.Errorf("ModelIdError: id property is mandatory")
	}

	if d.Cells == nil || len(d.Cells) == 0 {
		return fmt.Errorf("CellsError: cells property is mandatory")
	}

	for _, cell := range d.Cells {
		if cellErr := cell.Validate(); cellErr != nil {
			return fmt.Errorf("CellError: invalid cell '%s': \n - %s", cell.IdString(), cellErr)
		}
	}
	return nil
}
