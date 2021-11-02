package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type CellType int

const (
	unknownType = iota
	stringType
	booleanType
	integerType
	doubleType
	dateType
	arrayType
	enumType
)

var toString = map[CellType]string{
	stringType:  "string",
	booleanType: "boolean",
	integerType: "integer",
	doubleType:  "double",
	dateType:    "date",
	arrayType:   "array",
	enumType:    "enum",
}

var toID = map[string]CellType{
	"string":  stringType,
	"boolean": booleanType,
	"integer": integerType,
	"double":  doubleType,
	"date":    dateType,
	"array":   arrayType,
	"enum":    enumType,
}

func (c CellType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[c])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (c *CellType) UnmarshalJSON(b []byte) error {
	var j string
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}

	*c = toID[strings.ToLower(j)]
	return nil
}

func supportedTypesString() string {
	supportedTypes := ""
	keys := make([]string, 0, len(toID))

	for k := range toID {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		supportedTypes += k + ", "
	}

	return supportedTypes[:len(supportedTypes)-2]
}

func (c *CellType) Validate() error {
	if *c == 0 {
		return fmt.Errorf("CellTypeError: allowed cell_type is: %s", supportedTypesString())
	}

	return nil
}
