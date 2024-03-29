package hive

import (
	"github.com/CoopHive/coophive/pkg/jobcreator"
	optionsfactory "github.com/CoopHive/coophive/pkg/options"
	"github.com/CoopHive/coophive/pkg/system"
	"github.com/CoopHive/coophive/pkg/web3"
	"github.com/spf13/cobra"
)

func newJobCreatorCmd() *cobra.Command {
	options := optionsfactory.NewJobCreatorOptions()

	solverCmd := &cobra.Command{
		Use:     "jobcreator",
		Short:   "Start the hive job creator service.",
		Long:    "Start the hive job creator service.",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			options, err := optionsfactory.ProcessOnChainJobCreatorOptions(options, args)
			if err != nil {
				return err
			}
			return runJobCreator(cmd, options)
		},
	}

	optionsfactory.AddJobCreatorCliFlags(solverCmd, &options)

	return solverCmd
}

func runJobCreator(cmd *cobra.Command, options jobcreator.JobCreatorOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	// create the job creator and start it's control loop
	jobCreatorService, err := jobcreator.NewOnChainJobCreator(options, web3SDK)
	if err != nil {
		return err
	}

	jobCreatorErrors := jobCreatorService.Start(commandCtx.Ctx, commandCtx.Cm)

	for {
		select {
		case err := <-jobCreatorErrors:
			return err
		case <-commandCtx.Ctx.Done():
			return nil
		}
	}
}
