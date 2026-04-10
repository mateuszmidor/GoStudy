# kafka connect demo

Create a kafka connector to watch the file `process_list.txt` and automatically push it's contents into kafka topic `process_list`.

**note:** the `process_list.txt` is mounted in `kafka_connect` container under `/tmp/process_list.txt` for watching.

## Run

```sh
# run kafka_broker, kafka_connect, kafka_gui
docker-compose up # and wait until kafka-connect is up
```

```sh
# create new connector to watch process_list.txt and populate it's source file: process_list.txt
make
```

```sh
# kafka GUI - see topic "process_list" with process_list.txt file contents
firefox localhost:8080
```