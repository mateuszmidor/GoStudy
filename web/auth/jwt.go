package main

import (
	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
)

type CustomClaims struct {
	Email string `json:"email"`
}

// Use jose/jwt to create a signed JWT token with the "email" claim
// Secret key for HMAC signing (in production, use a secure way to manage this)
const key = "replace-your-secret-key"

func createJWT(email string) string {
	customClaims := CustomClaims{Email: email}

	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: []byte(key)}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		panic(err)
	}

	builder := jwt.Signed(signer).Claims(customClaims).Claims(jwt.Claims{})
	tokenStr, err := builder.CompactSerialize()
	if err != nil {
		panic(err)
	}
	return tokenStr
}

func decodeJWT(token string) (email string) {
	tok, err := jwt.ParseSigned(token)
	if err != nil {
		panic(err)
	}

	var customClaims CustomClaims
	var stdClaims jwt.Claims

	// Try to verify signature and decode
	err = tok.Claims([]byte(key), &customClaims, &stdClaims)
	if err != nil {
		panic(err)
	}

	return customClaims.Email
}
