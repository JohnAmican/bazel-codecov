package codecov

import "testing"

func TestRun(t *testing.T) {
	const code = "foobar"
	if err := run("create-commit"); err != nil {
		t.Fatal(err)
	}
	if err := run("create-report", "--code", code); err != nil {
		t.Fatal(err)
	}
	if err := run("do-upload", "--report-code", code); err != nil {
		t.Fatal(err)
	}
	if err := run("create-report-results", "--code", code); err != nil {
		t.Fatal(err)
	}
	if err := run("get-report-results", "--code", code); err != nil {
		t.Fatal(err)
	}
}
