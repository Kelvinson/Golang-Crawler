package parser

import (
	"github.com/kelvinson/go/crawler/engine"
	"regexp"
)
https://www.indeed.com/cmp?l=Houston%2C+TX
//<a href="/cmp?l=New+York%2C+NY" data-tn-link="">Best Companies in New York, NY</a>
const cityListRe = `<a href="/cmp?l=([0-9a-z=+%]+)"[^>]*>([^<]+)</a>`

func ParseCityList(
	contents []byte) engine.ParseResult {

	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	//limit := 10

	for _, m := range matches {

		result.Items = append(
			result.Items, "City "+string(m[2]))

		result.Requests = append(
			result.Requests, engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
		//	limit--
		//	if limit == 0 {
		//		break
		//	}
	}
	return result
}
