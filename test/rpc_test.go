package test

import (
	"encoding/json"
	"fmt"
	"github.com/JFJun/go-substrate-crypto/ss58"
	"github.com/rjman-self/substrate-go/client"
	"testing"
)

const url = "wss://supercube.pro/ws"
//const url = "wss://chainx.elara.patract.io"
//const url = "ws://127.0.0.1:8087"
//const url = "wss://xbridge.spiderx.pro/ws"
//const url = "wss://polkadot.elara.patract.io"

const testBlock = 215758

func Test_GetBlockByNumber(t *testing.T) {
	c, err := client.New(url)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPrefix(ss58.PolkadotPrefix)
	//expand.SetSerDeOptions(false)
	resp, err := c.GetBlockByNumber(testBlock)
	if err != nil {
		t.Fatal(err)
	}

	hash, err := c.Api.RPC.Chain.GetBlockHash(testBlock)
	if err != nil {
		fmt.Printf("GetBlockHash err\n")
	}

	block, err := c.Api.RPC.Chain.GetBlock(hash)
	if block != nil {
		currentBlock := int64(block.Block.Header.Number)
		fmt.Printf("block is %v\n", currentBlock)
	}

	d, _ := json.Marshal(resp)
	fmt.Println(string(d))
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
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
