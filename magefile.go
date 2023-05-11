//go:build mage
// +build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/sh"
)

const ARCHARM64 = "arm64"
const ARCHARM = "arm"
const ARCHAMD64 = "amd64"

const OSLINUX = "linux"

func Linux() error {
	return build(ARCHAMD64, OSLINUX)
}

func Arm64() error {
	return build(ARCHARM64, OSLINUX)
}

func Arm() error {
	return build(ARCHARM, OSLINUX)
}

func build(arch, os string) error {
	return sh.RunV("go", "build", "-o", fmt.Sprintf("pam-exec-oauth2.%s.%s", os, arch))
}
