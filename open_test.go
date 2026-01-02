package brwsr_test

import (
	"fmt"
	"github.com/speedyhoon/brwsr"
	"testing"
)

func TestOpen(t *testing.T) {
	tests := []struct {
		url     string
		wantErr bool
	}{
		{url: "http://localhost/"},
		{url: "127.0.0.1"},
		{url: "invalidProtocol://whoops"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test[%d]", i), func(t *testing.T) {
			err := brwsr.Open(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
