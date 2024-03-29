// Code to build a simple host with libp2p, this will start the peer to peer network
package main

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
)

func main() {
	// The context governs the lifetime of the libp2p node
	//ctx := context.Background()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// To construct a simple host with all the default settings, we just use 'new'
	node, err := libp2p.New(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello World, my hosts ID is %s\n", node.ID())
	fmt.Println("Listen addresses: ", node.Addrs())

	// Shut the node down
	if err := node.Close(); err != nil {
		panic(err)
	}

	// If you want more control over the configuration, you can specify some options

	// set your own keypair
	priv, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		panic(err)
	}

	node2, err := libp2p.New(ctx,
		//Use your own created keypair
		libp2p.Identity(priv),
		//set your own listen address
		//the config takes an arra of addresses, specify as many as you want
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/9000"),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello World, my hosts ID is %s\n", node2.ID())
	fmt.Println("Listen addresses: ", node2.Addrs())

}
