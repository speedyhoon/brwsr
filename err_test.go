package brwsr_test

import (
	"github.com/speedyhoon/brwsr"
	"testing"
)

func TestErr(t *testing.T) {
	const expectErr = "open failed"

	if expectErr != brwsr.ErrOpen.Error() {
		t.Errorf("error message is wrong,\nwant: `%s`,\n  got: `%s`", expectErr, brwsr.ErrOpen)
	}
}
