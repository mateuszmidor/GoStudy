package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
	"github.com/google/uuid"
)

type CustomClaims struct {
	Email      string `json:"email"`
	jwt.Claims        // Embed standard claims
}

func createJWT(email string) string {
	userIDbytes := sha256.Sum256([]byte(email))
	userIDhex := hex.EncodeToString(userIDbytes[:])
	now := time.Now().UTC()

	customClaims := CustomClaims{
		Email: email,
		Claims: jwt.Claims{
			Subject:   userIDhex,                                    // sub: unique user identifier
			Issuer:    "webauth",                                    // iss: who issued this token
			Audience:  jwt.Audience{"webauth"},                      // aud: intended recipient
			Expiry:    jwt.NewNumericDate(now.Add(5 * time.Minute)), // exp: token expires in 5 min
			IssuedAt:  jwt.NewNumericDate(now),                      // iat: when token was created
			NotBefore: jwt.NewNumericDate(now),                      // nbf: token valid immediately
			ID:        uuid.New().String(),                          // jti: unique token identifier for replay prevention
		},
	}

	// Use RS256 (asymmetric)
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: privateKey}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		panic(err)
	}

	builder := jwt.Signed(signer).Claims(customClaims)
	tokenStr, err := builder.CompactSerialize()
	if err != nil {
		panic(err)
	}

	return tokenStr
}

func validateAndParseJWT(tokenString string) *CustomClaims {
	token, err := jwt.ParseSigned(tokenString)
	if err != nil {
		panic(err)
	}

	claims := &CustomClaims{}
	if err := token.Claims(publicKey, claims); err != nil {
		panic(err)
	}

	// Validate standard claims
	expected := jwt.Expected{
		Issuer:   "webauth",
		Audience: jwt.Audience{"webauth"},
		Time:     time.Now().UTC(),
	}

	if err := claims.Validate(expected); err != nil {
		panic(err)
	}

	return claims
}

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func init() {
	var err error
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey = &privateKey.PublicKey
}
