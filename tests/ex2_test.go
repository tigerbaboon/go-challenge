package tests

import (
	"app/cmd"
	"testing"
)

func TestEx2Cmd(t *testing.T) {

	testCases := []struct {
		name     string
		data     string
		expected string
	}{
		{"Test ex2 case 1", "LLRR=", "210122"},
		{"Test ex2 case 2", "==RLL", "000210"},
		{"Test ex2 case 3", "=LLRR", "221012"},
		{"Test ex2 case 4", "RRL=R", "012001"},
	}

	for _, row := range testCases {
		t.Run(row.name, func(t *testing.T) {

			result := cmd.StartEx2(row.data)

			if result != row.expected {
				t.Errorf("Test %s failed: Expected %s, got %s", row.name, row.expected, result)
			} else {
				t.Logf("Test %s passed: Expected %s, got %s", row.name, row.expected, result)
			}
		})
	}
}
