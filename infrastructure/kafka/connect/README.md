# kafka connect demo

Watch the file `process_list.txt` and push it's contents into kafka topic `process_list`.

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