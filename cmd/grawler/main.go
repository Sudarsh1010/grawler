package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: grawler <url>")
		os.Exit(1)
	}

	base, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	httpClient := &http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", base.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d", resp.StatusCode)
	}

	htmlNode, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	seen := make(map[string]struct{})
	parsedURLs := []string{}

	for node := range htmlNode.Descendants() {
		if node.Type != html.ElementNode || node.DataAtom != atom.A {
			continue
		}

		for _, a := range node.Attr {
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
			parsedURLs = append(parsedURLs, urlStr)
		}
	}

	for _, v := range parsedURLs[:] {
		fmt.Println(v)
	}
}
