package parser

import (
	"github.com/kelvinson/go/crawler/engine"
	"regexp"
)

const cityRe = `<a href="/cmp?l=([0-9a-z=+%]+)"[^>]*>([^<]+)</a>`

func ParseCity(
	contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {

		CompanyName := string(m[2])
		result.Items = append(
			result.Items, "User "+CompanyName)

		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				ParserFunc: func(c []byte) engine.ParseResult {
					return ParseProfile(c, CompanyName)
				},
			})
	}
	return result
}
