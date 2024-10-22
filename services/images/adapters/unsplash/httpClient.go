package unsplash

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	apperrors "telegraminput/lib/errors"
	"time"
)

type httpClient struct {
	client        *http.Client
	unsplashToken string
}

func newHTTPClient(token string) httpClient {
	return httpClient{
		unsplashToken: token,
		client: &http.Client{
			Timeout:       time.Minute,
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
		},
	}
}

func (a httpClient) Do(req *http.Request) ([]byte, error) {
	baseURL := "https://api.unsplash.com"

	var err error

	req.URL, err = url.Parse(baseURL + req.URL.String())
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	q := req.URL.Query()
	q.Add("client_id", a.unsplashToken)
	req.URL.RawQuery = q.Encode()

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send http request: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	desiredCode := 200
	if resp.StatusCode != desiredCode {
		return nil, fmt.Errorf("%w: %s", apperrors.ErrExternal, body)
	}

	return body, nil
}
