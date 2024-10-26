package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

func ex1cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "ex1",
		Run: func(cmd *cobra.Command, args []string) {

			data, err := readJson()
			if err != nil {
				panic(err)
			}

			result := StartEx1(*data)

			fmt.Println(result)
		},
	}
	return cmd
}

func StartEx1(bigDatas [][]int) int {
	for i := len(bigDatas) - 2; i >= 0; i-- {

		for j := 0; j < len(bigDatas[i]); j++ {

			if bigDatas[i+1][j] > bigDatas[i+1][j+1] {
				bigDatas[i][j] += bigDatas[i+1][j]
			} else {
				bigDatas[i][j] += bigDatas[i+1][j+1]
			}
		}
	}

	return bigDatas[0][0]
}

func readJson() (*[][]int, error) {
	file, err := os.Open("json/hard.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var bigDatas = [][]int{}
	err = json.Unmarshal(byteValue, &bigDatas)
	if err != nil {
		return nil, err
	}

	return &bigDatas, nil
}
