package auspost

import (
	"errors"
	"fmt"
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
		fmt.Println("AusPost API Request")
		fmt.Println("===================")
		fmt.Println()
		fmt.Println("Response Info:")
		fmt.Println("Error      :", err)
		fmt.Println("Status Code:", resp.StatusCode())
		fmt.Println("Status     :", resp.Status())
		fmt.Println("Time       :", resp.Time())
		fmt.Println("Received At:", resp.ReceivedAt())
		fmt.Println("Body       :\n", resp)
		fmt.Println()

		fmt.Println("Request Trace Info:")
		ti := resp.Request.TraceInfo()
		fmt.Println("DNSLookup    :", ti.DNSLookup)
		fmt.Println("ConnTime     :", ti.ConnTime)
		fmt.Println("TLSHandshake :", ti.TLSHandshake)
		fmt.Println("ServerTime   :", ti.ServerTime)
		fmt.Println("ResponseTime :", ti.ResponseTime)
		fmt.Println("TotalTime    :", ti.TotalTime)
		fmt.Println("IsConnReused :", ti.IsConnReused)
		fmt.Println("IsConnWasIdle:", ti.IsConnWasIdle)
		fmt.Println("ConnIdleTime :", ti.ConnIdleTime)
		fmt.Println()
	}

	if err != nil {
		return model.PostcodeSearchResult{}, err
	}

	if r, ok := resp.Result().(model.PostcodeSearchResult); ok {
		return r, nil
	}

	return model.PostcodeSearchResult{}, ErrInvalidResult
}
