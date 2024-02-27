package app

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/costa92/krm/pkg/cli/cliflags"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var progressMessage = color.GreenString("==>")

type App struct {
	basename string
	name     string
	options  CliOptions
	cmd      *cobra.Command
	runFunc  RunFunc
}

func NewApp(name, basename string, opts ...Option) *App {
	app := &App{
		name:     name,
		basename: basename,
	}
	for _, o := range opts {
		o(app)
	}

	app.buildCommand()

	return app
}

func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		panic(err)
	}
}

// Option Path: pkg/server/server.go
// Compare this snippet from pkg/app/app.go:
// package app
type Option func(*App)

func WithOptions(opt CliOptions) Option {
	return func(a *App) {
		a.options = opt
	}
}

type RunFunc func(basename string) error

func WithRunFunc(run RunFunc) Option {
	return func(a *App) {
		a.runFunc = run
	}
}

func (a *App) buildCommand() {
	cmd := cobra.Command{
		Use:           formatBaseName(a.basename),
		Short:         a.name,
		SilenceUsage:  true,
		SilenceErrors: true,
		// Args:          a.args,
	}

	if a.runFunc != nil {
		cmd.RunE = a.runCommand
	}
	cliflags.InitFlags(cmd.Flags())
	var namedFlagSets cliflags.NamedFlagSets
	// Add flags from the options
	if a.options != nil {
		namedFlagSets = a.options.Flags()
		fs := cmd.Flags()
		for _, f := range namedFlagSets.FlagSets {
			fs.AddFlagSet(f)
		}
	}

	addConfigFlag(a.basename, namedFlagSets.FlagSet("global"))

	a.cmd = &cmd
}

func (a *App) runCommand(cmd *cobra.Command, _ []string) error {
	printWorkingDir()
	// Bind the flags to the options
	cliflags.PrintFlags(cmd.Flags())
	if err := viper.Unmarshal(a.options); err != nil {
		return err
	}
	// Apply the options
	if a.options != nil {
		if err := a.applyOptionRules(); err != nil {
			return err
		}
	}
	// Run the command
	if a.runFunc != nil {
		return a.runFunc(a.basename)
	}
	return nil
}

func formatBaseName(basename string) string {
	// Make case-insensitive and strip executable suffix if present
	if runtime.GOOS == "windows" {
		basename = strings.ToLower(basename)
		basename = strings.TrimSuffix(basename, ".exe")
	}
	return basename
}

// addConfigFlag adds the --config flag to the command :zh 添加 --config 标志到命令
func printWorkingDir() {
	wd, _ := os.Getwd()
	fmt.Printf("%v WorkingDir: %s \n", progressMessage, wd)
}

// applyOptionRules applies the rules to the options :zh 适用规则到选项
func (a *App) applyOptionRules() error {
	if completeOptions, ok := a.options.(CompleteableOptions); ok {
		if err := completeOptions.Complete(); err != nil {
			return err
		}
	}
	// Validate the options
	if errs := a.options.Validate(); len(errs) != 0 {
		panic(errs)
	}
	return nil
}
