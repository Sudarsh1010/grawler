package links

import (
	"slices"
	"strings"
	"testing"
)

func TestTracer(t *testing.T) {
	baseURL := "https://example.com"
	want := []string{"https://example.com/about"}

	s := `<a href="https://example.com/about">`
	r := strings.NewReader(s)

	resp, err := Extract(r, baseURL)
	if err != nil {
		t.Error(err)
	}

	if slices.Compare(resp, want) != 0 {
		t.Errorf(`Extract(r, %s) = %q, %v, want match for %#q, nil`, baseURL, resp, err, want)
	}
}

func TestResolveRelative(t *testing.T) {
	baseURL := "https://example.com"
	want := []string{"https://example.com/blog"}

	s := `<a href="/blog">`
	r := strings.NewReader(s)

	resp, err := Extract(r, baseURL)
	if err != nil {
		t.Error(err)
	}

	if slices.Compare(resp, want) != 0 {
		t.Errorf(`Extract(r, %s) = %q, %v, want match for %#q, nil`, baseURL, resp, err, want)
	}
}

func TestJunkIgnoreJunks(t *testing.T) {
	baseURL := "https://example.com"
	want := []string{"https://example.com"}

	s := `<a href="https://example.com#about"><a href="mailto:a@b.c"><a>no href</a>`
	r := strings.NewReader(s)

	resp, err := Extract(r, baseURL)
	if err != nil {
		t.Error(err)
	}

	if slices.Compare(resp, want) != 0 {
		t.Errorf(`Extract(r, %s) = %q, %v, want match for %#q, nil`, baseURL, resp, err, want)
	}
}
