//
// hashpassword.go wraps the bcrypt password hashing library
// with base64 encoding so we can deal in strings
//
package misc

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const (
	defaultHashCost = 10 // The cost we want to incur hashing a password
)

// HashPassword hashes the clear-text password and encodes it as base64,
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), defaultHashCost)
	if err != nil {
		return "", err
	}

	// Encode the hash as base64 and return
	hashBase64 := base64.StdEncoding.EncodeToString(hashedBytes)

	return hashBase64, nil
}

// ComparePassword hashes the test password and then compares
// the two hashes.
func ComparePassword(hashBase64, testPassword string) bool {

	// Decode the hashed password so bcrypt can compare
	hashBytes, err := base64.StdEncoding.DecodeString(hashBase64)
	if err != nil {
		fmt.Println("Error, we were given invalid base64 string", err)
		return false
	}

	err = bcrypt.CompareHashAndPassword(hashBytes, []byte(testPassword))
	return err == nil
}

// HashCost returns how much it cost (1-31) to hash this password
func HashCost(hashBase64 string) (int, error) {

	// Decode the hashed password so we can get the cost
	hashBytes, err := base64.StdEncoding.DecodeString(hashBase64)
	if err != nil {
		return -1, err
	}

	return bcrypt.Cost(hashBytes)
}
