package main

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	ctx := context.Background()

	// prepare a client with the given base address
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	// authenticate with a root token (insecure)
	if err := client.SetToken("root-token"); err != nil {
		log.Fatal(err)
	}

	// write a secret
	req := schema.KvV2WriteRequest{
		Data: map[string]any{
			"password1": "abc123",
			"password2": "correct horse battery staple",
		},
	}
	_, err = client.Secrets.KvV2Write(ctx, "my-secret", req, vault.WithMountPath("secret")) // "secret" is the Vault's default Key-Value engine name
	if err != nil {
		log.Fatal(err)
	}
	log.Println("secret written successfully")

	// read a secret
	s, err := client.Secrets.KvV2Read(ctx, "my-secret", vault.WithMountPath("secret"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("secret retrieved:", s.Data.Data)
}
