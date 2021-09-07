# p2p-tools for [libp2p](https://github.com/libp2p/libp2p)

## Overview
Command line tools to perform DHT lookups, ping peers, and more.

## Usage and Examples

Install Go 1.17 or later

### Clone and build

```
git clone https://github.com/notassigned/p2p-tools.git
cd p2p-tools
go build
./p2p-tools
```
## Ping a peerID using the libp2p ping protocol
```
$ ./p2p-tools ping QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN
time=118ms seq=0
time=100ms seq=1
time=100ms seq=2
time=99ms seq=3
time=99ms seq=4
^C
--- ping statistics ---
peer: QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN
5 packets transmitted, time 5712ms
rtt min/avg/max = 99/103/118 ms
```

## Lookup peer records on the Kademlia DHT
```
$ ./p2p-tools dht peer QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN
{
  "ID": "QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
  "Addrs": [
    "/ip4/127.0.0.1/tcp/4001",
    "/ip6/2604:1380:1000:6000::1/tcp/4001",
    "/ip6/2604:1380:1000:6000::1/udp/4001/quic",
    "/ip6/::1/udp/4001/quic",
    "/ip6/::1/tcp/4001",
    "/ip4/127.0.0.1/udp/4001/quic",
    "/ip4/127.0.0.1/tcp/8081/ws",
    "/dnsaddr/bootstrap.libp2p.io",
    "/ip4/147.75.109.213/tcp/4001",
    "/ip4/147.75.109.213/udp/4001/quic"
  ]
}
```

## Lookup content records on the DHT
In this case "test" will be converted to a CID. Quotations are optional. To search for an already-computed CID use the flag --cid to specify that the key is already a CID
```
$ ./p2p-tools dht content "test"
{
  "ID": "12D3KooWNyeZ6jSaNdHU2BRPFDADg2LWyuxHrASpvRq1T5hGqga2",
  "Addrs": [
    "/ip4/139.178.69.51/tcp/4001",
    "/ip4/127.0.0.1/tcp/4001"
  ]
},
{
  "ID": "12D3KooWEDMw7oRqQkdCJbyeqS5mUmWGwTp8JJ2tjCzTkHboF6wK",
  "Addrs": [
    "/ip6/2604:1380:45e1:2700::3/tcp/4002/ws",
    "/ip4/139.178.68.91/tcp/4001",
    "/ip6/2604:1380:45e1:2700::3/tcp/4001",
    "/ip4/139.178.68.91/tcp/4002/ws"
  ]
},
{
  "ID": "12D3KooWK42UnJJq3BMgVBEP7Bdmqcy45vrehSxcpQ8GVsj6dxkM",
  "Addrs": []
},
```
## Provide content on the DHT
Node will advertize the string on the DHT until the program is stopped.
```
$ ./p2p-tools provide "test"
Node ID: QmPQdpPUCg1C96BvQjERJcGh5FSWUFhJodqojZsvHTquoJ
Providing test

```
