package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/sudarsh1010/grawler/internal/links"
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

	parsedURLs, err := links.Extract(resp.Body, base.String())
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range parsedURLs {
		fmt.Println(v)
	}
}
