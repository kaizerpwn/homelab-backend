package utils

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/square/go-jose"
	"golang.org/x/crypto/hkdf"
)

func VerifyToken(c *gin.Context) {
	// >> Get cookie from request
	cookie, err := c.Request.Cookie("next-auth.session-token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}

	tokenString := cookie.Value

	// >> Generate the decryption key with an hdkf lib
	nextAuthSecret := os.Getenv("NEXTAUTH_SECRET")
	info := "NextAuth.js Generated Encryption Key"

	hash := sha256.New
	kdf := hkdf.New(hash, []byte(nextAuthSecret), []byte(""), []byte(info))
	encryptionKey := make([]byte, 32)
	_, _ = io.ReadFull(kdf, encryptionKey)

	// >> Parse and decrypt the JWE token
	parsedToken, err := jose.ParseEncrypted(tokenString)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden", "details": err.Error()})
		c.Abort()
		return
	}

	// >> Decrypt the token
	decryptedPayload, err := parsedToken.Decrypt(encryptionKey)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden", "details": err.Error()})
		c.Abort()
		return
	}

	payloadString := string(decryptedPayload)

	// Access the claims from the payload
	fmt.Println(payloadString)

	c.Next()
}
