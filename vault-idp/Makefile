vault:
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
	while true; do docker exec -e 'VAULT_ADDR=http://127.0.0.1:8200' -it my-vault sh -c 'sleep 1; vault login root-token && vault auth enable userpass'  && exit; echo "connecting to Vault..."; done;

idp:
	go run .

exec:
	docker exec -e 'VAULT_ADDR=http://127.0.0.1:8200' -it my-vault sh 

stop:
	docker stop my-vault
