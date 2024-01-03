default:
	docker run \
		-d \
		--rm \
		--name=my-vault \
		--cap-add=IPC_LOCK \
		-p 8200:8200 \
		-e 'VAULT_DEV_ROOT_TOKEN_ID=root-token' \
		-v `pwd`/storage:/vault/storage \
		-v `pwd`/config:/vault/config \
		hashicorp/vault:latest
	sleep 3
	docker exec -e 'VAULT_ADDR=http://127.0.0.1:8200' -it my-vault sh -c 'vault login root-token && vault auth enable userpass'

stop:
	docker stop my-vault
