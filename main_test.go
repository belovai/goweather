package main

import "testing"

func TestCalculateDirections(t *testing.T) {
	if CalculateDirections(0) != "N" {
		t.Error("Error in N 0")
	}

	if CalculateDirections(22) != "N" {
		t.Error("Error in N 22")
	}

	if CalculateDirections(23) == "N" {
		t.Error("Error in N 23")
	}

	if CalculateDirections(359) != "N" {
		t.Error("Error in N 359")
	}

	if CalculateDirections(338) != "N" {
		t.Error("Error in N 338")
	}

	if CalculateDirections(337) == "N" {
		t.Error("Error in N 337")
	}

	if CalculateDirections(90) != "E" {
		t.Error("Error in E 90")
	}

	if CalculateDirections(270) != "W" {
		t.Error("Error in W 270")
	}

	if CalculateDirections(45) != "NE" {
		t.Error("Error in NE 45")
	}

	if CalculateDirections(23) != "NE" {
		t.Error("Error in NE 23")
	}

}

func TestBuildParams(t *testing.T) {
	var params = map[string]string{
		"a": "avalue",
		"b": "bvalue",
	}
	if BuildParams(params) != "a=avalue&b=bvalue" {
		t.Error("Builded query not valid")
	}
}

func TestTableBuildParams(t *testing.T) {
	var tests = []struct {
		input    map[string]string
		expected string
	}{
		{
			map[string]string{
				"a": "avalue",
				"b": "bvalue",
			},
			"a=avalue&b=bvalue",
		},
		{
			map[string]string{
				"a": "avalue",
			},
			"a=avalue",
		},
		{
			map[string]string{
				"a": "avalue",
				"b": "bvalue",
				"c": "cvalue",
			},
			"a=avalue&b=bvalue&c=cvalue",
		},
	}

	for _, test := range tests {
		if output := BuildParams(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
