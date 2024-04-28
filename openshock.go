// Package openshock provides the generated client code for the OpenShock API.
package openshock

import "context"

// APIToken is a security value for OpenShockToken security scheme.
// It implements the [SecuritySource] interface.
type APIToken string

var _ SecuritySource = APIToken("")

// OpenShockToken implements the [SecuritySource] interface.
func (t APIToken) OpenShockToken(ctx context.Context, operationName string) (OpenShockToken, error) {
	return OpenShockToken{APIKey: string(t)}, nil
}
