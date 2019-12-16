package oauth

import (
	"bytes"
	"net/url"
	"strings"
)

type Config struct {
	// ClientKey is the application's ID.
	ClientKey string

	// ClientSecret is the application's secret.
	ClientSecret string
	Endpoint     Endpoint
	// RedirectURL is the URL to redirect users going through
	// the OAuth flow, after the resource owner's URLs.
	RedirectURL string

	// Scope specifies optional requested permissions.
	Scopes string
}
type Endpoint struct {
	// oauth url
	AuthURL string
	// token url
	TokenURL string
	// Refresh Token URL
	RefreshTokenURL string
	// client token url
	ClientTokenURL string
}

func (c *Config) AuthCodeURL(state string) string {
	var buf bytes.Buffer
	buf.WriteString(c.Endpoint.AuthURL)
	v := url.Values{
		"response_type": {responseTypeCode},
		"client_key":    {c.ClientKey},
	}
	if c.RedirectURL != "" {
		v.Set("redirect_uri", c.RedirectURL)
	}
	if len(c.Scopes) > 0 {
		v.Set("scope", c.Scopes)
	}
	if state != "" {
		// TODO(light): Docs say never to omit state; don't allow empty.
		v.Set("state", state)
	}
	if strings.Contains(c.Endpoint.AuthURL, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	return buf.String()
}
