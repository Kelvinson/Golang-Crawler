package parser

import (
	"github.com/kelvinson/go/crawler/engine"
	"github.com/kelvinson/go/crawler/model"
	"regexp"
)

var position = regexp.MustCompile(`<div class="m-btn pink" [^>]*>Position:([^<]+)</div>`)

func ParseProfile(
	contents []byte, CompanyName string) engine.ParseResult {
	profile := model.Profile{}

	profile.CompanyName = CompanyName
	profile.Position = extractString(
		contents, position)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(
	contents []byte, re *regexp.Regexp) string {

	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return "aaa"
	}
}
