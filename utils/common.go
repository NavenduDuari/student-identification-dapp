package utils

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/navenduduari/studentidentity/contracts"
)

func NewSession(ctx context.Context) (session contracts.IdentitySession) {
	envVars := LoadEnv()
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
	envVars := LoadEnv()

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
	UpdateEnvFile("CONTRACTADDR", contractAddress.Hex())
	return session
}

// LoadContract loads a contract if one exists
func LoadContract(session contracts.IdentitySession, client *ethclient.Client) contracts.IdentitySession {
	envVars := LoadEnv()
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

func LoadEnv() map[string]string {
	if envVars, err := godotenv.Read(".env"); err != nil {
		log.Printf("could not load env : %v", err)
	} else {
		return envVars
	}

	return nil
}

func UpdateEnvFile(k string, val string) {
	envVars := LoadEnv()
	envVars[k] = val
	err := godotenv.Write(envVars, ".env")
	if err != nil {
		log.Printf("failed to update %s: %v\n", err)
	}
}

func ReadStringStdin() string {
	reader := bufio.NewReader(os.Stdin)
	inputVal, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("invalid input: %v\n", err)
		return ""
	}

	output := strings.TrimSuffix(inputVal, "\n")
	return output
}
