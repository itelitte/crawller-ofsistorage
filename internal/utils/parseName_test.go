package utils

import (
	"fmt"
	"testing"
)

func TestExtractValues(t *testing.T) {
	t.Run("Teste ExtractValues", func(t *testing.T) {
		input := "ABBASIN 1: ABDUL AZIZ 2: n/a 3: n/a 4: n/a 5: n/a."
		output := ExtractValues(input)

		fmt.Println(output)
	})
}
