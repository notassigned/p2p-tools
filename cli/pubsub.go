package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

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
	MultiLine  bool `long:"multi" short:"m" desc:"Enable multiline input. Press Ctrl+] on its own line then Enter to send"`
	OutputText bool `long:"text" short:"t" desc:"Output message text rather than entire message struct"`
}

type PubsubArgs struct {
	Mode  string `desc:"Either listen or publish"`
	Topic string `desc:"Topic to listen to"`
}

func PubsubRun(r *cmd.Root, c *cmd.Sub) {
	if r.Flags.(*GlobalFlags).Log {
		logrus.SetLevel(logrus.DebugLevel)
	}
	flags := c.Flags.(*PubsubFlags)
	args := c.Args.(*PubsubArgs)

	//create node and subscribe to the topic
	node := p2p.NewP2P()
	topic, err := node.PubSub.Join(args.Topic)
	checkErr(err)
	subscription, err := topic.Subscribe()
	checkErr(err)

	scn := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(os.Stdin)

	switch args.Mode {
	case "listen":
		//read messages from the topic and output them to the console
		for {
			msg, err := subscription.Next(node.Ctx)
			checkErr(err)
			if flags.OutputText {
				fmt.Println("From: " + msg.ReceivedFrom.Pretty())
				fmt.Println(string(msg.Data[:]))
			} else {
				msgJson, _ := json.MarshalIndent(msg, "", "  ")
				fmt.Println(string(msgJson[:]) + ",")
			}
		}
	case "publish":
		//read messages from the console and publish them in the topic
		for {
			if flags.MultiLine {
				input := getLines(scn)
				topic.Publish(node.Ctx, input)
			} else {
				input, err := reader.ReadString('\n')
				checkErr(err)
				topic.Publish(node.Ctx, []byte(input))
			}
		}
	}
}

//enables getting multiline input e.g. paste
//user presses Ctrl+] on its own line then Enter to return
func getLines(scn *bufio.Scanner) (input []byte) {
	var lines []string
	for scn.Scan() {
		line := scn.Text()
		if len(line) == 1 {
			// Group Separator (GS ^]): ctrl-]
			if line[0] == '\x1D' {
				break
			}
		}
		lines = append(lines, line)
	}

	if len(lines) > 0 {
		for _, line := range lines {
			bytes := []byte(line)
			input = append(input[:], bytes[:]...)
		}
	}

	err := scn.Err()
	checkErr(err)

	if len(lines) == 0 {
		fmt.Println("Exited.")
		os.Exit(0)
	}

	return
}
