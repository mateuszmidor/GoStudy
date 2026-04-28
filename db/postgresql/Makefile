up:
	# database: postgres, user: postgres
	docker run --name my-postgres -p 5432:5432 -e POSTGRES_PASSWORD=pass postgres

fill:
	@echo "!!! Database password is: pass !!!"
	psql -h localhost -p 5432 -U postgres -d postgres -f ./series03-final-inserts.sql
	psql -h localhost -p 5432 -U postgres -d postgres -f ./series05-final-inserts.sql
	[[ ! -f series07-final-inserts.sql ]] && gunzip series07-final-inserts.sql.gz
	psql -h localhost -p 5432 -U postgres -d postgres -f ./series07-final-inserts.sql


start:
	docker start my-postgres # if container already created but was stopped

down:
	docker stop my-postgres