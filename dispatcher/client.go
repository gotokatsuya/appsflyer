package dispatcher

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/gotokatsuya/appsflyer/util"
)

type (
	Parameter struct {
		APIToken string
		AppID    string
		FromDate string
		ToDate   string
	}
	Client struct {
		HTTPClient *http.Client

		APIBaseURL   string
		APIParameter Parameter
	}
)

const (
	defaultAPIBaseURL = "https://hq.appsflyer.com"
)

func NewClient(appID, fromDate, toDate string) *Client {
	return NewClientWithParam(util.GetAPIToken(), appID, fromDate, toDate)
}

func NewClientWithParam(apiToken, appID, fromDate, toDate string) *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		APIBaseURL: defaultAPIBaseURL,
		APIParameter: Parameter{
			APIToken: apiToken,
			AppID:    appID,
			FromDate: fromDate,
			ToDate:   toDate,
		},
	}
}

func (c *Client) DispatchGetRequest(endpoint string) ([]byte, error) {
	u, err := url.Parse(c.APIBaseURL)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join("export", c.APIParameter.AppID, endpoint)
	urlString := u.String()

	values := url.Values{}
	values.Set("api_token", c.APIParameter.APIToken)
	values.Set("from", c.APIParameter.FromDate)
	values.Set("to", c.APIParameter.ToDate)

	resp, err := c.HTTPClient.Get(urlString + "?" + values.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
