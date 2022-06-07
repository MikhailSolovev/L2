package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	name    string
	isValid bool
	ps      PackedString
	want    string
}{
	{
		name:    "valid-001",
		isValid: true,
		ps:      "a4bc2d5e",
		want:    "aaaabccddddde",
	},
	{
		name:    "valid-002",
		isValid: true,
		ps:      "abcd",
		want:    "abcd",
	},
	{
		name:    "not valid-001",
		isValid: false,
		ps:      "45",
		want:    "",
	},
	{
		name:    "valid-003",
		isValid: true,
		ps:      "",
		want:    "",
	},
	{
		name:    "valid-004",
		isValid: true,
		ps:      "qwe\\4\\5",
		want:    "qwe45",
	},
	{
		name:    "valid-005",
		isValid: true,
		ps:      "qwe\\45",
		want:    "qwe44444",
	},
	{
		name:    "valid-006",
		isValid: true,
		ps:      "qwe\\\\5",
		want:    "qwe\\\\\\\\\\",
	},
	{
		name:    "not valid-002",
		isValid: false,
		ps:      "\\\\\\",
		want:    "",
	},
	{
		name:    "not valid-003",
		isValid: false,
		ps:      "a\\\\4\\",
		want:    "",
	},
}

func TestUnpack(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.isValid == true {
				res, err := test.ps.Unpack()
				if err != nil {
					t.Fatal(err)
				}
				if res != test.want {
					t.Fatal(fmt.Sprintf("Result: %s Want: %s", res, test.want))
				}
			} else {
				res, err := test.ps.Unpack()
				if err == nil {
					t.Fatal(err)
				}
				if res != test.want {
					t.Fatal(fmt.Sprintf("Result: %s Want: %s", res, test.want))
				}
			}
		})
	}
}
