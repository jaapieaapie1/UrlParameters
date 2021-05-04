package UrlParameters

import (
	"regexp"
	"strings"
)

type Path struct {
	RawPath string
	Regex *regexp.Regexp
	Parameters []Parameter
}

type Parameter struct {
	Name string
	Index int
}

func (p Path) Matches(path string) bool  {
	return p.Regex.MatchString(path)
}

func (p Path) Parse(path string) map[string]string {
	pathPieces := strings.Split(strings.TrimLeft(path, "/"), "/")

	values := make(map[string]string)
	for _, param := range p.Parameters {
		values[param.Name] = pathPieces[param.Index]
	}
	return values
}