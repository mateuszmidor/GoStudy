run: 
	[ -f .env ] && source .env; \
	PORT=9999 \
	SCOPES="openid profile email" \
	go run app/app.go

	
.PHONY: run