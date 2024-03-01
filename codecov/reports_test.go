package codecov

import "testing"

func TestDats(t *testing.T) {
	for _, cov := range coverageReports() {
		println(cov)
	}
}
