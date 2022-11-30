package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Hash password with Bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

// Check if two passwords match using Bcrypt's CompareHashAndPassword
// which return nil on success and an error on failure.
func doPasswordsMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currPassword))
	log.Println("password hashes seem to match!")

	if err != nil {
		log.Fatal("incorrect password")
	}
	return true

}

func hash(password string) string {
	hash := sha256.Sum256([]byte(password))
	output := hex.EncodeToString(hash[:])
	return output
}
