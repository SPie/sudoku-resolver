package main

import (
    "errors"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "strings"
)

func main() {
    field := buildSudokuField(os.Args[1])

    missingCells := getMissingCells(field)

    field, err := getMissingCellValues(field, missingCells, 0)
    if err != nil {
	panic("invalid sudoku")
    }

    printField(field)
}

func buildSudokuField(inputFile string) [][]string {
    field := make([][]string, 9, 9)

    file, _ := ioutil.ReadFile(inputFile)

    for y, row := range strings.Split(string(file), "\n") {
	if row != "" {
	    field[y] = strings.Split(row, " ")
	}
    }

    return field
}

func getMissingCells(field [][]string) [][]int {
    missingCells := [][]int{}
    for y, row := range field {
	for x, cell := range row {
	    if cell != "." {
		continue
	    }

	    missingCells = append(missingCells, []int{x, y})
	}
    }

    return missingCells
}

func getMissingCellValues(field [][]string, missingCells [][]int, index int) ([][]string, error) {
    if index >= len(missingCells) {
	return field, nil
    }

    x := missingCells[index][0]
    y := missingCells[index][1]

    for i := 1; i <= 9; i++ {
	value := strconv.Itoa(i)
	if !isValidValue(field, value, x, y) {
	    continue
	}

	field[y][x] = value

	filledField, err := getMissingCellValues(field, missingCells, index + 1)
	if err != nil {
	    field[y][x] = "."
	    continue
	}

	return filledField, nil
    }

    return field, errors.New("No value found")
}

func isValidValue(field [][]string, value string, x int, y int) bool {
    if valueExistsInRow(field, value, y) || valueExistsInColumn(field, value, x) || valueExistsInBlock(field, value, x, y) {
	return false
    }

    return true
}

func valueExistsInRow(field [][]string, value string, row int) bool {
    for x := 0; x < 9; x++ {
	if field[row][x] == value {
	    return true
	}
    }

    return false
}

func valueExistsInColumn(field [][]string, value string, column int) bool {
    for y := 0; y < 9; y++ {
	if field[y][column] == value {
	    return true
	}
    }
    
    return false
}

func valueExistsInBlock(field [][]string, value string, column int, row int) bool {
    hIndex := int(column / 3)
    vIndex := int(row / 3)
    for y := 0; y < 3; y++ {
	for x := 0; x < 3; x++ {
	    if field[y + vIndex * 3][x + hIndex * 3] == value {
		return true
	    }
	}
    }

    return false
}

func printField(field [][]string) {
    for i, row := range field {
	for j, cell := range row {
	    fmt.Printf("%s ", cell)

	    if j % 3 == 2 && j < 8 {
		fmt.Print("| ")
	    }
	}

	fmt.Print("\n")

	if i % 3 == 2 && i < 8 {
	    fmt.Println("------|-------|------")
	}
    }
}
