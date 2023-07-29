package main

import (
	"fmt"
	"time"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
)

type PrivateClaims struct {
	Color string `json:"color"` // struct tags is important to make claims compatible across json libraries
}

var sharedKey = []byte("secret")

func main() {
	// 1. write JWT

	// prepare signer
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: sharedKey}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		panic(err)
	}

	// include standard public claims
	publicClaims := jwt.Claims{
		Subject:   "my-subject",
		Issuer:    "my-issuer",
		ID:        "123",
		NotBefore: jwt.NewNumericDate(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
		Audience:  jwt.Audience{"andrzej", "krystyna"},
	}

	// include some custom, private claims
	privateCl := PrivateClaims{
		"green",
	}

	// create the JWT
	rawToken, err := jwt.Signed(signer).Claims(publicClaims).Claims(privateCl).CompactSerialize()
	if err != nil {
		panic(err)
	}

	fmt.Printf("raw JWT:\n%s\n\n", rawToken)

	// 2. read JWT
	tok, err := jwt.ParseSigned(rawToken)
	if err != nil {
		panic(err)
	}

	var publicClaimsOut jwt.Claims
	if err := tok.Claims(sharedKey, &publicClaimsOut); err != nil {
		panic(err)
	}
	var privateClaimsOut PrivateClaims
	if err := tok.Claims(sharedKey, &privateClaimsOut); err != nil {
		panic(err)
	}

	fmt.Printf("publicClaims:\n%+v\n\n", publicClaimsOut)
	fmt.Printf("privateClaims:\n%+v\n", privateClaimsOut)
}
