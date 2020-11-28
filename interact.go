package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/navenduduari/studentidentity/contractmanager"
	"github.com/navenduduari/studentidentity/contracts"
	"github.com/navenduduari/studentidentity/utils"
)

var ipfs *shell.Shell

type StudentDetails struct {
	name  string
	dob   string
	email string
}

func main() {
	envVars := utils.LoadEnv()
	ctx := context.Background()

	// connect to an ethereum node  hosted by infura
	client, err := ethclient.Dial(envVars["ETH_GATEWAY"])

	if err != nil {
		log.Fatal(utils.Bold, utils.Blue, "Network ::", utils.Red, " Unable to connect to network :: ", err, utils.Reset)
	} else {
		fmt.Println(utils.Bold, utils.Blue, "Network ::", utils.Green, " Successfully connected", utils.Reset)
	}

	defer client.Close()

	ipfs = shell.NewShell(envVars["IPFS_GATEWAY"])
	if ipfs.IsUp() == false {
		log.Fatal(utils.Bold, utils.Yellow, "IPFS ::", utils.Red, " Unable to connect to IPFS", utils.Reset)
	} else {
		fmt.Println(utils.Bold, utils.Yellow, "IPFS ::", utils.Green, "Successfully connected", utils.Reset)
	}

	session := contractmanager.NewSession(ctx)

	// Load or Deploy contract, and update session with contract instance
	if envVars["CONTRACTADDR"] == "" {
		fmt.Println(utils.Bold, utils.Cyan, "Contract ::", utils.Green, " Deploying new contract", utils.Reset)
		session = contractmanager.NewContract(session, client)
	}

	// If we have an existing contract, load it; if we've deployed a new contract, attempt to load it.
	if envVars["CONTRACTADDR"] != "" {
		fmt.Println(utils.Bold, utils.Cyan, "Contract ::", utils.Green, " Loading old contract", utils.Reset)
		session = contractmanager.LoadContract(session, client)
	}

	// Loop to implement simple CLI
	for {
		fmt.Println(
			utils.Bold,
			utils.Gray,
			"Pick an option:\n"+""+
				"1. setAcademicDetails.\n"+
				"2. getAcademicDetails.\n"+
				"3. Exit.\n"+
				"4. Reset and exit.\n",
			utils.Reset,
		)

		switch utils.ReadStringStdin() {
		case "1":
			setAcademicDetails(session)
			break
		case "2":
			getAcademicDetails(session)
			break
		case "3":
			fmt.Println(utils.Bold, utils.Purple, "Client ::", utils.Reset, " Exiting")
			return
		case "4":
			utils.UpdateEnvFile("CONTRACTADDR", "")
			fmt.Println(utils.Bold, utils.Purple, "Client ::", utils.Reset, " Cleared contract address. Exiting")
			return
		default:
			fmt.Println(utils.Bold, utils.Purple, "Client ::", utils.Reset, " Invalid option. Please try again.")
			break
		}
	}
}

func setAcademicDetails(session contracts.IdentitySession) {

	studentDetails := StudentDetails{}

	fmt.Println(utils.Bold, utils.Purple, "Client ::", utils.Reset, " Enter student name")
	studentDetails.name = utils.ReadStringStdin()

	fmt.Println(utils.Bold, utils.Purple, "Client ::", utils.Reset, " Enter student date of birth(dd/mm/yy)")
	studentDetails.dob = utils.ReadStringStdin()

	fmt.Println(utils.Bold, utils.Purple, "Client ::", utils.Reset, " Enter student email")
	studentDetails.email = utils.ReadStringStdin()

	studentId := utils.GenerateId([]byte(fmt.Sprintf("%v", studentDetails)))
	fmt.Println(utils.Bold, utils.Purple, "Client ::", utils.Green, " StudentId generated :: ", utils.Reset, studentId)

	fmt.Println(utils.Bold, utils.Purple, "Client ::", utils.Reset, " Enter sourcePath")
	sourcePath := utils.ReadStringStdin()

	cid, err := ipfs.AddDir(sourcePath)
	if err != nil {
		log.Fatal(utils.Bold, utils.Yellow, "IPFS ::", utils.Red, " Unable uploaded to IPFS :: ", err, utils.Reset)

		os.Exit(1)
	} else {
		fmt.Println(utils.Bold, utils.Yellow, "IPFS ::", utils.Green, "Successfully uploaded to IPFS", utils.Reset)
	}

	academicDetails := contracts.IdentityAcademicDetailsI{
		DegreeId: cid,
	}

	tx, err := session.SetAcademicDetails(studentId, academicDetails)

	if err != nil {
		log.Fatal(utils.Bold, utils.Cyan, "Contract ::", utils.Red, " Unable to set academic details", err, utils.Reset)
	} else {
		fmt.Println(utils.Bold, utils.Cyan, "Contract ::", utils.Green, " Successfully set academic details. Please wait for tx to be confirmed.", utils.Reset, "\nTx = ", tx.Hash().Hex())
	}
}

func getAcademicDetails(session contracts.IdentitySession) {

	studentDetails := StudentDetails{}

	fmt.Println(utils.Bold, utils.Purple, "Client ::", utils.Reset, " Enter student name")
	studentDetails.name = utils.ReadStringStdin()

	fmt.Println(utils.Bold, utils.Purple, "Client ::", utils.Reset, " Enter student date of birth(dd/mm/yy)")
	studentDetails.dob = utils.ReadStringStdin()

	fmt.Println(utils.Bold, utils.Purple, "Client ::", utils.Reset, " Enter student email")
	studentDetails.email = utils.ReadStringStdin()

	studentId := utils.GenerateId([]byte(fmt.Sprintf("%v", studentDetails)))
	fmt.Println(utils.Bold, utils.Purple, "Client ::", utils.Green, " StudentId generated :: ", utils.Reset, studentId)

	academicDetails, err := session.GetAcademicDetails(studentId)

	if err != nil || academicDetails.DegreeId == "" {
		log.Fatal(utils.Bold, utils.Cyan, "Contract ::", utils.Red, " Unable to get academic details", err, utils.Reset)
	} else {
		fmt.Println(utils.Bold, utils.Cyan, "Contract ::", utils.Green, " Successfully got academic details.", utils.Reset)
	}

	fmt.Println(utils.Bold, utils.Purple, "Client ::", utils.Reset, " Enter targetPath")
	targetPath := utils.ReadStringStdin()

	err = ipfs.Get(academicDetails.DegreeId, targetPath)
	if err != nil {
		log.Fatal(utils.Bold, utils.Yellow, "IPFS ::", utils.Red, " Unable to download from IPFS :: ", err, utils.Reset)
	} else {
		fmt.Println(utils.Bold, utils.Yellow, "IPFS ::", utils.Green, "Successfully downloaded from IPFS", utils.Reset)
	}
}
