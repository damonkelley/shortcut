package shortcut_test

import (
	"testing"

	"com.damonkelley/shortcut/internal/shortcut"
)

func TestRandomKeyGenerator(t *testing.T) {
	t.Run("it will generate a psuedorandom key", func(t *testing.T) {
		generator := shortcut.NewRandomKeyGenerator(&shortcut.RandomKeyConfig{
			Length: 10,
			Seed:   100,
		})

		input := "Not used for random keys"

		actual := generator.Generate(input)
		expected := "pqKqEKDzWf"

		if actual != expected {
			t.Errorf("Expected key %s but got %s", expected, actual)
		}
	})
}
