package homefacts

import (
	"fmt"
	"strings"

	"log"
	"net/http"
	"net/url"
	"os"
)

//curl  --data "&fulladdress=22280+S+209th+Way%2C+Queen+Creek%2C+AZ+85142" --compressed 'http://www.homefacts.com/hfreport.html'

const (
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.95 Safari/537.36"
	baseURL   = "http://www.homefacts.com/hfreport.html"
)

// Client is the GreatSchools client. It contains all the different resources available.
type Client struct {
	Debug  bool
	parser *parser
}

// New creates a new GreatSchools client
func New() *Client {
	return &Client{
		Debug:  false,
		parser: &parser{},
	}
}

// Request contains information to sent to the api endpoint
type Request struct {
	Address string
	Proxy   string
}

// Response contains Results from the API request
type Response struct {
	Result *Result `json:"result"`
}

// Result directly corresponds to the JSON returned by the API
type Result struct {
	CrimeRate            *resultProperty `json:"crimeRate"`
	SchoolRating         *resultProperty `json:"schoolRating"`
	RegisteredOffenders  *resultProperty `json:"registeredOffenders"`
	AvgHomePrice         *resultProperty `json:"avgHomePrice"`
	Foreclosures         *resultProperty `json:"foreclosures"`
	NaturalHazards       *resultProperty `json:"naturalHazards"`
	EnvironmentalHazards *resultProperty `json:"environmentalHazards"`
}

// resultProperty contains the class (related to the icon, indicates severity) and some free text
type resultProperty struct {
	Class string
	Text  string
}

// GetFacts fetches facts from homefacts.com
func (c *Client) GetFacts(r *Request) (*Response, error) {
	resp, err := c.getWebPageData(r)

	if err != nil {
		return nil, err
	}

	result, err := c.parser.parseHTML(resp)

	if err != nil {
		return nil, err
	}

	return &Response{
		Result: result,
	}, nil
}

// getWebPageData fetches facts from homefacts.com
func (c *Client) getWebPageData(r *Request) (*http.Response, error) {
	if len(r.Address) == 0 {
		return nil, fmt.Errorf("must provide full address")
	}

	client := &http.Client{}

	if len(r.Proxy) > 0 {
		_, err := url.Parse(r.Proxy)
		if err == nil {
			os.Setenv("HTTP_PROXY", r.Proxy)
		}
	}

	form := url.Values{}
	form.Add("fulladdress", r.Address)

	sBody := fmt.Sprintf("fulladdress=%s", r.Address)

	if c.Debug {
		log.Printf("post body: %s", sBody)
	}

	body := strings.NewReader(sBody)

	req, err := http.NewRequest("POST", baseURL, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Fatalln(err)
	}

	return client.Do(req)
}
