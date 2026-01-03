package brwsr

import "os/exec"

// Open a URL link in the user's default browser.
func Open(url string) (err error) {
	err = exec.Command("open", url).Start()
	if err != nil {
		return ErrOpen
	}

	return
}
