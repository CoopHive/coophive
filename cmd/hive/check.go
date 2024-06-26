package hive

import (
	"encoding/json"
	"fmt"

	"github.com/CoopHive/coophive/pkg/inspect"
	optionsfactory "github.com/CoopHive/coophive/pkg/options"
	"github.com/CoopHive/coophive/pkg/web3"
	"github.com/CoopHive/coophive/pkg/web3/bindings/storage"
	"github.com/spf13/cobra"
)

func newCheckCmd() *cobra.Command {
	options := optionsfactory.NewCheckOptions()
	runCmd := &cobra.Command{
		Use:     "check",
		Short:   "confirm your installation is running a service correctly",
		Long:    "confirm your installation is running a service correctly",
		Example: "check resource-provider-gpu",
		RunE: func(cmd *cobra.Command, args []string) error {

			options, err := optionsfactory.ProcessCheckOptions(options, args)
			if err != nil {
				return err
			}
			return check(cmd, options)
		},
	}
	optionsfactory.AddCheckCliFlags(runCmd, &options)
	return runCmd
}

func check(cmd *cobra.Command, options check.checkOptions) error {
	// try to run docker run --rm --runtime=nvidia --gpus all ubuntu nvidia-smi

	out, err := inspect.RunDockerCommand("--rm", "--runtime=nvidia", "--gpus", "all", "ubuntu", "nvidia-smi")
	if err != nil {
		return fmt.Errorf("failed to run docker command: %w", err)
	}

	fmt.Println(out)

	return nil
}
