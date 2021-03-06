package main

import (
	"os"

	"github.com/x1n13y84issmd42/oasis/src/contract"
	"github.com/x1n13y84issmd42/oasis/src/env"
	"github.com/x1n13y84issmd42/oasis/src/test"
	"github.com/x1n13y84issmd42/oasis/src/utility"
)

// Manual is an entry point for manual testing mode.
func Manual(args *env.Args, logger contract.Logger) {
	spec := utility.Load(args.Spec, logger)

	logger.TestingProject(spec)

	// Resolving the operations.
	specOps := utility.NewOperationResolver(spec, logger).Resolve(args.Ops)
	result := test.Success()

	if len(specOps) > 0 {
		for _, op := range specOps {
			logger.TestingOperation(op)

			// Stuffing it with data.
			op.Data().URL.Load(args.Use.PathParameters)
			op.Data().URL.Load(op.Resolve().Host(args.Host))
			op.Data().Query.Load(args.Use.Query)
			op.Data().Headers.Load(args.Use.Headers)
			op.Data().Body.Load(args.Use.Body)

			enrichment := []contract.RequestEnrichment{
				op.Data().Query,
				op.Data().Headers,
				op.Data().Body,

				op.Resolve().Security(args.Use.Security),
			}

			v := op.Resolve().Response(args.Expect.Status, args.Expect.CT)

			// Testing.
			result = result.And(test.Operation(op, &enrichment, v, logger))
		}

	} else {
		logger.PrintOperations(spec.Operations())
	}

	if !result.Success {
		os.Exit(255)
	}
}
