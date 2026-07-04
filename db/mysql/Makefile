USER = user
PASS = pass

.SILENT:

default: run

run: start-mysql wait-mysql-ready fill-mysql-table run-app stop-mysql 

start-mysql:
	@echo "Start MySQL"
	docker run -d --rm \
		--name=mysql \
		-p=3306:3306 \
		-e=MYSQL_ROOT_PASSWORD=admin \
		-e=MYSQL_USER=${USER} -e=MYSQL_PASSWORD=${PASS} -e=MYSQL_DATABASE=recordings \
		mysql:8.0.32
	@echo

stop-mysql:
	@echo "Stop MySQL"
	docker stop mysql
	@echo

wait-mysql-ready:
	while true; do echo "Waiting for MySQL..."; docker exec -it mysql mysqlsh ${USER}:${PASS}@localhost:3306 --sql -e "show databases;" > /dev/null 2>&1 && break || sleep 3; done
	@echo "Ready!"
	@echo

fill-mysql-table:
	@echo "Insert data into database"
	docker exec -it mysql mysqlsh ${USER}:${PASS}@localhost:3306 --sql -e "`cat data.sql`" 
	@echo

run-app:
	@echo "Run the Go app"
	DBUSER=${USER} DBPASS=${PASS} go run .
	@echo
