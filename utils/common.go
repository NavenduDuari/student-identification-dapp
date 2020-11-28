package utils

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

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

func GenerateId(data []byte) string {
	h := sha1.New()
	h.Write(data)
	sha1Hash := hex.EncodeToString(h.Sum(nil))
	return sha1Hash
}
