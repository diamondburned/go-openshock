// Code generated by ogen, DO NOT EDIT.

package openshock

import (
	"net/http"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenregex"
)

var regexMap = map[string]ogenregex.Regexp{
	"^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$":                                                                                        ogenregex.MustCompile("^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$"),
	"^[^ \f\n\r\t\v\u00a0\u1680\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200a\u2028\u2029\u202f\u205f\u3000\ufeff][^\r\n\u2028\u2029]*[^ \f\n\r\t\v\u00a0\u1680\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200a\u2028\u2029\u202f\u205f\u3000\ufeff]$": ogenregex.MustCompile("^[^ \f\n\r\t\v\u00a0\u1680\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200a\u2028\u2029\u202f\u205f\u3000\ufeff][^\r\n\u2028\u2029]*[^ \f\n\r\t\v\u00a0\u1680\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200a\u2028\u2029\u202f\u205f\u3000\ufeff]$"),
}

type (
	optionFunc[C any] func(*C)
)

type clientConfig struct {
	Client ht.Client
}

// ClientOption is client config option.
type ClientOption interface {
	applyClient(*clientConfig)
}

var _ ClientOption = (optionFunc[clientConfig])(nil)

func (o optionFunc[C]) applyClient(c *C) {
	o(c)
}

func newClientConfig(opts ...ClientOption) clientConfig {
	cfg := clientConfig{
		Client: http.DefaultClient,
	}
	for _, opt := range opts {
		opt.applyClient(&cfg)
	}
	return cfg
}

type baseClient struct {
	cfg clientConfig
}

func (cfg clientConfig) baseClient() (c baseClient, err error) {
	c = baseClient{cfg: cfg}
	return c, nil
}

// Option is config option.
type Option interface {
	ClientOption
}

// WithClient specifies http client to use.
func WithClient(client ht.Client) ClientOption {
	return optionFunc[clientConfig](func(cfg *clientConfig) {
		if client != nil {
			cfg.Client = client
		}
	})
}