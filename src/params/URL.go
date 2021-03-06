package params

import (
	"regexp"

	"github.com/x1n13y84issmd42/oasis/src/contract"
	"github.com/x1n13y84issmd42/oasis/src/errors"
)

// URLParameters is the source for URL path parameters.
// URLParameters have an implicit requirement for the @HOSTNAME parameter
// which is an API host name.
type URLParameters struct {
	contract.EntityTrait
	*Set

	Path string
}

// URL creates a new URLParameters instance.
func URL(path string, log contract.Logger) *URLParameters {
	p := &URLParameters{
		EntityTrait: contract.Entity(log),
		Set:         NewSet("URL"),
		Path:        path,
	}

	p.Require(KeyHost)

	return p
}

// Make creates a URL string value from path template
// and parameters it has.
func (params URLParameters) String() string {
	if err := params.Validate(); err != nil {
		errors.Report(err, "URLParameters", params.Log)
	}

	tpl := "{" + KeyHost + "}" + params.Path

	for p := range params.Iterate() {
		rx := regexp.MustCompile("\\{" + p.N + "\\}")

		if rx.Match([]byte(tpl)) {
			v := p.V()
			if v != "" {
				params.Log.UsingParameterExample(p.N, "path", p.Source, v)
				tpl = string(rx.ReplaceAll([]byte(tpl), []byte(v)))
			}
		}
	}

	return tpl
}
