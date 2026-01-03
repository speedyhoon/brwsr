// Package brwsr implements a way to open URL links in the user's default browser, compatible with most operating systems.
package brwsr

import "errors"

// ErrOpen is unable to open the default web browser for a URL.
var ErrOpen = errors.New("open failed")
