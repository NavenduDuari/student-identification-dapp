package contractmanager

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/navenduduari/studentidentity/contracts"
	"github.com/navenduduari/studentidentity/utils"
)

func NewSession(ctx context.Context) (session contracts.IdentitySession) {
	envVars := utils.LoadEnv()
	auth, err := bind.NewTransactor(strings.NewReader(envVars["KEY"]), envVars["KEYPASS"])
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// bind.NewTransactor() returns a bind.TransactOpts{} struct with the following field values:
	// From: auth.From,
	// Signer: auth.Signer,
	// Nonce: nil // Setting to nil uses nonce of pending state
	// Value: big.NewInt(0), // 0 because we're not transferring Eth
	// GasPrice: nil // Setting to nil automatically suggests a gas price
	// GasLimit: 0 // Setting to 0 automatically estimates gas limit

	// Return session without contract instance
	return contracts.IdentitySession{
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}
}

// NewContract deploys a contract if no existing contract exists
func NewContract(session contracts.IdentitySession, client *ethclient.Client) contracts.IdentitySession {
	envVars := utils.LoadEnv()

	if envVars["CONTRACTADDR"] != "" {
		return session
	}

	// Hash answer before sending it over Ethereum network.
	contractAddress, tx, instance, err := contracts.DeployIdentity(&session.TransactOpts, client)
	if err != nil {
		log.Fatalf("could not deploy contract: %v\n", err)
	}

	fmt.Printf("Contract deployed! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())

	session.Contract = instance
	utils.UpdateEnvFile("CONTRACTADDR", contractAddress.Hex())
	return session
}

// LoadContract loads a contract if one exists
func LoadContract(session contracts.IdentitySession, client *ethclient.Client) contracts.IdentitySession {
	envVars := utils.LoadEnv()
	if envVars["CONTRACTADDR"] == "" {
		log.Println("could not find a contract address to load")
		return session
	}
	addr := common.HexToAddress(envVars["CONTRACTADDR"])
	instance, err := contracts.NewIdentity(addr, client)
	if err != nil {
		log.Fatalf("could not load contract: %v\n", err)
	}

	session.Contract = instance
	return session
}
