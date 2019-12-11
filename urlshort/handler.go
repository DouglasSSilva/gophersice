package urlshort

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// RedirectPaths to be converted when receiven a file (.yaml, .json) or when receiveng from the db
type RedirectPaths struct {
	Path string `yaml:"- path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	redirect := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(pathsToUrls[r.URL.Path]) > 0 {
			http.Redirect(w, r, pathsToUrls[r.URL.Path], http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(w, r)
		}
	})

	return redirect
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	rPaths, err := convertYAML(yml)
	if err != nil {
		return nil, err
	}

	mapped := buildMap(rPaths)

	return MapHandler(mapped, fallback), nil
}

func convertYAML(yml []byte) ([]RedirectPaths, error) {
	var rPaths []RedirectPaths
	err := yaml.Unmarshal(yml, &rPaths)
	if err != nil {
		return rPaths, err
	}
	return rPaths, err
}

func buildMap(rPaths []RedirectPaths) map[string]string {
	var mapped map[string]string

	for i := range rPaths {
		mapped[rPaths[i].Path] = rPaths[i].URL
	}
	_, found := mapped[rPaths[0].Path]
	fmt.Println(found)
	return mapped
}
