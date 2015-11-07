package misc

import (
	"fmt"
	"github.com/stvp/assert"
	"testing"
	"time"
)

const (
	secretPassword = "super-duper-secret-password"
)

func Test_ComparePasswords(t *testing.T) {
	// Hash and salt the "real" password (for our purposes)
	hashBase64, _ := HashPasswordDefaultCost(secretPassword)
	fmt.Println("Hashed and base64 encoded:", hashBase64)

	// Compare with an invalid password
	assert.False(t, ComparePassword(hashBase64, "some invalid password"), "error, invalid password matched")

	// Compare with a valid password
	assert.True(t, ComparePassword(hashBase64, secretPassword), "error, valid password did not match")
}

func Test_HashCost(t *testing.T) {
	hashCost := 13
	start := time.Now()
	// Hash the password so we can test the cost of hashing it
	hashBase64, _ := HashPassword(secretPassword, hashCost)
	duration := time.Since(start)
	fmt.Println("Hashed and base64 encoded:", hashBase64)
	fmt.Println("time to hash password with cost", hashCost, "is:", duration, "(should be 0.5 to 1.5 seconds)")

	// Compare with an invalid password
	cost, _ := HashCost(hashBase64)
	assert.Equal(t, hashCost, cost, "error, invalid cost")
}
