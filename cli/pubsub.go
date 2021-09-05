package cli

import (
	"encoding/json"
	"fmt"

	"github.com/DataDrake/cli-ng/v2/cmd"
	p2p "github.com/notassigned/p2p-tools/libp2p"
	"github.com/sirupsen/logrus"
)

var Pubsub = cmd.Sub{
	Name:  "pubsub",
	Alias: "ps",
	Short: "Watch a pubsub topic",
	Args:  &PubsubArgs{},
	Run:   PubsubRun,
	Flags: &PubsubFlags{},
}

type PubsubFlags struct {
	CID bool `long:"cid" desc:"Specifies that the key is already a CID"`
}

type PubsubArgs struct {
	Topic string `desc:"Topic to listen to"`
}

func PubsubRun(r *cmd.Root, c *cmd.Sub) {
	if r.Flags.(*GlobalFlags).Log {
		logrus.SetLevel(logrus.DebugLevel)
	}
	//flags := c.Flags.(*PubsubFlags)
	args := c.Args.(*PubsubArgs)

	node := p2p.NewP2P()
	topic, err := node.PubSub.Join(args.Topic)
	checkErr(err)
	subscription, err := topic.Subscribe()
	checkErr(err)
	for {
		msg, err := subscription.Next(node.Ctx)
		checkErr(err)
		msgJson, err := json.MarshalIndent(msg, "", "  ")
		fmt.Println(string(msgJson[:]))
	}
}
