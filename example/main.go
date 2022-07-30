package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	env "github.com/takekazuomi/pflag-env"
)

var (
	FlagOne string
)

func init() {
	flag.StringVarP(&FlagOne, "flag-one", "f", "default value", "flag test")
}

func main() {
	os.Setenv("TEST_FLAG_ONE", "overwrite by env")

	// Parse の前に、環境に変数にBindする
	env.BindEnv("TEST")

	// pflagのParseを呼ぶ
	flag.Parse()

	flag.Usage()

	fmt.Printf("FlagOne: %s\n", FlagOne)
}
