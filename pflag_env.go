// Package pflagenv は、pflagに環境変数での上書きを追加するヘルパー
package pflagenv

import (
	"os"
	"strings"

	flag "github.com/spf13/pflag"
)

func toEnvName(prefix string, s string) string {
	s = strings.ToUpper(s)
	s = strings.Replace(s, "-", "_", -1)
	if len(prefix) != 0 {
		prefix = prefix + "_"
	}
	return prefix + s
}

// BindEnv は、pflag.CommandLine の値を環境変数で上書きする
func BindEnv(envPrefix string) {
	flag.CommandLine.VisitAll(func(f *flag.Flag) {
		e := toEnvName(envPrefix, f.Name)
		f.Usage = f.Usage + ", [" + e + "]"
		if n, ok := os.LookupEnv(e); ok {
			// avoid error. TODO panic or not?
			f.Value.Set(n)
		}
	})
}
