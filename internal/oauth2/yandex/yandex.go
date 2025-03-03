// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package yandex provides constants for using OAuth2 to access Yandex APIs.
package yandex // import "github.com/shimt/pam-exec-oauth2/internal/oauth2/yandex"

import (
	"github.com/shairozan/pam-exec-oauth2/internal/oauth2"
)

// Endpoint is the Yandex OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://oauth.yandex.com/authorize",
	TokenURL: "https://oauth.yandex.com/token",
}
