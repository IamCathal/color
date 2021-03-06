package color

import (
	"fmt"
	"strings"
)

func StyledText(msg string, modifiers ...string) string {
	result := msg
	if len(modifiers) == 0 {
		return msg
	}

	modifierMap := map[string]string{
		"resetall":               "\033[0m",
		"bold":                   "\033[1m",
		"dim":                    "\033[2m",
		"underlined":             "\033[4m",
		"blink":                  "\033[5m",
		"reverse":                "\033[7m",
		"hidden":                 "\033[8m",
		"resetbold":              "\033[21m",
		"resetdim":               "\033[22m",
		"resetunderlined":        "\033[24m",
		"resetblink":             "\033[25m",
		"resetreverse":           "\033[27m",
		"resethidden":            "\033[28m",
		"default":                "\033[39m",
		"black":                  "\033[30m",
		"red":                    "\033[31m",
		"green":                  "\033[32m",
		"yellow":                 "\033[33m",
		"blue":                   "\033[34m",
		"magenta":                "\033[35m",
		"cyan":                   "\033[36m",
		"lightgray":              "\033[37m",
		"darkgray":               "\033[90m",
		"lightred":               "\033[91m",
		"lightgreen":             "\033[92m",
		"lightyellow":            "\033[93m",
		"lighblue":               "\033[94m",
		"lightmagenta":           "\033[95m",
		"lightcyan":              "\033[96m",
		"white":                  "\033[97m",
		"backgrounddefault":      "\033[49m",
		"backgroundblack":        "\033[40m",
		"backgroundred":          "\033[41m",
		"backgroundgreen":        "\033[42m",
		"backgroundyellow":       "\033[43m",
		"backgroundnlue":         "\033[44m",
		"backgroundmagenta":      "\033[45m",
		"backgroundcyan":         "\033[46m",
		"backgroundlightgray":    "\033[47m",
		"backgrounddarkgray":     "\033[100m",
		"backgroundlightred":     "\033[101m",
		"backgroundlightgreen":   "\033[102m",
		"backgroundlightyellow":  "\033[103m",
		"backgroundlightblue":    "\033[104m",
		"backgroundlightmagenta": "\033[105m",
		"backgroundlightcyan":    "\033[106m",
		"backgroundwhite":        "\033[107m",
	}

	for _, modifier := range modifiers {
		lowerCaseModifier := strings.ToLower(modifier)
		if _, exists := modifierMap[strings.ToLower(lowerCaseModifier)]; exists {
			result = fmt.Sprintf("%s%s%s", modifierMap[lowerCaseModifier], result, modifierMap["resetall"])
		}
	}

	return result
}
