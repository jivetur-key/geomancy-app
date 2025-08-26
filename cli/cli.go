package cli

import (
	"flag"
	"fmt"
	"strings"
)

// Flags if planet flag is provided the value will be used for Gemini reuest
// If Rubeus is set to false readings will be generated if Rubeus is the first Mother Figure.
var (
	Planet string
	Rubeus bool
)

// Planet descriptions are provided for question types.
func InitFlags() {
	flag.StringVar(&Planet, "planet", "", " The Sun: Pertaining to life, success, authority, and powerful figures.\n"+
		" The Moon: Pertaining to secrets, the home, family, and emotions.\n"+
		" Mercury: Pertaining to communication, business, learning, and messages.\n"+
		" Venus: Pertaining to love, relationships, art, and social matters.\n"+
		" Mars: Pertaining to conflict, strife, courage, and passion.\n"+
		" Jupiter: Pertaining to wealth, expansion, wisdom, and law.\n"+
		" Saturn: Pertaining to endings, restrictions, karma, and time.\n")
	flag.BoolVar(&Rubeus, "rubeus", true, "When enabled readings are aborted if the First Mother is Rubeus.")
}

// planet numbers are what Gemini uses instead of names so names need to converted to numbers.
var planetNumber = map[string]string{
	"sun":     "1",
	"moon":    "2",
	"mercury": "3",
	"venus":   "4",
	"mars":    "5",
	"jupiter": "6",
	"saturn":  "7",
}

// Verify if String flag is set.
func IsFlagSet(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

// Converts planet name to number.
func GetPlanetNumber() (string, error) {
	key := strings.ToLower(Planet)

	val, ok := planetNumber[key]

	if !ok {
		return "", fmt.Errorf("Error: %s is an invailid argument.", val)
	}

	return val, nil
}
