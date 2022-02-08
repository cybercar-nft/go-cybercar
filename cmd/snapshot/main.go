package main

import (
	"context"
	"fmt"
	"github.com/cybercar-nft/go-cybercar/cyber"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"
	"math/big"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var app *cli.App

func init() {
	app = &cli.App{
		Name:    filepath.Base(os.Args[0]),
		Usage:   "CyberCar NFT CLI",
		Version: "0.1.0",
		Action:  ownerOf,
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		cancel()
	}()

	if err := app.RunContext(ctx, os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func ownerOf(c *cli.Context) error {
	addr := c.Args().First()
	block := c.Args().Get(1)
	//log.Printf("Contract address: %s, block: %s", addr, block)
	blockNumber, ok := big.NewInt(0).SetString(block, 10)
	if !ok {
		//log.Println("bad block number")
		return nil
	}
	ec, err := ethclient.DialContext(c.Context, "https://api.edennetwork.io/v1/beta")
	if err != nil {
		//log.Printf("connect rpc error: %s", err)
		return err
	}
	nft, err := cyber.NewCar(common.HexToAddress(addr), ec)
	if err != nil {
		//log.Printf("New Car error: %s", err)
		return err
	}
	opts := &bind.CallOpts{
		BlockNumber: blockNumber,
	}
	total, err := nft.TotalSupply(opts)
	if err != nil {
		return err
	}
	//log.Printf("TotalSupply: %s", total.String())
	i := big.NewInt(1)
	one := big.NewInt(1)
	for ; i.Cmp(total) < 1; i.Add(i, one) {
		owner, err := nft.OwnerOf(opts, i)
		if err != nil {
			//log.Printf("Error: %s", err)
			continue
		}
		fmt.Printf("%d,%s\n", i, owner)
	}
	return nil
}
