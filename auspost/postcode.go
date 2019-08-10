package auspost

import (
	"errors"
	"github.com/rhomber/auspost-postcode/auspost/model"
	"strconv"
)

var ErrInvalidResult = errors.New("invalid result returned")

const (
	ausPostUriPostcodeSearch = "postcode/search.json"
)

func (c *Client) PostcodeSearch(q string, state string, excludePostboxFlag bool) (model.PostcodeSearchResult, error) {
	resp, err := c.restClient.R().
		SetQueryParams(map[string]string{
			"q":                  q,
			"state":              state,
			"excludepostboxflag": strconv.FormatBool(excludePostboxFlag),
		}).
		SetResult(model.PostcodeSearchResult{}).
		Get(c.getUrl(ausPostUriPostcodeSearch))

	if err != nil {
		return model.PostcodeSearchResult{}, err
	}

	if r, ok := resp.Result().(model.PostcodeSearchResult); ok {
		return r, nil
	}

	return model.PostcodeSearchResult{}, ErrInvalidResult
}
