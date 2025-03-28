package pathheader

import (
	"context"
	"net/http"
)

type Config struct{}

func CreateConfig() *Config {
	return &Config{}
}

type PathHeader struct {
	next http.Handler
	name string
}

func New(_ context.Context, next http.Handler, _ *Config, name string) (http.Handler, error) {
	return &PathHeader{next: next, name: name}, nil
}

func (p *PathHeader) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Set raw path header
	req.Header.Set("X-Raw-Path", req.URL.Path)

	// Create normalized path by replacing numeric segments with *
	normalizedPath := req.URL.Path
	segments := strings.Split(normalizedPath, "/")
	for i, segment := range segments {
		if _, err := strconv.Atoi(segment); err == nil {
			segments[i] = "*"
		}
	}
	normalizedPath = strings.Join(segments, "/")

	// Set normalized path header
	req.Header.Set("X-Normalized-Path", normalizedPath)

	p.next.ServeHTTP(rw, req)
}