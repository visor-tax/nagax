package oauth

import (
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

type option func(*Module)

func WithUserStore(u UserStore) option {
	return func(m *Module) {
		m.userStore = u
	}
}

func WithRedirectPath(url string) option {
	return func(m *Module) {
		m.postOAuthRedirectPath = url
	}
}

func WithOAuthConfig(config *oauth2.Config) option {
	return func(m *Module) {
		m.oauthConfig = config
	}
}

func WithAuthCodeOptions(authCodeOptions ...oauth2.AuthCodeOption) option {
	return func(m *Module) {
		m.oauthOptions = authCodeOptions
	}
}

func WithSetState(setOAuthState func(req *http.Request) string) option {
	return func(m *Module) {
		m.setOAuthState = setOAuthState
	}
}

func WithGetCallbackRedirectPath(getCallbackRedirectPath func(userToken, state string) *url.URL) option {
	return func(m *Module) {
		m.getCallbackRedirectPath = getCallbackRedirectPath
	}
}

func WithLoginURL(loginURL string) option {
	return func(m *Module) {
		m.loginURL = loginURL
	}
}

func WithOAuthCallbackPath(p string) option {
	return func(m *Module) {
		m.oauthCallbackPath = p
	}
}
