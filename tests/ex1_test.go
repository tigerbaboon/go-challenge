package tests

import (
	"app/cmd"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var bigDatas = [][]int{}

var shortDatas = [][]int{
	{59},
	{73, 41},
	{52, 40, 53},
	{26, 53, 6, 34},
}

func TestEx1Cmd(t *testing.T) {

	err := readJson()
	if err != nil {
		panic(err)
	}

	testCases := []struct {
		name     string
		data     [][]int
		expected int
	}{
		{"Test ex1 short input", shortDatas, 237},
		{"Test ex1 long input", bigDatas, 7273},
	}

	for _, row := range testCases {
		t.Run(row.name, func(t *testing.T) {

			result := cmd.StartEx1(row.data)

			if result != row.expected {
				t.Errorf("Test %s failed: Expected %d, got %d", row.name, row.expected, result)
			} else {
				t.Logf("Test %s passed: Expected %d, got %d", row.name, row.expected, result)
			}
		})
	}
}

func readJson() error {
	file, err := os.Open("../json/hard.json")
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println(file)

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, &bigDatas)
	if err != nil {
		return err
	}

	return nil
}
