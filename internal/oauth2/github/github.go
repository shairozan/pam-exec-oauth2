// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package github provides constants for using OAuth2 to access Github.
package github // import "github.com/shimt/pam-exec-oauth2/internal/oauth2/github"

import (
	"github.com/shairozan/pam-exec-oauth2/internal/oauth2"
)

// Endpoint is Github's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://github.com/login/oauth/authorize",
	TokenURL: "https://github.com/login/oauth/access_token",
}
