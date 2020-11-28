package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/navenduduari/studentidentity/contracts"
	"github.com/navenduduari/studentidentity/utils"
)

var ipfs *shell.Shell

func main() {
	envVars := utils.LoadEnv()
	ctx := context.Background()

	// connect to an ethereum node  hosted by infura
	client, err := ethclient.Dial(envVars["GATEWAY"])

	if err != nil {
		log.Fatalf("Network :: Unable to connect to network:%v\n", err)
	} else {
		log.Println(utils.Bold, utils.Green, "Network ::", utils.Reset, "Successfully connected")
	}

	defer client.Close()

	ipfs = shell.NewShell("localhost:5001")
	if ipfs.IsUp() == false {
		log.Fatalf("IPFS :: Unable to connect to IPFS\n")
	} else {
		log.Println(utils.Bold, utils.Green, "IPFS ::", utils.Reset, "Successfully connected")
	}

	session := utils.NewSession(ctx)

	// Load or Deploy contract, and update session with contract instance
	if envVars["CONTRACTADDR"] == "" {
		fmt.Println("Contract :: Deploying new contract")
		session = utils.NewContract(session, client)
	}

	// If we have an existing contract, load it; if we've deployed a new contract, attempt to load it.
	if envVars["CONTRACTADDR"] != "" {
		fmt.Println("Contract :: Loading old contract")
		session = utils.LoadContract(session, client)
	}

	// Loop to implement simple CLI
	for {
		fmt.Printf(
			"Pick an option:\n" + "" +
				"1. setAcademicDetails.\n" +
				"2. getAcademicDetails.\n" +
				"3. Exit.\n" +
				"4. Reset and exit.\n",
		)

		switch utils.ReadStringStdin() {
		case "1":
			setAcademicDetails(session)
			break
		case "2":
			getAcademicDetails(session)
			break
		case "3":
			fmt.Println("Client :: Exiting")
			return
		case "4":
			utils.UpdateEnvFile("CONTRACTADDR", "")
			fmt.Println("Client :: Cleared contract address. Exiting")
			return
		default:
			fmt.Println("Client :: Invalid option. Please try again.")
			break
		}
	}
}

func setAcademicDetails(session contracts.IdentitySession) {
	fmt.Println("Client :: Enter studentId")
	studentId := utils.ReadStringStdin()
	fmt.Println("Client :: Enter sourcePath")
	sourcePath := utils.ReadStringStdin()

	cid, err := ipfs.AddDir(sourcePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	} else {
		fmt.Println("IPFS :: Successfully uploaded to IPFS")
	}

	academicDetails := contracts.IdentityAcademicDetailsI{
		DegreeId: cid,
	}

	tx, err := session.SetAcademicDetails(studentId, academicDetails)

	if err != nil {
		log.Fatalf("Contract :: Unable to set academic details :: %v\n", err)

	} else {
		fmt.Printf("Contract :: Successfully set academic details. Please wait for tx %s to be confirmed.\n", tx.Hash().Hex())
	}
}

func getAcademicDetails(session contracts.IdentitySession) {
	fmt.Println("Client :: Enter studentId")
	studentId := utils.ReadStringStdin()

	academicDetails, err := session.GetAcademicDetails(studentId)

	if err != nil {
		log.Fatalf("Contract :: Unable to get academic details :: %v\n", err)
	} else {
		fmt.Println("Contract :: Successfully got academic details")
	}

	fmt.Println("Client :: Enter targetPath")
	targetPath := utils.ReadStringStdin()

	err = ipfs.Get(academicDetails.DegreeId, targetPath)
	if err != nil {
		log.Fatalf("IPFS :: Unable download from IPFS :: %v\n", err)
	} else {
		fmt.Println("IPFS :: Successfully downloaded from IPFS")
	}
}
