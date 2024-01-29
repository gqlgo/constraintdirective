package constraintdirective_test

import (
	"testing"

	"github.com/gqlgo/constraintdirective"
	"github.com/gqlgo/gqlanalysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, constraintdirective.Analyzer([]string{"first"}), "a")
}
