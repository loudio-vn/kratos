package oidc

import (
	_ "embed"
)

//go:embed .schema/link.schema.json
var linkSchema []byte

//go:embed .schema/login.schema.json
var loginSchema []byte
