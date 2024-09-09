package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var rows [50][50]bool
var divider = strings.Repeat("-", 50)

func main() {
	rows[26][26] = true
	rows[26][25] = true
	rows[25][25] = true
	rows[25][24] = true
	rows[24][25] = true

	for {
		//fmt.Println(divider)
		doGameLogic()
		clearScreen()
		printGame()

		//time.Sleep(100 * time.Millisecond)
	}
}

func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ternary(condition bool, trueVal, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}

func printGame() {
	for i := range rows {
		for j := range rows[i] {
			fmt.Print(ternary(rows[i][j], "███", "░░░"))
		}
		fmt.Println()
	}
}

func activeNeighbourCount(rowIdx int, colIdx int) int {
	activeCount := 0
	for i := -1; i <= 1; i++ {
		nRow := rowIdx + i

		if nRow < 0 || nRow >= 50 {
			continue
		}

		for j := -1; j <= 1; j++ {
			nCol := colIdx + j

			if nCol < 0 || nCol >= 50 || (nRow == rowIdx && nCol == colIdx) {
				continue
			}

			if rows[nRow][nCol] {
				activeCount++
			}
		}
	}

	return activeCount
}

func doGameLogic() {
	newGame := [50][50]bool{}
	for i := range rows {
		for j := range rows[i] {
			activeNeighbourCount := activeNeighbourCount(i, j)

			if rows[i][j] {
				if activeNeighbourCount == 2 || activeNeighbourCount == 3 {
					newGame[i][j] = true
				}
			} else {
				if activeNeighbourCount == 3 {
					newGame[i][j] = true
				}
			}
		}
	}
	rows = newGame
}
