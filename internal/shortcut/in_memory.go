package shortcut

import (
	"fmt"
	"net/url"
)

type inMemory struct {
	generator KeyGenerator
	table     map[string]*url.URL
}

func (shortcuts *inMemory) Lookup(key string) (*url.URL, error) {
	url, ok := shortcuts.table[key]

	if !ok {
		return url, fmt.Errorf("No URL exists for key %s", key)
	}

	return url, nil
}

func (shortcuts *inMemory) Put(url *url.URL) string {
	key := shortcuts.generator.Generate(url.String())

	shortcuts.table[key] = url

	return key
}
