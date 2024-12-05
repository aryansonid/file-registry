package service

import (
	"context"
	"crypto/ecdsa"
	"file-registry/config"
	fileregistry "file-registry/fileregistry"
	"fmt"
	"math/big"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Service interface {
	SaveDetailsInContract(filePath, cid string) (string, error)
	GetCIDFromContract(filePath string) (string, error)
}

type ServiceRepository struct {
	rpcURL       string
	contractAddr string
	privateKey   string
}

func NewService(config config.Config) *ServiceRepository {
	s := &ServiceRepository{}
	s.SetupConfig(config)
	return s
}

func (s *ServiceRepository) SaveDetailsInContract(filePath, cid string) (string, error) {

	client, err := ethclient.Dial(s.rpcURL)
	if err != nil {
		return "", err
	}
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(s.privateKey)
	if err != nil {
		log.Errorf("Failed to get private key: %v", err)
		return "", fmt.Errorf("failed to get private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Errorf("Failed to get public key: %v", err)
		return "", fmt.Errorf("failed to get public key: %v", err)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Errorf("Failed to get nonce: %v", err)
		return "", fmt.Errorf("failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf("Failed to suggest gas price: %v", err)
		return "", fmt.Errorf("failed to suggest gas price: %v", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Errorf("Failed to get chain ID: %v", err)
		return "", fmt.Errorf("failed to get chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Errorf("Failed to create authorized transactor: %v", err)
		return "", fmt.Errorf("failed to create authorized transactor: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // No ETH is sent with this transaction; just for contract interaction.
	auth.GasLimit = uint64(3000000) // Gas limit
	auth.GasPrice = gasPrice.Mul(big.NewInt(2), gasPrice)

	// Interact with contract
	contractAddress := common.HexToAddress(s.contractAddr)
	instance, err := fileregistry.NewFileregistry(contractAddress, client)
	if err != nil {
		log.Errorf("Failed to instantiate a FileRegistry contract: %v", err)
		return "", fmt.Errorf("failed to instantiate a FileRegistry contract: %v", err)
	}

	tx, err := instance.Save(auth, filePath, cid)
	if err != nil {
		log.Errorf("Failed to save filePath and CID in the smart contract: %v", err)
		return "", fmt.Errorf("failed to save filePath and CID in the smart contract: %v", err)
	}

	log.Infof("Transaction sent: %s", tx.Hash().Hex())

	return tx.Hash().Hex(), nil
}

func (s *ServiceRepository) GetCIDFromContract(filePath string) (string, error) {
	client, err := ethclient.Dial(s.rpcURL)
	if err != nil {
		log.Errorf("Failed to connect to Ethereum client: %v", err)
		return "", fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	defer client.Close()

	contractAddress := common.HexToAddress(s.contractAddr)
	instance, err := fileregistry.NewFileregistry(contractAddress, client)
	if err != nil {
		log.Errorf("Failed to instantiate a FileRegistry contract: %v", err)
		return "", fmt.Errorf("failed to instantiate a FileRegistry contract: %v", err)
	}

	cid, err := instance.Get(&bind.CallOpts{}, filePath)
	if err != nil {
		log.Errorf("Failed to get CID from the smart contract: %v", err)
		return "", fmt.Errorf("failed to get CID from the smart contract: %v", err)
	}

	return cid, nil
}

func (s *ServiceRepository) SetupConfig(v config.Config) {
	s.rpcURL = v.ReadRpcUrlConfig()
	s.contractAddr = v.ReadContractAddressConfig()
	s.privateKey = v.ReadPrivateKey()
}
