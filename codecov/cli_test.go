package codecov

import "testing"

func TestRun(t *testing.T) {
	if err := run("--help"); err != nil {
		t.Fatal(err)
	}
}
