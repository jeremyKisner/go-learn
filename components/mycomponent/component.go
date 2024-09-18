package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Container struct {
	Components []Component `json:"components"`
	Hash       string      `json:"hash"`
}
type Component struct {
	Name           string         `json:"name"`
	Identification Identification `json:"identification"`
}

type Identification struct {
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

func main() {
	var comps []Component
	comp := Component{
		Name: "foobar",
		Identification: Identification{
			ID:   1,
			Hash: hashString("foobar" + fmt.Sprint(1)),
		},
	}
	comps = append(comps, comp)
	c := Container{
		Components: comps,
	}
	res, err := json.Marshal(c)
	if err != nil {
		return
	}
	fmt.Println(string(res))
}
