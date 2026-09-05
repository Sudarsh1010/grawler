package links

import (
	"reflect"
	"strings"
	"testing"
)

func TestExtract(t *testing.T) {
	baseURL := "https://example.com"

	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "tracer",
			input: `<a href="https://example.com/about">`,
			want:  []string{"https://example.com/about"},
		},

		{
			name:  "relative",
			input: `<a href="/blog">`,
			want:  []string{"https://example.com/blog"},
		},
		{
			name:  "junks",
			input: `<a href="https://example.com#about"><a href="mailto:a@b.c"><a>no href</a>`,
			want:  []string{"https://example.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			got, err := Extract(r, baseURL)
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v want %v", got, tt.want)
			}
		})
	}
}
