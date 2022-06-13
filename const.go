package gutowire

import (
	"go/ast"
	"text/template"
)

const (
	filePrefix  = "autowire"
	wireTag     = "@autowire"
	setTemplate = `// Code generated by go-autowire. DO NOT EDIT.

package {{ .Package }}

import (
	"github.com/google/wire"
)

var {{ .SetName }} = wire.NewSet({{ range $Item := .Items}} 
	{{ $Item }},
    {{ end }}
)
`
	initTemplateHead = `// Code generated by go-autowire. DO NOT EDIT.

// +build wireinject

package %s
`
	initItemTemplate = `
func Initialize%s(%s) (%s, func(), error) {
	panic(wire.Build(Sets))
}
`
)

var setTemp = template.Must(template.New("").Parse(setTemplate))

type (
	wireSet struct {
		Package string
		Items   []string
		SetName string
	}

	opt struct {
		searchPath string
		pkg        string
		genPath    string
		initWire   []string
	}

	Option func(*opt)

	searcher struct {
		sets           []string
		genPath        string
		pkg            string
		elementMap     map[string]map[string]element
		options        []Option
		modBase        string
		initElements   []element
		configElements []element
		initWire       []string
	}

	element struct {
		name        string
		constructor string
		fields      []string
		implements  []string
		pkg         string
		pkgPath     string
		typ         uint
		initWire    bool
		configWire  bool
	}

	tmpDecl struct {
		docs     string
		name     string
		isFunc   bool
		typeSpec *ast.TypeSpec
	}
)

func WithPkg(pkg string) Option {
	return func(o *opt) {
		o.pkg = pkg
	}
}

func InitWire(initStruct ...string) Option {
	return func(o *opt) {
		if len(initStruct) == 0 {
			initStruct = []string{"*"}
		}
		o.initWire = initStruct
	}
}

func WithSearchPath(path string) Option {
	return func(o *opt) {
		o.searchPath = path
	}
}
