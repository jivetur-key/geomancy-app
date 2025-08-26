package interpreter

import (
	"context"
	"fmt"
	"geomancy-app/geomancy"
	"github.com/google/generative-ai-go/genai"

	"google.golang.org/api/option"
	"os"
)

// Struct to split reading into parts for Gemini.
type reading struct {
	Mothers   []string
	Daughters []string
	Nieces    []string
	Witnesses []string
	Judge     string
}

// Number of Figures per group
const (
	mothers   = 4
	daughters = mothers + 4
	nieces    = daughters + 4
	witnesses = nieces + 2
	judge     = witnesses + 1
)

// Planet list repeated for Gemini
var questionTypes = map[string]string{
	"1": "The Sun: Pertaining to life, success, authority, and powerful figures.",
	"2": "The Moon: Pertaining to secrets, the home, family, and emotions.",
	"3": "Mercury: Pertaining to communication, business, learning, and messages.",
	"4": "Venus: Pertaining to love, relationships, art, and social matters.",
	"5": "Mars: Pertaining to conflict, strife, courage, and passion.",
	"6": "Jupiter: Pertaining to wealth, expansion, wisdom, and law.",
	"7": "Saturn: Pertaining to endings, restrictions, karma, and time.",
}

// Creates reading struct for Gemini based on generated geomancy.
func getReading(geo *geomancy.Geomancy) reading {

	r := reading{}
	for v := range judge {
		if v < mothers {
			r.Mothers = append(r.Mothers, geo.Name(v))
		} else if v >= mothers && v < daughters {
			r.Daughters = append(r.Daughters, geo.Name(v))
		} else if v >= daughters && v < nieces {
			r.Nieces = append(r.Nieces, geo.Name(v))
		} else if v >= nieces && v < witnesses {
			r.Witnesses = append(r.Witnesses, geo.Name(v))
		} else {
			r.Judge = geo.Name(v)
		}
	}

	return r
}

// Requests the interperation from Gemini, requires valid API key set as ENV variable.
func Interperet(geo *geomancy.Geomancy, planet string) (string, error) {

	r := getReading(geo)
	questionContext := questionTypes[planet]
	// --- Construct the Prompt for Gemini ---
	prompt := fmt.Sprintf(`
		Interpret the following geomancy shield reading.

		The question is of a specific nature, related to: "%s".
		Please focus your entire analysis through the lens of this topic.
		Provide a detailed analysis of the overall situation, the past, present, and future,
		and the final outcome, all as it pertains to the chosen context.

		Mothers: %v
		Daughters: %v
		Nieces: %v
		Witnesses: %v
		Judge: %s
	`, questionContext, r.Mothers, r.Daughters, r.Nieces, r.Witnesses, r.Judge)

	fmt.Println("--- Sending request to Gemini for interpretation... ---")

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return "", err
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-1.5-flash-latest")

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}
	result := ""

	fmt.Println("--- Gemini's Interpretation ---")
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				if txt, ok := part.(genai.Text); ok {
					result += string(txt)
				}
			}
		}
	}

	return result, nil
}
