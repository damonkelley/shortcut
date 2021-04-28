package main

import (
	"fmt"
	"net/url"
)

type TestDatabase struct {
	redirections map[string]*url.URL
}

func (db *TestDatabase) Lookup(key string) (*url.URL, error) {
	redirectTo, ok := db.redirections[key]

	if !ok {
		return redirectTo, fmt.Errorf("No redirection for %s", key)
	}

	return redirectTo, nil
}

func urlFrom(raw string) *url.URL {
	parsed, _ := url.Parse(raw)
	return parsed
}
