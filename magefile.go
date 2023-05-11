//go:build mage
// +build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/sh"
)

func Linux() error {
	return build("amd64", "linux")
}

func build(arch, os string) error {
	return sh.RunV("go", "build", "-o", fmt.Sprintf("pam-exec-oauth2.%s.%s", os, arch))
}
