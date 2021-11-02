package model

import "testing"

func TestSupportedTypesString(t *testing.T) {
	if st := supportedTypesString(); st != "array, boolean, date, double, enum, integer, string" {
		t.Fatalf("invalid supported type string: %s", st)
	}
}

func TestMarshal(t *testing.T) {
	var c CellType = stringType
	b, _ := c.MarshalJSON()

	var u CellType
	_ = u.UnmarshalJSON(b)

	if c != u {
		t.Fatalf("CellType marshal error")
	}
}

func TestNilUnmarshal(t *testing.T) {
	var u CellType
	if err := u.UnmarshalJSON(nil); err == nil {
		t.Fatalf("CellType unmarshal error")
	}
}
