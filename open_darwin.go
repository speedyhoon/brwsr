package brwsr

import "os/exec"

func Open(url string) (err error) {
	err = exec.Command("open", url).Start()
	if err != nil {
		return Err
	}

	return
}
