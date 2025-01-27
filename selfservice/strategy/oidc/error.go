package oidc

import "github.com/ory/herodot"

var (
	ErrScopeMissing = herodot.ErrBadRequest.
			WithError("authentication failed because a required scope was not granted").
			WithReasonf(`Unable to finish because one or more permissions were not granted. Please retry and accept all permissions.`)

	ErrIDTokenMissing = herodot.ErrBadRequest.
				WithError("authentication failed because id_token is missing").
				WithReasonf(`Authentication failed because no id_token was returned. Please accept the "openid" permission and try again.`)

	ErrApiTokenMissing = herodot.ErrBadRequest.
				WithError("request failed because properties id_token and access_token are missing").
				WithReasonf(`Request failed because properties id_token and access_token are missing.`)

	ErrProviderNoAPISupport = herodot.ErrBadRequest.
				WithError("request failed because oidc provider does not implement API flows").
				WithReasonf(`Request failed because oidc provider does not implement API flows.`)

	ErrAPIFlowNotSupported = herodot.ErrBadRequest.WithError("API-based flows are not supported for this method").
				WithReasonf("Social Sign In and OpenID Connect are only supported for flows initiated using the Browser endpoint.")
)
