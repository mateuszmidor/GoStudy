package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/go-jose/go-jose"
)

func main() {
	// Generate a public/private key pair to use for this example.
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Instantiate a signer using RSASSA-PSS (SHA512) with the given private key.
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.PS512, Key: privateKey}, nil)
	if err != nil {
		panic(err)
	}

	// Sign a sample payload. Calling the signer returns a protected JWS object,
	// which can then be serialized for output afterwards. An error would
	// indicate a problem in an underlying cryptographic primitive.
	var payload = []byte("Lorem ipsum dolor sit amet")
	fmt.Printf("input:\n%s\n\n", payload)
	object, err := signer.Sign(payload)
	if err != nil {
		panic(err)
	}

	// Serialize the signed object using the JWS JSON Serialization format.
	// Alternatively you can also use the compact format here by calling
	// object.CompactSerialize() instead.
	serialized := object.FullSerialize()
	fmt.Printf("serialized:\n%s\n\n", serialized)

	compactSerialized, err := object.CompactSerialize()
	if err != nil {
		panic(err)
	}
	fmt.Printf("compact serialized:\n%s\n\n", compactSerialized)

	// Parse the serialized, protected JWS object. An error would indicate that
	// the given input did not represent a valid message.
	object, err = jose.ParseSigned(compactSerialized)
	if err != nil {
		panic(err)
	}

	// Now we can verify the signature on the payload. An error here would
	// indicate that the message failed to verify, e.g. because the signature was
	// broken or the message was tampered with.
	output, err := object.Verify(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}

	fmt.Printf("deserialized:\n%s\n", output)

}
