package httpSource

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"
)

type HttpSource struct {
	url    string
	client *http.Client
}

func NewHttpSource(url string) *HttpSource {
	return &HttpSource{url: url, client: &http.Client{
		Timeout: 100 * time.Millisecond,
		Transport: &http.Transport{
			MaxConnsPerHost: 0,
		},
	}}
}

func (s *HttpSource) PerformQuery(ammo int, ctx context.Context) error {
	ammoStr := strconv.Itoa(ammo)

	req, err := http.NewRequestWithContext(ctx, "GET", s.url+"?id="+ammoStr, nil)
	if err != nil {
		return err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return nil
}
