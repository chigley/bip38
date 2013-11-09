package main

import (
	"bytes"
	"code.google.com/p/go.crypto/scrypt"
	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"
	"github.com/piotrnar/gocoin/btc"
	"log"
)

func main() {
	encryptedKey := "6PfMxA1n3cqYarHoDqPRPLpBBJGWLDY1qX94z8Qyjg7XAMNZJMvHLqAMyS"
	passphrase := "AaAaB"

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
		log.Printf("Has lot/sequence: %t", hasLotSequence)

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

			lotNumber := int(ownerSalt[4])*4096 + int(ownerSalt[5])*16 + int(ownerSalt[6])/16
			sequenceNumber := int(ownerSalt[6]&0x0f)*256 + int(ownerSalt[7])

			log.Printf("Lot number: %d", lotNumber)
			log.Printf("Sequence number: %d", sequenceNumber)
		} else {
			passFactor = prefactorA
		}

		log.Printf("passfactor: %s", hex.EncodeToString(passFactor))

		passpoint, err := btc.PublicFromPrivate(passFactor, true)
		if passpoint == nil {
			log.Fatal(err)
		}

		log.Printf("passpoint: %s", hex.EncodeToString(passpoint))

		encryptedpart1 := dec[15:23]
		encryptedpart2 := dec[23:39]

		addresshashplusownerentropy := bytes.Join([][]byte{dec[3:7], ownerSalt[:8]}, nil)

		derived, err := scrypt.Key(passpoint, addresshashplusownerentropy, 1024, 1, 1, 64)
		if derived == nil {
			log.Fatal(err)
		}

		derivedhalf2 := derived[32:]

		h, err := aes.NewCipher(derivedhalf2)
		if h == nil {
			log.Fatal(err)
		}

		unencryptedpart2 := make([]byte, 16)
		h.Decrypt(unencryptedpart2, encryptedpart2)
		h.Decrypt(unencryptedpart2, encryptedpart2) // TODO: necessary?
		for i := range unencryptedpart2 {
			unencryptedpart2[i] ^= derived[i+16]
		}

		encryptedpart1 = bytes.Join([][]byte{encryptedpart1, unencryptedpart2[:8]}, nil)

		unencryptedpart1 := make([]byte, 16)
		h.Decrypt(unencryptedpart1, encryptedpart2)
		h.Decrypt(unencryptedpart2, encryptedpart1) // TODO: necessary?
		for i := range unencryptedpart1 {
			unencryptedpart1[i] ^= derived[i]
		}

		seeddb := bytes.Join([][]byte{unencryptedpart1[:16], unencryptedpart2[8:]}, nil)

		sha := sha256.New()
		sha.Write(seeddb)
		singleHashed := sha.Sum(nil)
		sha.Reset()
		sha.Write(singleHashed)
		factorb := sha.Sum(nil)

		log.Printf("factorb: %s", hex.EncodeToString(factorb))
	} else {
		log.Fatal("Malformed byte slice")
	}
}
