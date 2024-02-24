package cliflags

import (
	goflag "flag"
	"github.com/costa92/logger"
	"strings"

	"github.com/spf13/pflag"
)

func WordSepNormalizeFunc(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}
	return pflag.NormalizedName(name)
}

func InitFlags(flags *pflag.FlagSet) {
	flags.SetNormalizeFunc(WordSepNormalizeFunc)
	flags.AddGoFlagSet(goflag.CommandLine)
}

func WarnWordSepNormalizeFunc(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		nname := strings.ReplaceAll(name, "_", "-")
		return pflag.NormalizedName(nname)
	}
	return pflag.NormalizedName(name)
}

func PrintFlags(flags *pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		//fmt.Printf(" FLAG: --%s=%q", flag.Name, flag.Value)
		logger.Infof(" FLAG: --%s=%q", flag.Name, flag.Value)
	})
}
