package main

import "github.com/urfave/cli/v2"

var (
	ConfigFlag = &cli.StringFlag{
		Name:    "config",
		Aliases: []string{"c"},
		Value:   "config.json",
		Usage:   "load configuration from `file`",
	}
	OwnerFlag = &cli.StringFlag{
		Name:    "owner",
		Aliases: []string{"r"},
		Usage:   "owner address",
	}
	addressListFlag = &cli.StringFlag{
		Name:    "addressList",
		Aliases: []string{"f"},
		Usage:   "owner address list `file`, in csv format",
	}
	amountFlag = &cli.IntFlag{
		Name:    "amount",
		Aliases: []string{"m"},
		Usage:   "amount of NFT",
	}
)
