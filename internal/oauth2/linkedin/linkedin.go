// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package linkedin provides constants for using OAuth2 to access LinkedIn.
package linkedin // import "github.com/shimt/pam-exec-oauth2/internal/oauth2/linkedin"

import (
	"github.com/shairozan/pam-exec-oauth2/internal/oauth2"
)

// Endpoint is LinkedIn's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:   "https://www.linkedin.com/oauth/v2/authorization",
	TokenURL:  "https://www.linkedin.com/oauth/v2/accessToken",
	AuthStyle: oauth2.AuthStyleInParams,
}
