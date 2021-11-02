package model

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestInvalidPath(t *testing.T) {
	path, _ := os.Getwd()
	workspacePath := filepath.Join(path, "../../../tests/sample")

	if _, err := LoadDumbi(workspacePath, "_id.df"); err == nil {
		t.Fatalf("Read error: %d", err)
	} else if !strings.HasPrefix(err.Error(), "FileIOError") {
		t.Fatalf("invalid error code: %d", err)
	}
}

func TestInvalidFormat(t *testing.T) {
	path, _ := os.Getwd()
	workspacePath := filepath.Join(path, "../../../tests/sample")

	if _, err := LoadDumbi(workspacePath, "invalid_format.df"); err == nil {
		t.Fatalf("Unmarshal error: %d", err)
	} else if !strings.HasPrefix(err.Error(), "FileFormatError") {
		t.Fatalf("invalid error code: %d", err)
	}
}

func TestIdIsNil(t *testing.T) {
	path, _ := os.Getwd()
	workspacePath := filepath.Join(path, "../../../tests/sample")

	if _, err := LoadDumbi(workspacePath, "id_nil.df"); err == nil {
		t.Fatalf("id test error: %d", err)
	} else if !strings.HasPrefix(err.Error(), "ModelIdError") {
		t.Fatalf("invalid error code: %d", err)
	}
}

func TestCellsParamIsNil(t *testing.T) {
	path, _ := os.Getwd()
	workspacePath := filepath.Join(path, "../../../tests/sample")

	if _, err := LoadDumbi(workspacePath, "cells_nil.df"); err == nil {
		t.Fatalf("load error: %d", err)
	} else if !strings.HasPrefix(err.Error(), "CellsError") {
		t.Fatalf("invalid error code: %d", err)
	}
}

func TestCellsParamIsEmpty(t *testing.T) {
	path, _ := os.Getwd()
	workspacePath := filepath.Join(path, "../../../tests/sample")

	if _, err := LoadDumbi(workspacePath, "cells_empty.df"); err == nil {
		t.Fatalf("load error: %d", err)
	} else if !strings.HasPrefix(err.Error(), "CellsError") {
		t.Fatalf("invalid error code: %d", err)
	}
}

func TestCellItemIsNil(t *testing.T) {
	path, _ := os.Getwd()
	workspacePath := filepath.Join(path, "../../../tests/sample")

	if _, err := LoadDumbi(workspacePath, "cell_id_nil.df"); err == nil {
		t.Fatalf("load error: %d", err)
	} else if !strings.HasPrefix(err.Error(), "CellError") {
		t.Fatalf("invalid error code: %d", err)
	}
}

func TestValidFile(t *testing.T) {
	path, _ := os.Getwd()
	workspacePath := filepath.Join(path, "../../../tests/sample/default")

	if _, err := LoadDumbi(workspacePath, "index.df"); err != nil {
		t.Fatalf("ivalid default file: %d", err)
	}
}
