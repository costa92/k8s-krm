package cli

import (
	"fmt"

	"github.com/costa92/krm/pkg/app"
	"github.com/costa92/krm/pkg/cli/cliflags"
	"github.com/costa92/krm/pkg/cli/globalflag"
	"github.com/costa92/krm/pkg/util/term"
	"github.com/spf13/cobra"
)

type RunFunc func(basename string) error

func BuildCommand(basename string, runFunc RunFunc) *cobra.Command {
	cmd := cobra.Command{
		Use:   basename,
		Short: fmt.Sprintf("%s is a Kubernetes resource manager", basename),
		Long:  fmt.Sprintf("%s is a Kubernetes resource manager", basename),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runFunc(basename)
		},
	}
	var namedFlagSets cliflags.NamedFlagSets
	cliflags.InitFlags(cmd.Flags())
	app.AddConfigFlag(basename, namedFlagSets.FlagSet("global"))
	globalflag.AddGlobalFlags(namedFlagSets.FlagSet("global"), cmd.Name())
	// cmd.SetHelpCommand(app.HelpCommand(FormatBaseName(basename)))
	cmd.Flags().AddFlagSet(namedFlagSets.FlagSet("global"))
	// 设置分组
	// addCmdTemplate(&cmd, namedFlagSets)
	return &cmd
}

func chCommand() *cobra.Command {
	return nil
}

func addCmdTemplate(cmd *cobra.Command, namedFlagSets cliflags.NamedFlagSets) {
	usageFmt := "Usage:\n  %s\n"
	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
		cliflags.PrintSections(cmd.OutOrStderr(), namedFlagSets, cols)
		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
		cliflags.PrintSections(cmd.OutOrStdout(), namedFlagSets, cols)
	})
}
