package shortcut_test

import (
	"net/url"
	"testing"

	"com.damonkelley/shortcut/internal/shortcut"
)

type TestKeyGenerator struct{}

func (generator *TestKeyGenerator) Generate(key string) string {
	return key
}

func TestInMemoryShortcuts(t *testing.T) {
	t.Run("a url can be added", func(t *testing.T) {
		urlToAdd, _ := url.Parse("https://example.com")

		shortcuts := shortcut.NewShortcuts(&TestKeyGenerator{})

		key := shortcuts.Put(urlToAdd)

		foundUrl, _ := shortcuts.Lookup(key)

		if foundUrl.String() != urlToAdd.String() {
			t.Errorf("Expected to find %s but got %s", urlToAdd, foundUrl)
		}
	})

	t.Run("it will generate a key", func(t *testing.T) {
		urlToAdd, _ := url.Parse("https://example.com")

		generator := &TestKeyGenerator{}

		key := shortcut.NewShortcuts(generator).
			Put(urlToAdd)

		expectedKey := generator.Generate(urlToAdd.String())

		if expectedKey != key {
			t.Errorf("Expected to have key %q but got %q", expectedKey, key)
		}
	})

	t.Run("looking up a key that doesn't exist with be an error", func(t *testing.T) {
		generator := &TestKeyGenerator{}

		_, err := shortcut.NewShortcuts(generator).
			Lookup("does not exist")

		if err == nil {
			t.Error("Expected Lookup error to be present, but was not")
		}
	})
}
