package termout

import (
	"fmt"
	"geomancy-app/geomancy"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

/* Magic Numbers, common counting patterns */
const (
	zero      = 0
	one       = 1
	two       = 2
	three     = 3
	four      = 4
	minHeight = 30
	minWidth  = 105
)

/* Clears Terminal screen, MAC, Linux and windows*/
func ClearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

/* Center align a UTF-8 String with padding */
func centerUTF8(text string, width int) string {
	runes := []rune(text)
	textLength := len(runes)
	if textLength >= width {
		return text
	}
	padding := (width - textLength) / two
	return strings.Repeat(" ", padding) + text + strings.Repeat(" ", width-textLength-padding)
}

/* Inserts the verticle lines for interior and exterior cells*/
func printVerticleLines(start int, index int, value string) string {
	if index == start {
		return "┃" + value + "┃"
	}
	return value + "┃"
}

/* Plots the points of the figure by binary value one row of points at a time*/
func printRowOfFigures(geo *geomancy.Geomancy, start int, stop, indent int) {

	onePoint := centerUTF8(" ◇ ", indent)
	twoPoint := centerUTF8("◇ ◇", indent)

	for i := range four {
		for j := start; j >= stop; j-- {
			var points string
			if geo.String(j)[i] == '0' {
				points = twoPoint
			} else {
				points = onePoint
			}
			points = printVerticleLines(start, j, points)
			fmt.Printf("%s", points)
		}
		fmt.Printf("\n")
	}
}

/* Centers and prints the figure names in each cell of a row*/
func printFigureNames(geo *geomancy.Geomancy, start int, stop, indent int) {

	for i := range two {
		for j := start; j >= stop; j-- {
			var figure string
			name := strings.Split(geo.Name(j), " ")
			if len(name) == two {
				figure = centerUTF8(name[i], indent)
			} else if len(name) == one && i == one {
				figure = centerUTF8(name[zero], indent)
			} else {
				figure = centerUTF8(" ", indent)
			}

			figure = printVerticleLines(start, j, figure)
			fmt.Printf("%s", figure)
		}
		fmt.Printf("\n")
	}
}

/*Prints the shield reading figures 0 15 with their names and in proper order.*/
func PrintGeomancy(geo *geomancy.Geomancy, rubeus bool) {

	var (
		prev      int = 0
		start     int = four * two
		step      int = four * two
		width     int = four * three
		fullWidth int = minWidth - two
	)

	ClearTerminal()
	line := strings.Repeat(string("━"), fullWidth)
	top := "┏" + line + "┓"
	interior := "┣" + line + "┫"
	bottom := "┗" + line + "┛"

	if rubeus && geo.Name(0) == "Rubeus" {
		fmt.Println("Rubues is the first figure cast.  Reading Destroyed, wait thirty" +
			" \nminutes before consulting again.")
		return
	}

	for i := range four {
		if i == zero {
			fmt.Println(top)
		} else {
			fmt.Println(interior)
		}
		printRowOfFigures(geo, start-one, prev, width)
		printFigureNames(geo, start-one, prev, width)
		step = step / two
		prev = start
		start = step + start
		width = width*two + one
	}
	fmt.Println(bottom)
}

func PrintInerpreter(txt string) {
	fmt.Println(wrapTextPreservingNewlines(txt, minWidth))
}
