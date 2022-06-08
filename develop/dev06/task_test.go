package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	name string
	args []string
	want string
}{
	{
		name: "test-001",
		args: []string{"-f4", "\"cat\tdog\thorse\tcow\tduck\telephant\tchicken\tparrot\""},
		want: "cow",
	},
	{
		name: "test-002",
		args: []string{"-f-3", "\"cat\tdog\thorse\tcow\tduck\telephant\tchicken\tparrot\""},
		want: "cat\tdog\thorse",
	},
	{
		name: "test-003",
		args: []string{"-f3-", "\"cat\tdog\thorse\tcow\tduck\telephant\tchicken\tparrot\""},
		want: "horse\tcow\tduck\telephant\tchicken\tparrot",
	},
	{
		name: "test-004",
		args: []string{"-f2-5", "\"cat\tdog\thorse\tcow\tduck\telephant\tchicken\tparrot\""},
		want: "dog\thorse\tcow\tduck",
	},
	{
		name: "test-005",
		args: []string{`-d" "`, "-f2-5", "\"cat dog horse cow duck elephant chicken parrot\""},
		want: "dog horse cow duck",
	},
	{
		name: "test-006",
		args: []string{`-d" "`, "-s", "-f2-5", "\"cat\tdog\thorse\tcow\tduck\telephant\tchicken\tparrot\""},
		want: "",
	},
	{
		name: "test-008",
		args: []string{"-s", "-f10-12", "\"cat\tdog\thorse\tcow\tduck\telephant\tchicken\tparrot\""},
		want: "",
	},
}

func TestRunCut(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := RunCut(test.args)
			if res != test.want {
				t.Fatal(fmt.Sprintf("Result: %s Want: %s", res, test.want))
			}
		})
	}
}
