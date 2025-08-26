package interpreter

import (
	"context"
	"fmt"
	"geomancy-app/geomancy"
	"github.com/google/generative-ai-go/genai"

	"google.golang.org/api/option"
	"os"
)

type reading struct {
	Mothers   []string
	Daughters []string
	Nieces    []string
	Witnesses []string
	Judge     string
}

const (
	mothers   = 4
	daughers  = 8
	nieces    = 12
	witnesses = 14
	judge     = 15
)

var questionTypes = map[string]string{
	"1": "The Sun: Pertaining to life, success, authority, and powerful figures.",
	"2": "The Moon: Pertaining to secrets, the home, family, and emotions.",
	"3": "Mercury: Pertaining to communication, business, learning, and messages.",
	"4": "Venus: Pertaining to love, relationships, art, and social matters.",
	"5": "Mars: Pertaining to conflict, strife, courage, and passion.",
	"6": "Jupiter: Pertaining to wealth, expansion, wisdom, and law.",
	"7": "Saturn: Pertaining to endings, restrictions, karma, and time.",
}

func getReading(geo *geomancy.Geomancy) reading {

	r := reading{}
	for v := range judge {
		if v < mothers {
			r.Mothers = append(r.Mothers, geo.Name(v))
		} else if v >= mothers && v < daughers {
			r.Daughters = append(r.Daughters, geo.Name(v))
		} else if v >= daughers && v < nieces {
			r.Nieces = append(r.Nieces, geo.Name(v))
		} else if v >= nieces && v < witnesses {
			r.Witnesses = append(r.Witnesses, geo.Name(v))
		} else {
			r.Judge = geo.Name(v)
		}
	}

	return r
}

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

	// --- Connect to the Gemini API ---
	ctx := context.Background()

	// The client will automatically use the GEMINI_API_KEY environment variable.
	// If you need to pass the key directly, use option.WithAPIKey("YOUR_API_KEY")
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return "", err
	}
	defer client.Close()

	// --- Select the Gemini Model ---
	// For text-based tasks, "gemini-1.5-flash-latest" is a good choice.
	model := client.GenerativeModel("gemini-1.5-flash-latest")

	// --- Send the Prompt and Get the Response ---
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}
	result := ""
	// --- Print the Interpretation ---
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
