package test

import (
	"encoding/json"
	"fmt"
	"github.com/JFJun/go-substrate-crypto/ss58"
	"github.com/rjman-ljm/substrate-go/client"
	"github.com/rjman-ljm/substrate-go/expand"
	"testing"
)

//const url = "wss://chainx.elara.patract.io"
//const url = "ws://127.0.0.1:9944"
//const url = "wss://xbridge.spiderx.pro/ws"
//const url = "wss://polkadot.elara.patract.io"
//const url = "wss://dot.supercube.pro/ws"
//const url = "wss://chainx.supercube.pro/ws"
//const url = "ws://localhost:9977"
const url = "wss://supercube.pro/ws"

func Test_GetBlockByNumber(t *testing.T) {
	c, err := client.New(url)
	if err != nil {
		t.Fatal(err)
	}
	c.SetPrefix(ss58.ChainXPrefix)
	c. Name = expand.ClientNameChainX

	//expand.SetSerDeOptions(false)
	resp, err := c.GetBlockByNumber(11255)
	if err != nil {
		t.Fatal(err)
	}

	hash, err := c.Api.RPC.Chain.GetBlockHash(11255)
	if err != nil {
		fmt.Printf("GetBlockHash err\n")
	}

	block, err := c.Api.RPC.Chain.GetBlock(hash)
	if err != nil {
		fmt.Printf("GetBlock err\n")
		//api, _ := gsrpc.NewSubstrateAPI(url)
		//
		//hash, err := api.RPC.Chain.GetBlockHash(4744169)
		//if err != nil {
		//	fmt.Printf("GetBlockHash err\n")
		//}
		//
		//block, err := api.RPC.Chain.GetBlock(hash)
		//if err != nil {
		//	fmt.Printf("Get Block err\n")
		//}
		//
		//if block != nil {
		//	currentBlock := int64(block.Block.Header.Number)
		//	fmt.Printf("block is %v\n", currentBlock)
		//}
	}

	if block != nil {
		currentBlock := int64(block.Block.Header.Number)
		fmt.Printf("block is %v\n", currentBlock)
	}

	d, _ := json.Marshal(resp)
	fmt.Println(string(d))
}

func Test_GetAccountInfo(t *testing.T) {
	c, err := client.New(url)
	if err != nil {
		t.Fatal(err)
	}
	c.SetPrefix(ss58.PolkadotPrefix)
	ai, err := c.GetAccountInfo("15oF4uVJwmo4TdGW7VfQxNLavjCXviqxT9S1MgbjMNHr6Sp5")
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(ai)
	fmt.Println(string(d))
	fmt.Println(ai.Data.Free.String())
}
