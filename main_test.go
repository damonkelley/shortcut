package main

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"
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

func TestRedirectMe(t *testing.T) {

	database := &TestDatabase{
		redirections: map[string]*url.URL{
			"/abc123": urlFrom("http://subdomain.example.com"),
		},
	}

	t.Run("the status code is 301 when the key is found", func(t *testing.T) {
		request := httptest.NewRequest("GET", "http://example.com/abc123", nil)
		recorder := httptest.NewRecorder()

		RedirectMe(database).ServeHTTP(recorder, request)

		expected := 301

		if recorder.Code != expected {
			t.Errorf("Expected %d but got %d", expected, recorder.Code)
		}
	})

	t.Run("it redirects the request when the key is found", func(t *testing.T) {
		request := httptest.NewRequest("GET", "http://example.com/abc123", nil)
		recorder := httptest.NewRecorder()

		RedirectMe(database).ServeHTTP(recorder, request)

		location, err := recorder.Result().Location()
		expectedLocation, _ := database.Lookup("/abc123")

		if err != nil {
			t.Fatal("Location header was not set")
		}

		if *location != *expectedLocation {
			t.Errorf("Expected %s but got %s", expectedLocation, location)
		}
	})

	t.Run("it is not found if the key doesn't exist", func(t *testing.T) {
		request := httptest.NewRequest("GET", "http://example.com/notfound", nil)
		recorder := httptest.NewRecorder()

		RedirectMe(database).ServeHTTP(recorder, request)

		actual, expected := recorder.Code, 404

		if expected != actual {
			t.Errorf("Expected %d but got %d", expected, actual)
		}
	})
}

func TestAdmin(t *testing.T) {
	t.Run("a new link can be added", func(t *testing.T) {
	})
}

func urlFrom(raw string) *url.URL {
	parsed, _ := url.Parse(raw)
	return parsed
}
