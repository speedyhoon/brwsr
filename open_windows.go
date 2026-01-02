package brwsr

import "os/exec"

func Open(url string) (err error) {
	err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	if err != nil {
		return Err
	}

	return
}
