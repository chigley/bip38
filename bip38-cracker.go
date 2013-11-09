package main

import (
	"encoding/hex"
	"github.com/piotrnar/gocoin/btc"
	"log"
)

func main() {
	encryptedKey := "6PfMxA1n3cqYarHoDqPRPLpBBJGWLDY1qX94z8Qyjg7XAMNZJMvHLqAMyS"

	dec := btc.Decodeb58(encryptedKey)[:39] // trim to length 39 (not sure why needed)
	if dec == nil {
		log.Fatal("Cannot decode base58 string " + encryptedKey)
	}

	log.Printf("Decoded base58 string to %s (length %d)", hex.EncodeToString(dec), len(dec))

	if dec[0] == 0x01 && dec[1] == 0x42 {
		log.Print("EC multiply mode not used")

	} else if dec[0] == 0x01 && dec[1] == 0x43 {
		log.Print("EC multiply mode used")

		ownerSalt := dec[7:15]
		includeHashStep := dec[2]&0x04 == 0x04

		log.Printf("Owner salt: %s", hex.EncodeToString(ownerSalt))
		log.Printf("Include hash step: %t", includeHashStep)
	} else {
		log.Fatal("Malformed byte slice")
	}
}
