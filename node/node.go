package node

import (
	"context"
	"errors"
	"fmt"
	"github.com/cybercar-nft/go-cybercar/cyber"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/xyths/hs"
	"go.uber.org/zap"
	"io/ioutil"
	"math/big"
	"os"
	"time"
)

type Config struct {
	Log      hs.LogConf `json:"log"`
	RPC      string     `json:"rpc"`
	Contract string     `json:"contract"`
	Mnemonic string     `json:"mnemonic"`
	Account  int        `json:"account"`
}

type Node struct {
	cfg Config

	Sugar *zap.SugaredLogger

	wallet  *hdwallet.Wallet
	account accounts.Account

	ec  *ethclient.Client
	nft *cyber.Car
}

func New(cfg Config) *Node {
	return &Node{
		cfg: cfg,
	}
}

func (n *Node) Init(ctx context.Context) error {
	l, err := hs.NewZapLogger(n.cfg.Log)
	if err != nil {
		return err
	}
	n.Sugar = l.Sugar()
	n.Sugar.Info("logger initialized")

	mnemonic, err := loadMnemonic(n.cfg.Mnemonic)
	if err != nil {
		n.Sugar.Errorf("load mnemonic error: %s", err)
		return err
	}

	n.wallet, err = hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		n.Sugar.Errorf("new hd wallet error: %s", err)
		return err
	}

	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", n.cfg.Account))
	n.account, err = n.wallet.Derive(path, false)
	if err != nil {
		n.Sugar.Errorf("derive account error: %s", err)
		return err
	}
	n.Sugar.Info("wallet initialized")

	n.ec, err = ethclient.DialContext(ctx, n.cfg.RPC)
	if err != nil {
		n.Sugar.Errorf("connect rpc error: %s", err)
		return err
	}
	n.Sugar.Info("dial success")
	n.nft, err = cyber.NewCar(common.HexToAddress(n.cfg.Contract), n.ec)
	if err != nil {
		n.Sugar.Errorf("New Car error: %s", err)
		return err
	}
	n.Sugar.Info("initialize success")
	return nil
}

type Quota struct {
	Minted uint8
	Cap    uint8
}

func (n *Node) AirdropQuota(ctx context.Context, owner common.Address) (Quota, error) {
	return n.nft.AirdropQuota(&bind.CallOpts{Context: ctx}, owner)
}

func (n *Node) Paused(ctx context.Context) (bool, error) {
	return n.nft.Paused(&bind.CallOpts{Context: ctx})
}

func (n *Node) Phase(ctx context.Context) (int8, error) {
	return n.nft.Phase(&bind.CallOpts{Context: ctx})
}

func (n *Node) MintQuota(ctx context.Context, owner common.Address) (Quota, error) {
	return n.nft.MintQuota(&bind.CallOpts{Context: ctx}, owner)
}

func (n *Node) AddAirdrop(ctx context.Context, owners []common.Address, amount uint8) error {
	for _, owner := range owners {
		n.Sugar.Infof("AddAirdrop for %s", owner.String())
	}
	chainId, err := n.ec.ChainID(ctx)
	if err != nil {
		n.Sugar.Errorf("Get chainId error: %s", err)
		return err
	}
	privateKey, err := n.wallet.PrivateKey(n.account)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		n.Sugar.Errorf("Get transactor error: %s", err)
		return err
	}
	nonce, err := n.ec.NonceAt(ctx, n.account.Address, nil)
	if err != nil {
		n.Sugar.Errorf("Get nonce error: %s", err)
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units
	gasPrice, err := n.ec.SuggestGasPrice(ctx)
	if err != nil {
		n.Sugar.Errorf("SuggestGasPrice error: %s", err)
		return err
	}
	auth.GasPrice = gasPrice

	tx, err := n.nft.AddAirdrop(auth, owners, amount)
	if err != nil {
		n.Sugar.Errorf("AddAirdrop error: %s", err)
		return err
	}
	pending := true
	for pending {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Second):
			_, isPending, err := n.ec.TransactionByHash(ctx, tx.Hash())
			if err != nil {
				return err
			}
			if !isPending {
				pending = false // break `for`
			}
		}
	}
	receipt, err := n.ec.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return err
	}
	if receipt.Status == 0 {
		msg := fmt.Sprintf("transaction reverted, hash %s", receipt.TxHash.String())
		return errors.New(msg)
	}

	return nil
}

func (n *Node) AddWhitelist(ctx context.Context, owners []common.Address, amount uint8) error {
	chainId, err := n.ec.ChainID(ctx)
	if err != nil {
		n.Sugar.Errorf("Get chainId error: %s", err)
		return err
	}
	privateKey, err := n.wallet.PrivateKey(n.account)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		n.Sugar.Errorf("Get transactor error: %s", err)
		return err
	}
	nonce, err := n.ec.NonceAt(ctx, n.account.Address, nil)
	if err != nil {
		n.Sugar.Errorf("Get nonce error: %s", err)
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units
	gasPrice, err := n.ec.SuggestGasPrice(ctx)
	if err != nil {
		n.Sugar.Errorf("SuggestGasPrice error: %s", err)
		return err
	}
	auth.GasPrice = gasPrice

	tx, err := n.nft.AddWhitelist(auth, owners, amount)
	if err != nil {
		n.Sugar.Errorf("AddWhitelist error: %s", err)
		return err
	}
	pending := true
	for pending {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Second):
			_, isPending, err := n.ec.TransactionByHash(ctx, tx.Hash())
			if err != nil {
				return err
			}
			if !isPending {
				pending = false // break `for`
			}
		}
	}
	receipt, err := n.ec.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return err
	}
	if receipt.Status == 0 {
		msg := fmt.Sprintf("transaction reverted, hash %s", receipt.TxHash.String())
		return errors.New(msg)
	}

	return nil
}

