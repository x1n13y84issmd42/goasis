package utility

import (
	"github.com/x1n13y84issmd42/oasis/src/api"
	"github.com/x1n13y84issmd42/oasis/src/api/openapi3"
	"github.com/x1n13y84issmd42/oasis/src/contract"
)

//go:generate pwd

// Load loads an API spec file.
func Load(path string, logger contract.Logger) contract.Spec {
	logger.LoadingSpec(path)

	//TODO: later we'll need a way to select loaders based on format.
	spec, specErr := openapi3.Load(path, logger)
	if specErr != nil {
		return api.NoSpec(specErr, logger)
	}

	return spec
}
