package UrlParameters

import (
	"fmt"
	"regexp"
	"strings"
)

func CreatePath(rawPath string) *Path {
	pathPieces := strings.Split(strings.TrimLeft(rawPath, "/"), "/")
	path := Path{
		RawPath: rawPath,
	}
	pattern := ""
	for i, pathPiece := range pathPieces {
		fmt.Println(pathPiece)
		if a, _ := regexp.MatchString("{.*}", pathPiece); a {
			param := Parameter{
				Name: strings.TrimRight(strings.TrimLeft(pathPiece, "{"), "}"),
				Index: i,
			}
			path.Parameters = append(path.Parameters, param)

			pattern += "/.*"
		} else {
			pattern += "/" + pathPiece
		}
	}

	path.Regex = regexp.MustCompile(pattern)
	return &path
}