package libseat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	hc     *http.Client
	key    string
	server string
	u      url.URL
}

func NewClient(server, key string) (*Client, error) {
	u, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	return &Client{
		hc: &http.Client{
			Transport: &http.Transport{
				IdleConnTimeout: 15 * time.Second,
			},
			Timeout: 15 * time.Second,
		},
		key:    key,
		server: server,
		u:      *u,
	}, nil
}

type requestApi struct {
	Method    string
	Path      string
	URLParams url.Values

	Body   interface{}
	Output interface{}
}

func (s Client) performRequest(a requestApi) error {
	var b io.ReadWriter

	if a.Body != nil {
		b = new(bytes.Buffer)
		if err := json.NewEncoder(b).Encode(&a.Body); err != nil {
			return err
		}
	}

	if a.Method == "" {
		a.Method = "GET"
	}

	u := s.u
	u.Path = a.Path
	if a.URLParams != nil {
		u.RawQuery = a.URLParams.Encode()
	}

	req, err := http.NewRequest(a.Method, u.String(), b)
	if err != nil {
		return err
	}

	req.Header.Add("X-Token", s.key)
	req.Header.Add("Accept", "application/json")

	resp, err := s.hc.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		return fmt.Errorf("libseat: unexpected status code of %d", resp.StatusCode)
	}

	if a.Output != nil {
		if err := json.NewDecoder(resp.Body).Decode(&a.Output); err != nil {
			return err
		}
	}

	return nil
}
