package shortcut

import "net/url"

type Database interface {
	Lookup(string) (*url.URL, error)
}
