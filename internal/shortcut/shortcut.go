package shortcut

import (
	"net/url"
)

type Lookup interface {
	Lookup(string) (*url.URL, error)
}

type Put interface {
	Put(*url.URL) string
}

type ReadWrite interface {
	Lookup
	Put
}

type KeyGenerator interface {
	Generate(key string) string
}

func NewShortcuts(generator KeyGenerator) ReadWrite {
	return &inMemory{
		generator: generator,
		table:     make(map[string]*url.URL),
	}
}
