default:
	docker run \
		-d \
		--rm \
		--name=my-vault \
		--cap-add=IPC_LOCK \
		-p 8200:8200 \
		-e 'VAULT_DEV_ROOT_TOKEN_ID=root-token' \
		-e 'VAULT_DEV_LISTEN_ADDRESS=0.0.0.0:8200' \
		hashicorp/vault:latest
	go run .

stop:
	docker stop my-vault
