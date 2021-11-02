package model

import (
	"strings"
	"testing"
)

func TestCellIdIsNil(t *testing.T) {
	cell := &Cell{Id: nil}

	if err := cell.Validate(); err == nil {
		t.Fatalf("id error: %d", err)
	} else if !strings.HasPrefix(err.Error(), "CellIdError") {
		t.Fatalf("invalid error code: %d", err)
	}
}

func TestCellTypeIsNil(t *testing.T) {
	id := "1"
	cell := &Cell{Id: &id, Type: nil}

	if err := cell.Validate(); err == nil {
		t.Fatalf("cell id error: %d", err)
	} else if !strings.HasPrefix(err.Error(), "CellTypeError") {
		t.Fatalf("invalid error code: %d", err)
	}
}

func TestCellTypeIsUnknown(t *testing.T) {
	id := "1"
	var cellType CellType = unknownType
	cell := &Cell{Id: &id, Type: &cellType}

	if err := cell.Validate(); err == nil {
		t.Fatalf("cell type error: %d", err)
	} else if !strings.HasPrefix(err.Error(), "CellTypeError") {
		t.Fatalf("invalid error code: %d", err)
	}
}

func TestValidCellType(t *testing.T) {
	id := "1"

	var cellType CellType = stringType
	cell := &Cell{Id: &id, Type: &cellType}

	if err := cell.Validate(); err != nil {
		t.Fatalf("invalid type: %s", err)
	}
}

func TestIdStringIsNil(t *testing.T) {
	cell := &Cell{Id: nil}

	if cell.IdString() != "" {
		t.Fatalf("invalid cell id string")
	}
}

func TestIdString(t *testing.T) {
	id := "1"

	var cellType CellType = stringType
	cell := &Cell{Id: &id, Type: &cellType}

	if cell.IdString() != "1" {
		t.Fatalf("invalid cell id string")
	}
}
