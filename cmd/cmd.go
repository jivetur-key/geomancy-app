package cmd

import (
	"fmt"
	"geomancy-app/geomancy"
	"geomancy-app/interpreter"
	"geomancy-app/termout"
	"golang.org/x/term"
	"os"
)

// Enforced required terminal width
const (
	minWidth  = 105
	minHeight = 30
)

// Generates the reading, if rubues is true will exit if that is the first mother.  Will connect to the Gemini AP for
// interpretation if planet parameter is something other then None.
func Cmd(planet string, rubeus bool) error {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))

	if err != nil {
		return fmt.Errorf("Error getting terminal size: %v.", err)

	}
	if width < minWidth || height < minHeight {
		return fmt.Errorf("Error terminal must have at least a width of %d and a height of %d\nCurrent width: %d"+
			"current height: %d\n", minWidth, minHeight, width, height)
	}
	geo := geomancy.New()

	err = geo.Generate()

	if err != nil {
		return err
	}

	termout.PrintGeomancy(geo, rubeus)
	if planet != "None" {
		result, err := interpreter.Interperet(geo, planet)

		if err != nil {
			return err
		}

		termout.PrintInerpreter(result)
	}
	return nil
}
