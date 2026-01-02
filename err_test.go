package brwsr_test

import (
	"github.com/speedyhoon/brwsr"
	"testing"
)

func TestErr(t *testing.T) {
	const errMsg = "unable to open web browser for URL"

	if errMsg != brwsr.Err.Error() {
		t.Errorf("error message is wrong,\nwant: `%s`,\n  got: `%s`", errMsg, brwsr.Err)
	}
}
