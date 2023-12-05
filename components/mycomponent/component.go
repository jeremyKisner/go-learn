package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Identifcation struct {
	ID   int    `json:"id"`
	Hash string `json:"hash"`
}

func hashString(input string) string {
	// Create a new SHA-256 hasher
	hasher := sha256.New()

	// Write the string to the hasher
	hasher.Write([]byte(input))

	// Get the hashed bytes
	hashedBytes := hasher.Sum(nil)

	// Convert the hashed bytes to a hex-encoded string
	hashedString := hex.EncodeToString(hashedBytes)

	return hashedString
}

type Component struct {
	Name          string        `json:"name"`
	Identifcation Identifcation `json:"identification"`
}

func main() {
	name := "foobar"
	id := 1
	comp := &Component{
		Name: name,
		Identifcation: Identifcation{
			ID:   id,
			Hash: hashString(name + fmt.Sprint(id)),
		},
	}
	fmt.Println(comp)
	res, err := json.Marshal(comp)
	if err != nil {
		return
	}
	fmt.Println(string(res))
}
