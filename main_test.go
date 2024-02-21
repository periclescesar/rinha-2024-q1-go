package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/periclescesar/rinha-2024-q1-go/configs"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Paths:  []string{"features"},
	Format: "pretty", // can define default values
}

func init() {
	configs.InitConfigs(".test.env")

	godog.BindCommandLineFlags("godog.", &opts) // godog v0.11.0 (latest)
}

func TestMain(m *testing.M) {
	pflag.Parse()
	opts.Paths = pflag.Args()

	status := godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	// Optional: Run `testing` package's logic besides godog.
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}

//func TestFeatures(t *testing.T) {
//	suite := godog.TestSuite{
//		ScenarioInitializer: InitializeScenario,
//		Options: &godog.Options{
//			Format:   "pretty",
//			Paths:    []string{"features"},
//			TestingT: t, // Testing instance that will run subtests.
//		},
//	}
//
//	if suite.Run() != 0 {
//		t.Fatal("non-zero status returned, failed to run feature tests")
//	}
//}

// assertExpectedAndActual is a helper function to allow the step function to call
// assertion functions where you want to compare an expected and an actual value.
func assertExpectedAndActual(a expectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	a(&t, expected, actual, msgAndArgs...)
	return t.err
}

type expectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

// asserter is used to be able to retrieve the error reported by the called assertion
type asserter struct {
	err error
}

// Errorf is used by the called assertion to report an error
func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}
