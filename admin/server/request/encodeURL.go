package request

import (
	"net/url"
)

// EncodeURL add and encode parameters.
func EncodeURL(baseURL string, p Parameters) (string, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	query := url.Query()

	for k := range p {
		query.Set(k, p[k])
	}

	url.RawQuery = query.Encode()
	urlPath := url.String()

	return urlPath, nil
}
