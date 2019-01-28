// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "golang-foo": status variables - See vendor/fabric8-services/fabric8-common/goasupport/status/generator.go
//
// Command:
// $ goagen
// --design=github.com/golang-starters/golang-rest-http/design
// --out=$(GOPATH)/src/github.com/golang-starters/golang-rest-http/app
// --version=v1.3.0

package app

import 	"time"

var (
	// Commit current build commit set by build script
	Commit = "0"
	// BuildTime set by build script in ISO 8601 (UTC) format: YYYY-MM-DDThh:mm:ssTZD (see https://www.w3.org/TR/NOTE-datetime for details)
	BuildTime = "0"
	// StartTime in ISO 8601 (UTC) format
	StartTime = time.Now().UTC().Format("2006-01-02T15:04:05Z")
)