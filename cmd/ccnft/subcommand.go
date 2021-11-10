package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/cybercar-nft/go-cybercar/node"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
	"github.com/xyths/hs"
	"os"
	"strconv"
)

var (
	userCommand = &cli.Command{
		Name:    "user",
		Aliases: []string{"u"},
		Usage:   "User interfaces to interactive with the NFT contract",
		Flags: []cli.Flag{
		},
		Subcommands: []*cli.Command{
			{
				Action: airdropQuota,
				Name:   "airdropQuota",
				Usage:  "check airdrop quota of an owner",
				Flags: []cli.Flag{
					OwnerFlag,
				},
			},
			{
				Action: paused,
				Name:   "paused",
				Usage:  "check if contract paused",
			},
			{
				Action: phase,
				Name:   "phase",
				Usage:  "check mint phase",
			},
			{
				Action: mintQuota,
				Name:   "mintQuota",
				Usage:  "check mint quota of an whitelist owner",
				Flags: []cli.Flag{
					OwnerFlag,
				},
			},
		},
	}
	adminCommand = &cli.Command{
		Name:    "admin",
		Aliases: []string{"a"},
		Usage:   "Admin interfaces to manage the NFT contract",
		Flags: []cli.Flag{
		},
		Subcommands: []*cli.Command{
			{
				Action: addAirdrop,
				Name:   "addAirdrop",
				Usage:  "add airdrop list with quota",
				Flags: []cli.Flag{
					addressListFlag,
					amountFlag,
				},
			},
			{
				Action: addWhitelist,
				Name:   "addWhitelist",
				Usage:  "add mint whitelist with quota",
				Flags: []cli.Flag{
					addressListFlag,
					amountFlag,
				},
			},
			{
				Action: pause,
				Name:   "pause",
				Usage:  "pause",
			},
			{
				Action: unpause,
				Name:   "unpause",
				Usage:  "unpause",
			},
			{
				Action: setPhase,
				Name:   "setPhase",
				Usage:  "set phase of operation",
				Flags: []cli.Flag{
				},
				ArgsUsage: "phase (0-2)",
			},
		},
	}
)

func airdropQuota(ctx *cli.Context) error {
	owner := ctx.String(OwnerFlag.Name)
	configFile := ctx.String(ConfigFlag.Name)
	cfg := node.Config{}
	if err := hs.ParseJsonConfig(configFile, &cfg); err != nil {
		return err
	}
	s := node.New(cfg)
	if err := s.Init(ctx.Context); err != nil {
		return err
	}
	quota, err := s.AirdropQuota(ctx.Context, common.HexToAddress(owner))
	if err != nil {
		return err
	}
	fmt.Printf("Minted: %d, Cap: %d", quota.Minted, quota.Cap)
	return nil
}

func paused(ctx *cli.Context) error {
	configFile := ctx.String(ConfigFlag.Name)
	cfg := node.Config{}
	if err := hs.ParseJsonConfig(configFile, &cfg); err != nil {
		return err
	}
	s := node.New(cfg)
	if err := s.Init(ctx.Context); err != nil {
		return err
	}
	p, err := s.Paused(ctx.Context)
	if err != nil {
		return err
	}
	fmt.Printf("Paused: %v", p)
	return nil
}

func phase(ctx *cli.Context) error {
	configFile := ctx.String(ConfigFlag.Name)
	cfg := node.Config{}
	if err := hs.ParseJsonConfig(configFile, &cfg); err != nil {
		return err
	}
	s := node.New(cfg)
	if err := s.Init(ctx.Context); err != nil {
		return err
	}
	p, err := s.Phase(ctx.Context)
	if err != nil {
		return err
	}
	fmt.Printf("Phase: %d", p)
	return nil
}

func mintQuota(ctx *cli.Context) error {
	owner := ctx.String(OwnerFlag.Name)
	configFile := ctx.String(ConfigFlag.Name)
	cfg := node.Config{}
	if err := hs.ParseJsonConfig(configFile, &cfg); err != nil {
		return err
	}
	s := node.New(cfg)
	if err := s.Init(ctx.Context); err != nil {
		return err
	}
	quota, err := s.MintQuota(ctx.Context, common.HexToAddress(owner))
	if err != nil {
		return err
	}
	fmt.Printf("Minted: %d, Cap: %d", quota.Minted, quota.Cap)
	return nil
}

func addAirdrop(ctx *cli.Context) error {
	amount := ctx.Int(amountFlag.Name)
	addrFile := ctx.String(addressListFlag.Name)
	f, err := os.Open(addrFile)
	if err != nil {
		return err
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}
	var owners []common.Address
	for _, line := range records {
		for _, token := range line {
			owners = append(owners, common.HexToAddress(token))
		}
	}
	configFile := ctx.String(ConfigFlag.Name)
	cfg := node.Config{}
	if err = hs.ParseJsonConfig(configFile, &cfg); err != nil {
		return err
	}
	s := node.New(cfg)
	if err := s.Init(ctx.Context); err != nil {
		return err
	}
	err = s.AddAirdrop(ctx.Context, owners, uint8(amount))
	if err != nil {
		return err
	}
	return nil
}

func addWhitelist(ctx *cli.Context) error {
	amount := ctx.Int(amountFlag.Name)
	addrFile := ctx.String(addressListFlag.Name)
	f, err := os.Open(addrFile)
	if err != nil {
		return err
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}
	var owners []common.Address
	for _, line := range records {
		for _, token := range line {
			owners = append(owners, common.HexToAddress(token))
		}
	}
	configFile := ctx.String(ConfigFlag.Name)
	cfg := node.Config{}
	if err = hs.ParseJsonConfig(configFile, &cfg); err != nil {
		return err
	}
	s := node.New(cfg)
	if err := s.Init(ctx.Context); err != nil {
		return err
	}
	err = s.AddWhitelist(ctx.Context, owners, uint8(amount))
	if err != nil {
		return err
	}
	return nil
}

func pause(ctx *cli.Context) error {
	configFile := ctx.String(ConfigFlag.Name)
	cfg := node.Config{}
	if err := hs.ParseJsonConfig(configFile, &cfg); err != nil {
		return err
	}
	s := node.New(cfg)
	if err := s.Init(ctx.Context); err != nil {
		return err
	}
	if err := s.Pause(ctx.Context); err != nil {
		return err
	}
	return nil
}

func unpause(ctx *cli.Context) error {
	configFile := ctx.String(ConfigFlag.Name)
	cfg := node.Config{}
	if err := hs.ParseJsonConfig(configFile, &cfg); err != nil {
		return err
	}
	s := node.New(cfg)
	if err := s.Init(ctx.Context); err != nil {
		return err
	}
	if err := s.Unpause(ctx.Context); err != nil {
		return err
	}
	return nil
}

func setPhase(ctx *cli.Context) error {
	args := ctx.Args()
	if args.Len() == 0 {
		return errors.New("input phase 0-2")
	}
	phase_, err := strconv.ParseInt(args.Get(0), 0, 8)
	if err != nil {
		return errors.New("input phase should be 0-2")
	}
	newPhase := int8(phase_)
	configFile := ctx.String(ConfigFlag.Name)
	cfg := node.Config{}
	if err = hs.ParseJsonConfig(configFile, &cfg); err != nil {
		return err
	}
	s := node.New(cfg)
	if err = s.Init(ctx.Context); err != nil {
		return err
	}
	err = s.SetPhase(ctx.Context, newPhase)
	if err != nil {
		return err
	}
	return nil
}
