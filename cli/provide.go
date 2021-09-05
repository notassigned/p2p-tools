package cli

import (
	"fmt"

	"github.com/DataDrake/cli-ng/v2/cmd"
	p2p "github.com/notassigned/p2p-tools/libp2p"
	"github.com/sirupsen/logrus"
)

var Provide = cmd.Sub{
	Name:  "provide",
	Alias: "pvd",
	Short: "Advertise content on the DHT",
	Args:  &ProvideArgs{},
	Run:   ProvideRun,
}

type ProvideArgs struct {
	Key string `desc:"String to advertise on the DHT"`
}

func ProvideRun(r *cmd.Root, c *cmd.Sub) {
	if r.Flags.(*GlobalFlags).Log {
		logrus.SetLevel(logrus.DebugLevel)
	}
	node := p2p.NewP2P()
	fmt.Println("Node ID:", node.Host.ID().Pretty())
	fmt.Println("Providing", c.Args.(*ProvideArgs).Key)
	node.Advertise(c.Args.(*ProvideArgs).Key)
	select {} //wait forever
}
