package termout

import "strings"

func wrapLine(line string, width int) string {
	words := strings.Fields(line)
	var result strings.Builder
	lineLen := 0

	for _, word := range words {
		if lineLen+len(word)+1 > width {
			result.WriteString("\n")
			lineLen = 0
		} else if lineLen > 0 {
			result.WriteString(" ")
			lineLen++
		}
		result.WriteString(word)
		lineLen += len(word)
	}

	return result.String()
}

func wrapTextPreservingNewlines(text string, width int) string {
	lines := strings.Split(text, "\n")
	var result strings.Builder

	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			// Preserve blank lines
			result.WriteString("\n")
			continue
		}
		wrapped := wrapLine(line, width)
		result.WriteString(wrapped)
		if i < len(lines)-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}
