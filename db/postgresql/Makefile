up:
	# database: postgres, user: postgres
	docker run --name my-postgres -p 5432:5432 -e POSTGRES_PASSWORD=pass postgres

fill:
	@echo "!!! Database password is: pass !!!"
	psql -h localhost -p 5432 -U postgres -d postgres -f ./series03-final-inserts.sql

down:
	docker stop my-postgres