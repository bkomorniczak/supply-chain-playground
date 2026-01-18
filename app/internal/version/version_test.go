package version

import (
	"testing"
)

func TestInfoHasFields(t *testing.T) {
	v := Info()
	if v.Commit == "" {
		t.Error("expected Commit to be set")
	}
	if v.BuildTime == "" {
		t.Error("expected BuildTime to be set")
	}
	if v.GoVersion == "" {
		t.Error("expected GoVersion to be set")
	}
}