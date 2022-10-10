package main

import (
	"regexp"
	"strings"
	"testing"
)

func TestRegular(t *testing.T) {
	s := "/go.dev/doc/install"
	r := regexp.MustCompile("/(.+?)/")
	host := strings.Trim(r.FindString(s), "/")
	if host != "go.dev" {
		t.Errorf("parse failed")
	}
}
