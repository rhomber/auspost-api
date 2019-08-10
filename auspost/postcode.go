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
	req := c.restClient.R()

	if c.trace {
		req.EnableTrace()
	}

	resp, err := req.
		SetQueryParams(map[string]string{
			"q":                  q,
			"state":              state,
			"excludepostboxflag": strconv.FormatBool(excludePostboxFlag),
		}).
		SetResult(model.PostcodeSearchResult{}).
		Get(c.getUrl(ausPostUriPostcodeSearch))

	if c.trace {
		spewTrace(resp, err)
	}

	if err != nil {
		return model.PostcodeSearchResult{}, err
	}

	if r, ok := resp.Result().(*model.PostcodeSearchResult); ok {
		return *r, nil
	}

	return model.PostcodeSearchResult{}, ErrInvalidResult
}
