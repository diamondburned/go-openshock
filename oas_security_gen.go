// Code generated by ogen, DO NOT EDIT.

package openshock

import (
	"context"
	"net/http"

	"github.com/go-faster/errors"
)

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// OpenShockToken provides OpenShockToken security value.
	// API Token Authorization header.
	OpenShockToken(ctx context.Context, operationName string) (OpenShockToken, error)
}

func (s *Client) securityOpenShockToken(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.OpenShockToken(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"OpenShockToken\"")
	}
	req.Header.Set("OpenShockToken", t.APIKey)
	return nil
}
