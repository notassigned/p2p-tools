package cli

import (
	"encoding/json"
	"fmt"

	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/libp2p/go-libp2p-core/peer"
	p2p "github.com/notassigned/p2p-tools/libp2p"
	"github.com/sirupsen/logrus"
)

var Dht = cmd.Sub{
	Name:  "dhtlookup",
	Alias: "dht",
	Short: "Lookup records on the DHT",
	Args:  &DhtArgs{},
	Run:   DhtRun,
	Flags: &DhtFlags{},
}

type DhtFlags struct {
	CID bool `long:"cid" desc:"Specifies that the key is already a CID"`
}

type DhtArgs struct {
	RecordType string `desc:"Either 'peer' or 'content' records"`
	Key        string `desc:"PeerID or Key to search for"`
}

func DhtRun(r *cmd.Root, c *cmd.Sub) {
	if r.Flags.(*GlobalFlags).Log {
		logrus.SetLevel(logrus.DebugLevel)
	}
	flags := c.Flags.(*DhtFlags)
	args := c.Args.(*DhtArgs)
	if args.RecordType != "peer" && args.RecordType != "content" {
		fmt.Println("RecordType must be either 'peer' or 'content'")
		return
	}

	node := p2p.NewP2P()

	switch args.RecordType {
	case "peer":
		id, err := peer.Decode(args.Key)
		checkErr(err)
		peer, err := node.KadDHT.FindPeer(node.Ctx, id)
		checkErr(err)
		peerJson, err := json.MarshalIndent(peer, "", "  ")
		checkErr(err)
		fmt.Println(string(peerJson))
		break
	case "content":
		logrus.Debugln("Searching for peers providing", args.Key)
		total := p2p.FindPeers(node, args.Key, flags.CID)
		logrus.Debugf("Found %d peers.\n", total)
		break
	}
}
