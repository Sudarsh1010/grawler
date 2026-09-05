package links

import (
	"io"
	"net/url"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func Extract(r io.Reader, baseURL string) ([]string, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	htmlNodes, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	seen := make(map[string]struct{})
	dedupeURLs := []string{}

	for n := range htmlNodes.Descendants() {
		if n.Type != html.ElementNode || n.DataAtom != atom.A {
			continue
		}

		for _, a := range n.Attr {
			if a.Key != "href" {
				continue
			}

			ref, err := url.Parse(a.Val)
			if err != nil {
				continue
			}

			u := base.ResolveReference(ref)
			u.Fragment = ""

			if u.Host != base.Host {
				continue
			}

			urlStr := u.String()
			if _, exists := seen[urlStr]; exists {
				continue
			}

			seen[urlStr] = struct{}{}
			dedupeURLs = append(dedupeURLs, urlStr)
		}
	}

	return dedupeURLs, nil
}
