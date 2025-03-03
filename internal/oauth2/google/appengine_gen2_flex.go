// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !appengine
// +build !appengine

// This file applies to App Engine second generation runtimes (>= Go 1.11) and App Engine flexible.

package google

import (
	"context"
	"log"
	"sync"

	"github.com/shairozan/pam-exec-oauth2/internal/oauth2"
)

var logOnce sync.Once // only spam about deprecation once

// See comment on AppEngineTokenSource in appengine.go.
func appEngineTokenSource(ctx context.Context, scope ...string) oauth2.TokenSource {
	logOnce.Do(func() {
		log.Print("google: AppEngineTokenSource is deprecated on App Engine standard second generation runtimes (>= Go 1.11) and App Engine flexible. Please use DefaultTokenSource or ComputeTokenSource.")
	})
	return ComputeTokenSource("")
}
