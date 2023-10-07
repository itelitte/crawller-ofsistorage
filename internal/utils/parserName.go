package utils

import (
	"regexp"
	"strings"
)

// input  ABBASIN 1: ABDUL AZIZ 2: n/a 3: n/a 4: n/a 5: n/a.
// output  ABBASIN  , ABDUL AZIZ
//func ExtractValues(input string) string {
//	//remove numbers and n\a
//	in := strings.ReplaceAll(input, "n/a", "")
//	in = strings.ReplaceAll(in, "1", "")
//	in = strings.ReplaceAll(in, "2", "")
//	in = strings.ReplaceAll(in, "3", "")
//	in = strings.ReplaceAll(in, "4", "")
//	in = strings.ReplaceAll(in, "5", "")
//
//	slice := strings.Split(in, ":")
//	var output string
//	for _, n := range slice {
//		if len(n) > 3 {
//			output = output + "," + n
//		}
//	}
//
//	output = strings.Replace(output, ",", "", 1)
//
//	return output
//}

func ExtractValues(input string) string {
	// Remove "n/a" e números usando uma expressão regular
	re := regexp.MustCompile(`\d|n/a`)
	in := re.ReplaceAllString(input, "")

	// Remove espaços extras
	in = strings.Join(strings.Fields(in), " ")

	parts := strings.Split(in, ":")
	var output []string

	for _, part := range parts {
		trimmedPart := strings.TrimSpace(part)
		if len(trimmedPart) > 0 {
			output = append(output, trimmedPart)
		}
	}

	if len(output) > 0 {
		lastItem := output[len(output)-1]
		if lastItem == "." || lastItem == "," {
			output = output[:len(output)-1]
		}
	}

	return strings.Join(output, ", ")
}
