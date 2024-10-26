package cmd

import (
	"fmt"
	"math"

	"github.com/spf13/cobra"
)

func ex2Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "ex2",
		Run: func(cmd *cobra.Command, args []string) {
		loop:
			var code string
			fmt.Print("Enter code: ")
			fmt.Scanln(&code)

			result := StartEx2(code)
			fmt.Println(result)

			var exit string
			fmt.Print("Do you want to exit? Y/N: ")
			fmt.Scanln(&exit)

			if exit == "Y" || exit == "y" {
				goto loop
			}
		},
	}
	return cmd
}

func StartEx2(code string) string {
	n := len(code) + 1
	bestString := ""
	minSum := math.MaxInt64

	var backtrack func(pos int, arr []int)
	backtrack = func(pos int, arr []int) {

		if pos == n {

			sum := 0
			for _, num := range arr {
				sum += num
			}

			if sum < minSum {
				minSum = sum
				bestString = ""
				for _, num := range arr {
					bestString += fmt.Sprintf("%d", num)
				}
			}
			return
		}

		for i := 0; i <= 9; i++ {
			if pos > 0 {
				prev := arr[pos-1]

				switch code[pos-1] {
				case 'L':
					if prev <= i {
						continue
					}
				case 'R':
					if prev >= i {
						continue
					}
				case '=':
					if prev != i {
						continue
					}
				}
			}

			backtrack(pos+1, append(arr, i))
		}
	}

	backtrack(0, []int{})

	return bestString
}
