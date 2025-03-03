// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mailchimp provides constants for using OAuth2 to access MailChimp.
package mailchimp // import "github.com/shimt/pam-exec-oauth2/internal/oauth2/mailchimp"

import (
	"github.com/shairozan/pam-exec-oauth2/internal/oauth2"
)

// Endpoint is MailChimp's OAuth 2.0 endpoint.
// See http://developer.mailchimp.com/documentation/mailchimp/guides/how-to-use-oauth2/
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://login.mailchimp.com/oauth2/authorize",
	TokenURL: "https://login.mailchimp.com/oauth2/token",
}
