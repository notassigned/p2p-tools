package cli

import (
	"log"

	"github.com/DataDrake/cli-ng/v2/cmd"
)

var appVersion string = "0.0.1"

type GlobalFlags struct {
	Log bool `short:"l" long:"log" desc:"Output debug log."`
}

var Root *cmd.Root

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	Root = &cmd.Root{
		Name:    "p2p-tools",
		Short:   "Tools for libp2p",
		Version: appVersion,
		Flags:   &GlobalFlags{},
	}

	cmd.Register(&cmd.Help)
	cmd.Register(&GenKey)
	cmd.Register(&Dht)
	cmd.Register(&Ping)
	cmd.Register(&Provide)
	cmd.Register(&Pubsub)
}
