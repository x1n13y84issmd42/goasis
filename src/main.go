package main

import (
	"fmt"
	"os"

	"github.com/x1n13y84issmd42/oasis/src/api"
	"github.com/x1n13y84issmd42/oasis/src/log"
	"github.com/x1n13y84issmd42/oasis/src/test"
	"github.com/x1n13y84issmd42/oasis/src/utility"
)

func main() {
	//	Command line args
	var inSpec string
	var inScript string
	var inHost string
	// var inOp string
	var inOps []string
	var useCT string = "*"
	var useSecurity string
	var expCT string = "*"
	var expStatus int64 = 0

	expExecute := utility.CLIQL().Flag("execute").Capture(&inScript)
	expFrom := utility.CLIQL().Flag("from").Capture(&inSpec)
	expTest := utility.CLIQL().Flag("test").CaptureStringSlice(&inOps)
	// expTest := utility.CLIQL().Flag("test").Capture(&inOp)
	expHost := utility.CLIQL().Flag("@").Capture(&inHost)

	expUse := utility.CLIQL().Flag("use").Repeat(utility.CLIQL().Any([]*utility.CLIQLParser{
		utility.CLIQL().Flag("security").Capture(&useSecurity),
	}), 0, 1)

	expExpect := utility.CLIQL().Flag("expect").Repeat(utility.CLIQL().Any([]*utility.CLIQLParser{
		utility.CLIQL().Flag("CT").Capture(&expCT),
		utility.CLIQL().Flag("status").CaptureInt64(&expStatus),
	}), 0, 2)

	utility.CLIQL().Repeat(utility.CLIQL().Any([]*utility.CLIQLParser{
		expExecute,
		expFrom,
		expTest,
		expUse,
		expExpect,
		expHost,
	}), 1, 4).Parse(os.Args[1:])

	if inScript != "" {
		//	Executing a test script
		//	TODO
	} else if inSpec != "" {
		//	Running a single test
		spec, specErr := api.Load(inSpec)

		if specErr == nil {
			runner := test.Runner{
				Spec: spec,
				Log:  log.Simple{},
			}

			runner.Log.TestingProject(spec.GetProjectInfo())

			if hostOK := runner.UseHost(inHost); hostOK {
				testResult := true
				for _, inOp := range inOps {
					testResult = runner.Test(inOp, useCT, int(expStatus), expCT) && testResult
				}
				if !testResult {
					os.Exit(255)
				}
			} else {
				os.Exit(255)
			}
		} else {
			fmt.Println(specErr)
			os.Exit(255)
		}
	} else {
		fmt.Println("Please specify at least a spec file & an operation to test.")
		fmt.Println("Example:")
		fmt.Println("oasis from path/to/oas_spec.yaml test operation_id")
	}
}
