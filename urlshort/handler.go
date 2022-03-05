package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		location, ok := pathsToUrls[r.URL.Path]
		if !ok {
			fallback.ServeHTTP(rw, r)
		} else {
			http.Redirect(rw, r, location, http.StatusMovedPermanently)
		}
	}
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
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(ys []byte) ([]map[interface{}]interface{}, error) {
	m := make([]map[interface{}]interface{}, 0)
	err := yaml.Unmarshal(ys, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func buildMap(parsedYaml []map[interface{}]interface{}) map[string]string {
	m := make(map[string]string)
	for _, v := range parsedYaml {
		m[v["path"].(string)] = v["url"].(string)
	}
	return m
}
