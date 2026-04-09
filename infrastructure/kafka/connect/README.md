# kafka connect demo

Watch the file `process_list.txt` and push it's contents into kafka topic `process_list`.

**note:** the `process_list.txt` is mounted in `kafka-connect` container under `/tmp/process_list.txt`.

## Run

```
docker-compose up # and wait until kafka-connect is up
```

```sh
make
```

```sh
firefox localhost:8080 # kafka GUI - see topic with file contents
```