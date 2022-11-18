package example_iwanta

import (
	"testing"

	gutowire "github.com/windrivder/go-autowire"
	"github.com/windrivder/go-autowire/example/dependencies"
	"github.com/windrivder/go-autowire/example_zoo"
)

var zoo example_zoo.Dog
var testI dependencies.Test
var local Local

// run go test ./example_iwanta/...
// then it will generate the init file
var _ = gutowire.IWantA(&local)
var _ = gutowire.IWantA(&zoo)
var _ = gutowire.IWantA(&testI)

func TestRun(t *testing.T) {}
