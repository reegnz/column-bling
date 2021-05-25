package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

var (
	defaultColors = []colorFunc{
		color.YellowString,
		color.GreenString,
		color.CyanString,
		color.BlueString,
	}
)

type colorFunc func(string, ...interface{}) string

func readColumnLengths(header string) [][]int {
	// Parse header and find witdth of each header.
	// This regex assumes that each column header is one or more words,
	// And that words are separated by one space at most.
	// Additional assumption is that headers are separated with at least two spaces.
	regex := regexp.MustCompile(`\S+(\s\S+)*\s*`)
	res := regex.FindAllStringIndex(header, -1)
	return res
}

func colorizeColumn(text string, column int) string {
	colorize := defaultColors[column%len(defaultColors)]
	return colorize(text)
}

func colorizeColumns(line string, columnLengths [][]int) string {
	runes := []rune(line)

	var colorizedColumns []string
	// patch columnLengths to length of line
	columnLengths[len(columnLengths)-1][1] = len(runes)
	for i, v := range columnLengths {
		column := string(runes[v[0]:v[1]])
		colorizedColumn := colorizeColumn(column, i)
		colorizedColumns = append(colorizedColumns, colorizedColumn)
	}
	colorizedLine := strings.Join(colorizedColumns, "")
	return colorizedLine
}

func main() {
	color.NoColor = false
	scanner := bufio.NewScanner(os.Stdin)
	var columnLengths [][]int
	if scanner.Scan() {
		header := scanner.Text()
		columnLengths = readColumnLengths(header)
		headerColor := color.New(color.Underline, color.Bold)
		fmt.Println(headerColor.Sprintf(header))
	} else {
		os.Exit(1)
	}
	for scanner.Scan() {
		line := scanner.Text()
		coloredLine := colorizeColumns(line, columnLengths)
		fmt.Println(coloredLine)
	}
}
