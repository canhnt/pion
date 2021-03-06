package main

import (
	"flag"
	"os"

	"github.com/kpn/pion/pkg/app/proxy"
	"github.com/spf13/pflag"
)

func init() {
	flags := pflag.NewFlagSet("", pflag.ExitOnError)
	flag.Set("logtostderr", "true")
	flags.AddGoFlagSet(flag.CommandLine)
	flags.Parse(os.Args)
}

func main() {
	proxy.NewApp().Start()
}
