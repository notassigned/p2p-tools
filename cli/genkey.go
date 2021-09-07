package cli

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/mr-tron/base58/base58"
	"github.com/multiformats/go-multihash"
	p2p "github.com/notassigned/p2p-tools/libp2p"
)

var GenKey = cmd.Sub{
	Name:  "genkey",
	Alias: "g",
	Short: "Generate a key pair",
	Run:   GenKeyRun,
}

type Identity struct {
	//peerId
	Id string

	//serialized public key in protobuf format
	PubKey []byte

	//serialized private key in protobuf format
	PrivKey []byte
}

// GenKeyRun handles the execution of the genkey command.
func GenKeyRun(r *cmd.Root, c *cmd.Sub) {

	// Generate key pair and convert to serialized versions
	privKey, pubKey, err := p2p.GeneratePrivKey()
	checkErr(err)
	pubKeyMarshalled, err := crypto.MarshalPublicKey(pubKey)
	checkErr(err)
	privKeyMarshalled, err := crypto.MarshalPrivateKey(privKey)
	checkErr(err)
	peerId := getPeerID(pubKey)

	id := &Identity{
		Id:      peerId,
		PubKey:  pubKeyMarshalled,
		PrivKey: privKeyMarshalled,
	}

	idJSON, err := json.MarshalIndent(id, "", "  ")
	checkErr(err)

	fmt.Println(string(idJSON[:]))
}

//cant find the function for this so we'll just make it
//returns peer id encoded as a raw base58btc multihash
func getPeerID(pubKey crypto.PubKey) (result string) {
	pubKeyBytes, _ := pubKey.Bytes()
	hash := sha256.New()
	hash.Write(pubKeyBytes)
	multi, _ := multihash.Encode(hash.Sum(nil), multihash.SHA2_256)
	result = base58.Encode(multi)
	return
}