func (n *Node) Pause(ctx context.Context) error {
	paused, err := n.Paused(ctx)
	if err != nil {
		n.Sugar.Errorf("check paused error: %s", err)
		return err
	}
	if paused {
		n.Sugar.Info("already paused")
		return nil
	}
	chainId, err := n.ec.ChainID(ctx)
	if err != nil {
		n.Sugar.Errorf("Get chainId error: %s", err)
		return err
	}
	privateKey, err := n.wallet.PrivateKey(n.account)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		n.Sugar.Errorf("Get transactor error: %s", err)
		return err
	}
	nonce, err := n.ec.NonceAt(ctx, n.account.Address, nil)
	if err != nil {
		n.Sugar.Errorf("Get nonce error: %s", err)
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units
	gasPrice, err := n.ec.SuggestGasPrice(ctx)
	if err != nil {
		n.Sugar.Errorf("SuggestGasPrice error: %s", err)
		return err
	}
	auth.GasPrice = gasPrice

	tx, err := n.nft.Pause(auth)
	if err != nil {
		n.Sugar.Errorf("Pause error: %s", err)
		return err
	}
	pending := true
	for pending {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Second):
			_, isPending, err := n.ec.TransactionByHash(ctx, tx.Hash())
			if err != nil {
				return err
			}
			if !isPending {
				pending = false // break `for`
			}
		}
	}
	receipt, err := n.ec.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return err
	}
	if receipt.Status == 0 {
		msg := fmt.Sprintf("transaction reverted, hash %s", receipt.TxHash.String())
		return errors.New(msg)
	}

	return nil
}

func (n *Node) Unpause(ctx context.Context) error {
	paused, err := n.Paused(ctx)
	if err != nil {
		n.Sugar.Errorf("check paused error: %s", err)
		return err
	}
	if !paused {
		n.Sugar.Info("already non-paused")
		return nil
	}
	chainId, err := n.ec.ChainID(ctx)
	if err != nil {
		n.Sugar.Errorf("Get chainId error: %s", err)
		return err
	}
	privateKey, err := n.wallet.PrivateKey(n.account)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		n.Sugar.Errorf("Get transactor error: %s", err)
		return err
	}
	nonce, err := n.ec.NonceAt(ctx, n.account.Address, nil)
	if err != nil {
		n.Sugar.Errorf("Get nonce error: %s", err)
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units
	gasPrice, err := n.ec.SuggestGasPrice(ctx)
	if err != nil {
		n.Sugar.Errorf("SuggestGasPrice error: %s", err)
		return err
	}
	auth.GasPrice = gasPrice

	tx, err := n.nft.Unpause(auth)
	if err != nil {
		n.Sugar.Errorf("Pause error: %s", err)
		return err
	}
	pending := true
	for pending {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Second):
			_, isPending, err := n.ec.TransactionByHash(ctx, tx.Hash())
			if err != nil {
				return err
			}
			if !isPending {
				pending = false // break `for`
			}
		}
	}
	receipt, err := n.ec.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return err
	}
	if receipt.Status == 0 {
		msg := fmt.Sprintf("transaction reverted, hash %s", receipt.TxHash.String())
		return errors.New(msg)
	}

	return nil
}

func (n *Node) SetPhase(ctx context.Context, newPhase int8) error {
	chainId, err := n.ec.ChainID(ctx)
	if err != nil {
		n.Sugar.Errorf("Get chainId error: %s", err)
		return err
	}
	privateKey, err := n.wallet.PrivateKey(n.account)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		n.Sugar.Errorf("Get transactor error: %s", err)
		return err
	}
	nonce, err := n.ec.NonceAt(ctx, n.account.Address, nil)
	if err != nil {
		n.Sugar.Errorf("Get nonce error: %s", err)
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units
	gasPrice, err := n.ec.SuggestGasPrice(ctx)
	if err != nil {
		n.Sugar.Errorf("SuggestGasPrice error: %s", err)
		return err
	}
	auth.GasPrice = gasPrice

	tx, err := n.nft.SetPhase(auth, newPhase)
	if err != nil {
		n.Sugar.Errorf("AddWhitelist error: %s", err)
		return err
	}
	pending := true
	for pending {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Second):
			_, isPending, err := n.ec.TransactionByHash(ctx, tx.Hash())
			if err != nil {
				return err
			}
			if !isPending {
				pending = false // break `for`
			}
		}
	}
	receipt, err := n.ec.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return err
	}
	if receipt.Status == 0 {
		msg := fmt.Sprintf("transaction reverted, hash %s", receipt.TxHash.String())
		return errors.New(msg)
	}

	return nil
}

func loadMnemonic(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
