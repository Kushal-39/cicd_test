package cowsay

import (
	"fmt"
	"strings"
)

func WrapText(msg string, width int) string {
	words := strings.Fields(msg)
	if len(words) == 0 {
		return ""
	}

	lines := []string{}
	line := words[0]

	for _, word := range words[1:] {
		if len(line)+1+len(word) > width {
			lines = append(lines, line)
			line = word
		} else {
			line += " " + word
		}
	}
	lines = append(lines, line)
	return strings.Join(lines, "\n")
}

func Say(msg string) string {
	lines := strings.Split(msg, "\n")
	longest := 0
	for _, l := range lines {
		if len(l) > longest {
			longest = len(l)
		}
	}

	top := fmt.Sprintf(" %s\n", strings.Repeat("_", longest+2))
	bottom := fmt.Sprintf(" %s\n", strings.Repeat("-", longest+2))

	mid := ""
	for _, l := range lines {
		mid += fmt.Sprintf("< %s >\n", l)
	}

	cow := `
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
`
	return top + mid + bottom + cow
}
