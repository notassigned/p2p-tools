package cli

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
	p2p "github.com/notassigned/p2p-tools/libp2p"
	"github.com/sirupsen/logrus"
)

var Ping = cmd.Sub{
	Name:  "ping",
	Alias: "p",
	Short: "Ping a peer",
	Args:  &PingArgs{},
	Run:   PingRun,
}

type PingArgs struct {
	ID string `desc:"PeerID of peer to ping"`
}

func PingRun(r *cmd.Root, c *cmd.Sub) {
	if r.Flags.(*GlobalFlags).Log {
		logrus.SetLevel(logrus.DebugLevel)
	}
	args := c.Args.(*PingArgs)
	node := p2p.NewP2P()
	id, err := peer.Decode(args.ID)
	checkErr(err)
	pinger := ping.NewPingService(node.Host)
	ts := pinger.Ping(node.Ctx, id)
	var count int64 = 0
	rtts := make([]uint32, 0)
	startTime := time.Now()

	//capture ctrl+C and print statistics before program ends
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)
		<-ch
		endTime := time.Now()
		time := endTime.Sub(startTime)
		min, max, avg := minMaxAvg(rtts)

		fmt.Printf("\n--- ping statistics ---\n")
		fmt.Printf("peer: %s\n", args.ID)
		fmt.Printf("%d packets transmitted, time %dms\n", count, time.Milliseconds())
		fmt.Printf("rtt min/avg/max = %d/%d/%d ms\n", min, avg, max)

		os.Exit(0)
	}()

	for {
		select {
		case res := <-ts:
			checkErr(res.Error)
			time.Sleep(1 * time.Second)
			fmt.Printf("time=%dms seq=%d\n", res.RTT.Milliseconds(), count)
			rtts = append(rtts, uint32(res.RTT.Milliseconds()))
			count++
		case <-time.After(time.Second * 5):
			fmt.Println("Failed to receive ping after 5 seconds.")
			return
		}
	}

}

func minMaxAvg(array []uint32) (uint32, uint32, uint32) {
	var max uint32 = array[0]
	var min uint32 = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	var sum uint64 = 0
	for i := range array {
		sum += uint64(array[i])
	}
	avg := uint32(sum / uint64(len(array)))
	return min, max, avg
}
