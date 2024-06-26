package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const configFlagName = "config"

var cfgFile string

// nolint: gochecknoinits
func init() {
	pflag.StringVarP(&cfgFile, "config", "c", cfgFile, "Read configuration from specified `FILE`, "+
		"support JSON, TOML, YAML, HCL, or Java properties formats.")
}

func addConfigFlag(basename string, fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(configFlagName))
	viper.AutomaticEnv()
	viper.SetEnvPrefix(strings.ReplaceAll(strings.ToUpper(basename), "-", "_"))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	cobra.OnInitialize(func() {
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		} else {
			viper.SetConfigName(basename)
			viper.AddConfigPath(".")
			viper.AddConfigPath("$HOME/.config")
			viper.AddConfigPath("/etc")
			viper.AddConfigPath("configs")
			viper.SetConfigType("yaml")
		}
		if err := viper.ReadInConfig(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: failed to read configuration file(%s): %v\n", cfgFile, err)
			os.Exit(1)
		}
	})
}
