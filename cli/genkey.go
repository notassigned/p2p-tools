package cli

import (
	"encoding/hex"
	"fmt"

	"github.com/DataDrake/cli-ng/v2/cmd"
	p2p "github.com/notassigned/p2p-tools/libp2p"
)

var GenKey = cmd.Sub{
	Name:  "genkey",
	Alias: "g",
	Short: "Generate a private key",
	Run:   GenKeyRun,
}

// GenKeyRun handles the execution of the genkey command.
func GenKeyRun(r *cmd.Root, c *cmd.Sub) {

	// Generate Private Key
	keyBytes, err := p2p.GeneratePrivKey()
	checkErr(err)
	fmt.Printf("%s\n", hex.EncodeToString(keyBytes))
}
