package router_test

import (
	"testing"

	. "github.com/rhuancaetano/enceladus/pkg/router"
)

type PathTest struct {
	path   string
	output bool
}

var pathIsValidTests = []PathTest{
	{"", false},
	{"test", false},
	{"/", true},
	{"/test", true},
	{"/test/", true},
	{"test/", false},
	{"/:test", true},
	{"/:test/test", true},
	{"/:test/test/:test", false},
	{"/test/test/test", true},
	{"/:red/test/:red", false},
	{"/:red/test/:blue", true},
	{"/red/test/blue", true},
	{"/?", false},
	{"/test?", false},
	{"/test@", false},
}

func TestPathIsValid(t *testing.T) {
	for _, test := range pathIsValidTests {
		path := Path(test.path)
		if isValid, _ := path.IsValid(); isValid != test.output {
			t.Errorf("path.isValid '%s' expected %t but got %t", test.path, test.output, isValid)
		}
	}
}

var pathIsParamTests = []PathTest{
	{"", false},
	{"test", false},
	{"/", false},
	{"/test", false},
	{"/test/", false},
	{"test/", false},
	{"/:test", true},
	{"/:test/test", true},
	{"/:test/test/:test", true},
	{"/test/:test", false},
	{"/?", false},
	{"/test?", false},
	{"/test@", false},
}

func TestPathIsParam(t *testing.T) {
	for _, test := range pathIsParamTests {
		path := Path(test.path)
		if IsParam := path.IsParam(); IsParam != test.output {
			t.Errorf("path.IsParam '%s' expected %t but got %t", test.path, test.output, IsParam)
		}
	}
}
