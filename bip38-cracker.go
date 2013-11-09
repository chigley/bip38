package main

import (
	"bytes"
	"code.google.com/p/go.crypto/scrypt"
	"crypto/sha256"
	"encoding/hex"
	"github.com/piotrnar/gocoin/btc"
	"log"
)

func main() {
	encryptedKey := "6PgGWtx25kUg8QWvwuJAgorN6k9FbE25rv5dMRwu5SKMnfpfVe5mar2ngH"
	passphrase := ""

	dec := btc.Decodeb58(encryptedKey)[:39] // trim to length 39 (not sure why needed)
	if dec == nil {
		log.Fatal("Cannot decode base58 string " + encryptedKey)
	}

	log.Printf("Decoded base58 string to %s (length %d)", hex.EncodeToString(dec), len(dec))

	if dec[0] == 0x01 && dec[1] == 0x42 {
		log.Print("EC multiply mode not used")
		log.Fatal("TODO: implement decryption when EC multiply mode not used")
	} else if dec[0] == 0x01 && dec[1] == 0x43 {
		log.Print("EC multiply mode used")

		ownerSalt := dec[7:15]
		hasLotSequence := dec[2]&0x04 == 0x04

		log.Printf("Owner salt: %s", hex.EncodeToString(ownerSalt))
		log.Printf("Has lot sequence: %t", hasLotSequence)

		prefactorA, err := scrypt.Key([]byte(passphrase), ownerSalt, 16384, 8, 8, 32)
		if prefactorA == nil {
			log.Fatal(err)
		}

		var passFactor []byte

		if hasLotSequence {
			prefactorB := bytes.Join([][]byte{prefactorA, ownerSalt}, nil)

			h := sha256.New()
			h.Write(prefactorB)
			singleHashed := h.Sum(nil)
			h.Reset()
			h.Write(singleHashed)
			doubleHashed := h.Sum(nil)

			passFactor = doubleHashed
		} else {
			passFactor = prefactorA
		}

		log.Print(passFactor)
	} else {
		log.Fatal("Malformed byte slice")
	}
}
