// Package brwsr implements a way to open URL links in the user's default browser, compatible with most operating systems.
package brwsr

import "errors"

var Err = errors.New("unable to open web browser for URL")
